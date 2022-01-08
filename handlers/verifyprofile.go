package handlers

import (
	"encoding/json"
	"linka/tools"
	"linka/tools/auth"
	"linka/tools/htmltemplates"
	"log"
	"net/http"
)

func (m *Repository) VerifyProfile(w http.ResponseWriter, r *http.Request) {

	type Req struct {
		BrowserCode string `json:"browser_code"`
		Token       string `json:"token"`
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

		uEBool, eUserName := auth.UserAuth(m.DataBase, rq.Token)

		if uEBool {

			var uEmail, uVCode string
			vCodeQ, vCodeError := m.DataBase.Query("SELECT * FROM `linka_verify_code` WHERE email = (SELECT `email` FROM `linka_users` WHERE `username` = '" + eUserName + "')")
			if vCodeError == nil {

				for vCodeQ.Next() {

					vCodeQ.Scan(&uEmail, &uVCode)

				}

			} else {

				log.Println("User Profile Not Metch:", eUserName, vCodeError)
				json.NewEncoder(w).Encode(&jsonRes)
				return
			}

			tools.SendMail(uEmail, "Verify Your Email Address", htmltemplates.EmailVarify(eUserName, "", m.FrontURL + "verify/" + uEmail + "/" + uVCode))

			//fmt.Println(uEmail, uVCode)
			jsonRes.Success = true
		} else {
			jsonRes.Success = false
		}

		json.NewEncoder(w).Encode(&jsonRes)

	}

}
