package view

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"gitlab.com/david_mbuvi/go_asterisks"
)

func InitView() (string, []byte) {
	//ask for admin user creation
	fmt.Println("Welcome! Create an admin account to start.")
	fmt.Printf("Enter your admin username: ")

	reader := bufio.NewReader(os.Stdin)
	username, err := reader.ReadString('\n')

	if len(strings.TrimSpace(username)) == 0 {
		err = fmt.Errorf("Your FirstName can't be empty %v", username)
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

	return username, password
}
