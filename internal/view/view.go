package view

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"

	userOp "github.com/epousa/todoListProject/internal/user"
	"gitlab.com/david_mbuvi/go_asterisks"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)

func InitView() (string, []byte) {
	//ask for admin user creation
	fmt.Println("Welcome! Create an admin account to start.")
	fmt.Printf("Enter your admin username: ")

	username, err := reader.ReadString('\n')

	if len(strings.TrimSpace(username)) == 0 {
		err = fmt.Errorf("Your Username can't be empty %v", username)
		fmt.Println(err.Error())
		os.Exit(1)
	}
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Enter your admin password: ")
	password, err := go_asterisks.GetUsersPassword("", true, os.Stdin, os.Stdout)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Printf("\n\n### You registered successfully! :) ###")

	return username, password
}

func CleanStdout() {
	time.Sleep(3 * time.Second)
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func AdminView(users *[]userOp.User) {
	//ask for admin user creation

}

func StandardView(user userOp.User) {
	//ask for admin user creation

}
