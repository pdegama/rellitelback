package handlers

import (
	"encoding/json"
	"fmt"
	"linka/tools"
	"linka/tools/auth"
	"net/http"
)

func (m *Repository) ProfilePhoto(w http.ResponseWriter, r *http.Request) {

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

				fmt.Println(passError)

			}

			//enc pass
			if uPass == rq.UserNewInfo.CurPass {

				setNewPass, setNewPassError := m.DataBase.Prepare("UPDATE `linka_users` SET `password` = ? WHERE `username` = ?")

				if setNewPassError == nil {

					setNewPass.Exec(rq.UserNewInfo.NewPass, eUserName)
					jsonRes.Success = true

				} else {
					fmt.Println(setNewPassError)
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
