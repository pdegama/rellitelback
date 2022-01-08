package handlers

import (
	"encoding/json"
	"fmt"
	"linka/tools"
	"linka/tools/auth"
	"net/http"
)

func (m *Repository) UserProfile(w http.ResponseWriter, r *http.Request) {

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
		UserInfo  struct {
			UserName   string `json:"user_name"`
			UserFull   string `json:"user_fullname"`
			UserEmail  string `json:"user_email"`
			UserPhone  string `json:"user_phone"`
			UserPhoto  string `json:"user_photo"`
			UserVerify int    `json:"user_verify"`
		} `json:"user_info"`
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

		userE, userName := auth.UserAuth(m.DataBase, rq.Token)

		if userE {

			profileD, profileError := m.DataBase.Query("SELECT `username`, `email`, `fullname`, `verify`, `phone`, `photo` FROM `linka_users` WHERE `username` = '" + userName + "'")
			if profileError == nil {

				var uName, uEmail, uFullName, uPhone, uPhoto string
				var uVerify int

				for profileD.Next() {

					profileD.Scan(&uName, &uEmail, &uFullName, &uVerify, &uPhone, &uPhoto)

				}

				jsonRes.UserInfo.UserName = uName
				jsonRes.UserInfo.UserEmail = uEmail
				jsonRes.UserInfo.UserFull = uFullName
				jsonRes.UserInfo.UserPhone = uPhone
				jsonRes.UserInfo.UserPhoto = uPhoto
				jsonRes.UserInfo.UserVerify = uVerify

				jsonRes.Success = true

			} else {

				fmt.Println(profileError)

			}

		} else {

			jsonRes.Success = false

		}

		json.NewEncoder(w).Encode(&jsonRes)

	}

}
