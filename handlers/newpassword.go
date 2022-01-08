package handlers

import (
	"encoding/json"
	"linka/tools"
	"linka/tools/auth"
	"log"
	"net/http"
)

func (m *Repository) NewPassword(w http.ResponseWriter, r *http.Request) {

	type Req struct {
		BrowserCode string `json:"browser_code"`
		Token       string `json:"token"`
		UserNewInfo struct {
			CurPass string `json:"cureent_pass"`
			NewPass string `json:"new_pass"`
		} `json:"user_pass"`
	}

	var rq Req

	tools.StorRequest(r, &rq)

	SetHeaders(m, w)

	type resStruct struct {
		ErrorCode int32 `json:"error_code"`
		Success   bool  `json:"success"`
		Error     struct {
			UserPass bool `json:"user_pass"`
		} `json:"pass_error"`
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

		uEBool, eUserName := auth.UserAuth(m.DataBase, rq.Token)

		if uEBool {

			var uPass string

			passD, passError := m.DataBase.Query("SELECT `password` FROM `linka_users` WHERE `username` = '" + eUserName + "'")
			if passError == nil {

				for passD.Next() {

					passD.Scan(&uPass)

				}

			} else {

				log.Println(passError)

			}

			//enc or dse
			if uPass == tools.AesEncrypt(rq.UserNewInfo.CurPass, m.SecKey) {

				setNewPass, setNewPassError := m.DataBase.Prepare("UPDATE `linka_users` SET `password` = ? WHERE `username` = ?")

				if setNewPassError == nil {

					//new password enc
					encPass := tools.AesEncrypt(rq.UserNewInfo.NewPass, m.SecKey)
					_, err := setNewPass.Exec(encPass, eUserName)

					if err != nil {

						log.Println("New Pass Set Error(Reset 2) User: ", eUserName)
						jsonRes.Success = false
	
					} else {
						jsonRes.Success = true
					}

					log.Println("User Password Update:", eUserName)

				} else {
					log.Println("New Pass Set Error User: ", eUserName, setNewPassError)
					jsonRes.Success = false
				}

			} else {
				jsonRes.Success = false
				jsonRes.Error.UserPass = true
			}

		} else {
			jsonRes.Success = false
		}

		json.NewEncoder(w).Encode(&jsonRes)

	}

}
