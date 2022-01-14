package tools

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"linka/config"
	"os/exec"

	"github.com/manifoldco/promptui"
)

func InstallApp() {

	var UserNameMain = "admin"
	var UserKeyMain = "CKFK9-QNGF2-D34FM-99QX2-8XC4K"

	out, err := exec.Command("clear").Output()
	if err != nil {
		fmt.Printf("%s", err)
	}
	output := string(out[:])
	fmt.Println(output)

	//serverstart.ServerStart(":8080")

	//user menu
	prompt := promptui.Select{
		Label: "Select Option",
		Items: []string{"Install", "About"},
	}

	_, userSelection, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	if userSelection == "Install" {

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

			fmt.Println("\nHost Details.")
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

			fmt.Println("\nDatabase Details.")
			prompt = promptui.Prompt{
				Label:   "DataBase Hostname",
				Default: "localhost",
			}

			userHostDetails1, err := prompt.Run()

			if err != nil {
				fmt.Printf("Prompt failed %v\n", err)
				return
			}

			prompt = promptui.Prompt{
				Label:   "DataBase Usrename",
				Default: "root",
			}

			userHostDetails2, err := prompt.Run()

			if err != nil {
				fmt.Printf("Prompt failed %v\n", err)
				return
			}

			prompt = promptui.Prompt{
				Label:   "DataBase Password",
				Default: "",
				Mask:    '*',
			}

			userHostDetails3, err := prompt.Run()

			if err != nil {
				fmt.Printf("Prompt failed %v\n", err)
				return
			}

			prompt = promptui.Prompt{
				Label:   "DataBase Name",
				Default: "",
			}

			userHostDetails4, err := prompt.Run()

			if err != nil {
				fmt.Printf("Prompt failed %v\n", err)
				return
			}

			fmt.Println("\nSecret Key.")
			secKeyV := func(input string) error {
				if len(input) != 24 {
					return errors.New("key must have 24 characters")
				}
				return nil
			}
			prompt = promptui.Prompt{
				Label:    "Key",
				Default:  "",
				Validate: secKeyV,
			}

			userHostDetails5, err := prompt.Run()

			if err != nil {
				fmt.Printf("Prompt failed %v\n", err)
				return
			}

			fmt.Println("\nFrontend Url.")
			prompt = promptui.Prompt{
				Label:   "Url",
				Default: "",
			}

			userHostDetails6, err := prompt.Run()

			if err != nil {
				fmt.Printf("Prompt failed %v\n", err)
				return
			}

			fmt.Println("\nPath Config.")
			prompt = promptui.Prompt{
				Label:   "PageLogo Path",
				Default: "",
			}

			userHostDetails7, err := prompt.Run()

			if err != nil {
				fmt.Printf("Prompt failed %v\n", err)
				return
			}

			prompt = promptui.Prompt{
				Label:   "Image Path",
				Default: "",
			}

			userHostDetails8, err := prompt.Run()

			if err != nil {
				fmt.Printf("Prompt failed %v\n", err)
				return
			}

			fmt.Println("\nHeader Config.")
			prompt = promptui.Prompt{
				Label:   "Access-Control-Allow-Origin",
				Default: "",
			}

			userHostDetails9, err := prompt.Run()

			if err != nil {
				fmt.Printf("Prompt failed %v\n", err)
				return
			}

			prompt = promptui.Prompt{
				Label:   "Access-Control-Allow-Headers",
				Default: "",
			}

			userHostDetails10, err := prompt.Run()

			if err != nil {
				fmt.Printf("Prompt failed %v\n", err)
				return
			}

			logFileV := func(input string) error {
				if len(input) == 0 {
					return errors.New("Require")
				}
				return nil
			}

			fmt.Println("\nLogs File Path.")
			prompt = promptui.Prompt{
				Label:   "Path",
				Default: "",
				Validate: logFileV,
			}

			userHostDetails11, err := prompt.Run()

			if err != nil {
				fmt.Printf("Prompt failed %v\n", err)
				return
			}

			//fmt.Println("Server Start On " + userHostValidate + ":" + userPostValidate)
			//serverstart.ServerStart(userHostValidate + ":" + userPostValidate)

			var mainConfig config.MainConfig

			mainConfig.UserName = userNameValidate
			mainConfig.UserKey = userKeyValidate
			mainConfig.Host.HostName = userHostValidate
			mainConfig.Host.HostPort = userPostValidate
			mainConfig.DataBase.Host = userHostDetails1
			mainConfig.DataBase.User = userHostDetails2
			mainConfig.DataBase.Pass = userHostDetails3
			mainConfig.DataBase.Data = userHostDetails4
			mainConfig.SecKey = userHostDetails5
			mainConfig.FrontURL = userHostDetails6
			mainConfig.StaticPath = userHostDetails7
			mainConfig.StaticPath2 = userHostDetails8
			mainConfig.ACAO = userHostDetails9
			mainConfig.ACAH = userHostDetails10
			mainConfig.LogFile = userHostDetails11

			conFile, _ := json.MarshalIndent(mainConfig, "", " ")
			_ = ioutil.WriteFile("/etc/rellitel.ink/main.json", conFile, 0644)

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
