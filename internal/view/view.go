package view

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"gitlab.com/david_mbuvi/go_asterisks"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)

var adminChoices = []string{"Add Task", "Remove Task", "Check todo List", "Add user", "Remove user", "Access users list", "Log out"}
var standardChoices = []string{"Add Task", "Remove Task", "Check todo List", "Log out"}
var commonChoices = []string{"Register/Login", "Exit"}

type model struct {
	cursor  int
	choice  string
	choices []string
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q", "esc":
			return m, tea.Quit

		case "enter":
			// Send the choice on the channel and exit.
			m.choice = m.choices[m.cursor]
			return m, tea.Quit

		case "down", "j":
			m.cursor++
			if m.cursor >= len(m.choices) {
				m.cursor = 0
			}

		case "up", "k":
			m.cursor--
			if m.cursor < 0 {
				m.cursor = len(m.choices) - 1
			}
		}
	}

	return m, nil
}

func (m model) View() string {
	s := strings.Builder{}
	s.WriteString("What kind of Bubble Tea would you like to order?\n\n")

	for i := 0; i < len(m.choices); i++ {
		if m.cursor == i {
			s.WriteString("(•) ")
		} else {
			s.WriteString("( ) ")
		}
		s.WriteString(m.choices[i])
		s.WriteString("\n")
	}
	s.WriteString("\n(press q to quit)\n")

	return s.String()
}

func AdminView() string {

	p := tea.NewProgram(model{0, "", adminChoices})

	// Run returns the model as a tea.Model.
	m, err := p.Run()
	if err != nil {
		fmt.Println("Oh no:", err)
		os.Exit(1)
	}

	// Assert the final tea.Model to our local model and print the choice.
	if m, ok := m.(model); ok && m.choice != "" {
		return m.choice
	}

	// CleanStdout()
	return ""
}

func StandardView() string {
	p := tea.NewProgram(model{0, "", standardChoices})

	// Run returns the model as a tea.Model.
	m, err := p.Run()
	if err != nil {
		fmt.Println("Oh no:", err)
		os.Exit(1)
	}

	// Assert the final tea.Model to our local model and print the choice.
	if m, ok := m.(model); ok && m.choice != "" {
		return m.choice
	}

	// CleanStdout()
	return ""

}

func CommonView() string {
	p := tea.NewProgram(model{0, "", commonChoices})

	// Run returns the model as a tea.Model.
	m, err := p.Run()
	if err != nil {
		fmt.Println("Oh no:", err)
		os.Exit(1)
	}

	// Assert the final tea.Model to our local model and print the choice.
	if m, ok := m.(model); ok && m.choice != "" {
		return m.choice
	}

	// CleanStdout()
	return ""
}

func InitView() (string, []byte) {
	//ask for admin user creation
	fmt.Printf("Enter your username: ")

	username, err := reader.ReadString('\n')

	if len(strings.TrimSpace(username)) == 0 {
		err = fmt.Errorf("Your Username can't be empty %v", username)
		fmt.Println(err.Error())
		os.Exit(1)
	}
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Enter your password: ")
	password, err := go_asterisks.GetUsersPassword("", true, os.Stdin, os.Stdout)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Printf("\n\n### You registered successfully! :) ###")

	// CleanStdout()
	return username, password
}

func AddTaskView() string {
	fmt.Printf("Enter your todo: ")

	task, err := reader.ReadString('\n')
	if len(strings.TrimSpace(task)) == 0 {
		err = fmt.Errorf("Your Username can't be empty %v", task)
		fmt.Println(err.Error())
		os.Exit(1)
	}
	if err != nil {
		log.Fatal(err)
	}

	// CleanStdout()
	return task
}

func CleanStdout() {
	time.Sleep(3 * time.Second)
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
