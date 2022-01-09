package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"linka/config"
	"linka/serverstart"
	"linka/tools"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

// main is the main function
func main() {

	if _, err := os.Stat("/etc/rellitel.ink/main.json"); err != nil {

		if os.IsNotExist(err) {
			tools.InstallApp()
		} else {
			fmt.Println("Install Error")
		}

	} else {

		//fmt.Println(tools.AesDecrypt("uHhMcWExFdJKzJRZiwjDHw", "17813736645r73er8t4w2c8t"))

		jsonFile, err := os.Open("/etc/rellitel.ink/main.json")

		byteValFile, _ := ioutil.ReadAll(jsonFile)

		if err != nil {
			fmt.Println(err)
		}

		json.Unmarshal(byteValFile, &config.AppMainConfig)

		var UserNameMain = "admin"
		var UserKeyMain = "CKFK9-QNGF2-D34FM-99QX2-8XC4K"
		if config.AppMainConfig.UserName == UserNameMain && config.AppMainConfig.UserKey == UserKeyMain {
			serverstart.ServerStart(config.AppMainConfig.Host.HostName + ":" + config.AppMainConfig.Host.HostPort)
		} else {
			os.Remove("/etc/rellitel.ink/main.json")
			fmt.Println("Reinstall App Please...")
		}

	}

}
