package main

import (
	"errors"
	"fmt"
	"linka/serverstart"
	"os/exec"

	_ "github.com/go-sql-driver/mysql"
	"github.com/manifoldco/promptui"
)

var UserNameMain = "admin"
var UserKeyMain = "CKFK9-QNGF2-D34FM-99QX2-8XC4K"

// main is the main function
func main() {

	//user display
	out, err := exec.Command("clear").Output()
	if err != nil {
		fmt.Printf("%s", err)
	}
	output := string(out[:])
	fmt.Println(output)

	//user display
	out, err = exec.Command("figlet", "RellBack").Output()
	if err != nil {
		fmt.Printf("%s", err)
	}
	output = string(out[:])
	fmt.Println(output)

	serverstart.ServerStart(":8080")

	//user menu
	prompt := promptui.Select{
		Label: "Select Option",
		Items: []string{"Start", "About"},
	}

	_, userSelection, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	if userSelection == "Start" {

		userValidate := func(input string) error {
			if len(input) < 3 {
				return errors.New("username must have more than 3 characters")
			}
			return nil
		}

		prompt := promptui.Prompt{
			Label:    "Username",
			Validate: userValidate,
			Default:  "admin",
		}

		userNameValidate, err := prompt.Run()

		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}

		userValidate = func(input string) error {
			if len(input) != 29 {
				return errors.New("key must have 29 characters")
			}
			return nil
		}

		prompt = promptui.Prompt{
			Label:    "Key",
			Validate: userValidate,
			Mask:     '*',
		}

		userKeyValidate, err := prompt.Run()

		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}

		if userNameValidate == UserNameMain && userKeyValidate == UserKeyMain {

			prompt := promptui.Prompt{
				Label:   "Hostname",
				Default: "localhost",
			}

			userHostValidate, err := prompt.Run()

			if err != nil {
				fmt.Printf("Prompt failed %v\n", err)
				return
			}

			userValidate := func(input string) error {
				switch input {
					case "8000":
					case "8080":
					case "8888":
					case "3000":
					case "3030":
					case "3333":
					default:
						return errors.New("available ports is 8000, 8080, 8888, 3000, 3030 and 3333")	
				}
				return nil
			}

			prompt = promptui.Prompt{
				Label:    "Port",
				Validate: userValidate,
				Default:  "8080",
			}

			userPostValidate, err := prompt.Run()

			if err != nil {
				fmt.Printf("Prompt failed %v\n", err)
				return
			}

			fmt.Println("Server Start On " + userHostValidate + ":" + userPostValidate)
			serverstart.ServerStart(userHostValidate + ":" + userPostValidate)

		} else {
			fmt.Println("Username or key invalid.")
		}

	} else {
		fmt.Println("RellBack Is Rellitel.ink Backend")
		fmt.Println("RellBack Written In Go")
		fmt.Println("Author: The Rellitel.ink Author")
		fmt.Println("Version: 1.0 alpha")
	}

}
