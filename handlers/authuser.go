package handlers

import (
	"encoding/json"
	"fmt"
	"linka/tools"
	"log"
	"net/http"
	"strings"
)

func (m *Repository) AuthUser(w http.ResponseWriter, r *http.Request) {

	type UserData struct {
		UserOrEmail string `json:"uoe"`
		UserPass    string `json:"pass"`
	}

	type Req struct {
		BrowserCode string   `json:"browser_code"`
		UserData    UserData `json:"userdata"`
	}

	var rq Req

	tools.StorRequest(r, &rq)

	rq.UserData.UserOrEmail = strings.ToLower(rq.UserData.UserOrEmail)

	SetHeaders(m, w)

	type resStruct struct {
		ErrorCode int32  `json:"error_code"`
		Success   bool   `json:"success"`
		Token     string `json:"token"`
	}

	if rq.BrowserCode == "" {

		json.NewEncoder(w).Encode(&struct {
			ErrorCode int32 `json:"error_code"`
		}{
			404,
		})

	} else {

		var jsonRes resStruct

		jsonRes.ErrorCode = 200

		//check exist user or not exist user
		checkUser, checkUserError := m.DataBase.Query("SELECT count(*) FROM `linka_users` where `username` = '" + rq.UserData.UserOrEmail + "' OR `email` = '" + rq.UserData.UserOrEmail + "'")

		if checkUserError != nil {

			fmt.Println(checkUser)

		}

		var userNumRow int

		for checkUser.Next() {

			checkUser.Scan(&userNumRow)

		}

		if userNumRow >= 1 {

			//check password

			checkPWD, checkPWDError := m.DataBase.Query("SELECT `username`, `password` FROM `linka_users` WHERE `username` = '" + rq.UserData.UserOrEmail + "' OR `email` = '" + rq.UserData.UserOrEmail + "'")

			if checkPWDError != nil {

				fmt.Println(checkPWD)

			}

			var orPass string
			var orUser string

			for checkPWD.Next() {

				checkPWD.Scan(&orUser, &orPass)

			}

			//enc or dec
			if orPass == tools.AesEncrypt(rq.UserData.UserPass, m.SecKey) {

				value := map[string]string{
					"token": orUser + orPass,
				}

				if encodedToken, eCrr := m.SecureCookie.Encode("cookie-name", value); eCrr == nil {

					//set token database
					setToknenSta, setTokenErr := m.DataBase.Prepare("INSERT INTO `linka_user_token` VALUES(?, ?)")

					if setTokenErr != nil {
						fmt.Println(setTokenErr)
					}

					_, setTokenErrX := setToknenSta.Exec(encodedToken, orUser)

					if setTokenErrX == nil {
						//set token response
						jsonRes.Token = encodedToken
					} else {
						fmt.Println(setTokenErrX)
					}

				} else {

					fmt.Println(eCrr)

				}

				jsonRes.Success = true
				log.Println("User Login", orUser)

			}

			//fmt.Println(orUser, orPass)

		}

		json.NewEncoder(w).Encode(&jsonRes)

	}

}
