package handlers

import (
	"encoding/json"
	"fmt"
	"linka/tools"
	"log"
	"net/http"
)

func (m *Repository) ResetPass(w http.ResponseWriter, r *http.Request) {

	type Req struct {
		BrowserCode string `json:"browser_code"`
		UserEmail   string `json:"user_email"`
		UserCode    string `json:"user_code"`
		UserPass    string `json:"user_pass"`
	}

	var rq Req

	tools.StorRequest(r, &rq)

	SetHeaders(m, w)

	type resStruct struct {
		ErrorCode int32 `json:"error_code"`
		Success   bool  `json:"success"`
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

		if rq.UserEmail != "" {
			//fmt.Println(rq)

			countUser, err := m.DataBase.Query("SELECT COUNT(*) FROM `linka_forgot_code` WHERE `email` = '" + rq.UserEmail + "' AND `code` = '" + rq.UserCode + "'")
			if err != nil {
				log.Println("User Code Check Error(Forgot Link Send 1) Email: ", rq.UserEmail)
				jsonRes.Success = false
				json.NewEncoder(w).Encode(&jsonRes)
				return
			}

			userCount := 0

			for countUser.Next() {
				countUser.Scan(&userCount)
			}

			//fmt.Println(userCount)

			if userCount != 0 && rq.UserPass != ""{

				setNewPass, err := m.DataBase.Prepare("UPDATE `linka_users` SET `password` = ? WHERE `email` = ?")

				if err != nil {

					log.Println("New Pass Set Error(Reset 1) User: ", rq.UserEmail, err)
					jsonRes.Success = false
					json.NewEncoder(w).Encode(&jsonRes)
					return

				} else {
					//new password enc
					encPass := tools.AesEncrypt(rq.UserPass, m.SecKey)
					_, err = setNewPass.Exec(encPass, rq.UserEmail)

					if err != nil {

						fmt.Println("New Pass Set Error(Reset 2) User: ", rq.UserEmail, err)
						jsonRes.Success = false
						json.NewEncoder(w).Encode(&jsonRes)
						return

					}

					jsonRes.Success = true
					log.Println("User Password Update:", rq.UserEmail)
				}

				jsonRes.Success = true

			} else {

				jsonRes.Success = false

			}

			deleteCode, err := m.DataBase.Prepare("DELETE FROM `linka_forgot_code` WHERE `email` = ?")
			if err != nil {
				log.Panicln("Forgot Code Delete error (1) User: ", rq.UserEmail, err)
			} else {
				_, err = deleteCode.Exec(rq.UserEmail)
				if err != nil {
					log.Panicln("Forgot Code Delete error (1) User: ", rq.UserEmail, err)
				}
			}
		}

		json.NewEncoder(w).Encode(&jsonRes)

	}

}
