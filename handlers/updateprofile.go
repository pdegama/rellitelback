package handlers

import (
	"encoding/base32"
	"encoding/base64"
	"encoding/json"
	"linka/tools"
	"linka/tools/auth"
	"log"
	"net/http"
	"strings"
	"time"
)

func (m *Repository) UpdateProfile(w http.ResponseWriter, r *http.Request) {

	type Req struct {
		BrowserCode string `json:"browser_code"`
		Token       string `json:"token"`
		UserNewInfo struct {
			UEmail string `json:"uemail"`
			UName  string `json:"uname"`
			UPhone string `json:"uphone"`
		} `json:"user_new_info"`
	}

	var rq Req

	tools.StorRequest(r, &rq)

	SetHeaders(m, w)

	type resStruct struct {
		ErrorCode int32 `json:"error_code"`
		Success   bool  `json:"success"`
		Already   struct {
			UserEmail bool `json:"user_email"`
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

		uEBool, eUserName := auth.UserAuth(m.DataBase, rq.Token)

		if uEBool {

			var uEmail, uFullName, uPhone string
			var uVerify int

			profileD, profileError := m.DataBase.Query("SELECT `email`, `fullname`, `verify`, `phone` FROM `linka_users` WHERE `username` = '" + eUserName + "'")
			if profileError == nil {

				for profileD.Next() {

					profileD.Scan(&uEmail, &uFullName, &uVerify, &uPhone)

				}

			} else {

				log.Println("User Profile Not Metch:", eUserName, profileError)

			}

			//fmt.Println(uVerify)

			if rq.UserNewInfo.UEmail == "" {
				rq.UserNewInfo.UEmail = uEmail
			} else {
				rq.UserNewInfo.UEmail = strings.ToLower(rq.UserNewInfo.UEmail)
				uVerify = 0
			}

			if rq.UserNewInfo.UName == "" {
				rq.UserNewInfo.UName = uFullName
			} else {
				rq.UserNewInfo.UName = strings.ToLower(rq.UserNewInfo.UName)
			}

			if rq.UserNewInfo.UPhone == "" {
				rq.UserNewInfo.UPhone = uPhone
			} else {
				rq.UserNewInfo.UPhone = strings.ToLower(rq.UserNewInfo.UPhone)
			}

			//email username
			umCheck, cuError := m.DataBase.Query("SELECT count(*), `username` FROM `linka_users` WHERE `email` = '" + rq.UserNewInfo.UEmail + "'")
			//fmt.Println("SELECT count(*), `username` FROM `linka_users` WHERE `email` = '" + rq.UserNewInfo.UEmail + "'")
			if cuError != nil {
				log.Println("profile user check error: ", cuError)
			}

			defer umCheck.Close()

			var existUseremail int
			var existUserName string

			for umCheck.Next() {
				umCheck.Scan(&existUseremail, &existUserName)
			}

			tmpBool1 := true

			if existUserName == eUserName {
				tmpBool1 = false
			}

			if existUseremail != 0 && tmpBool1 {
				jsonRes.Already.UserEmail = true
			} else {

				uInfo, uInfoError := m.DataBase.Prepare("UPDATE `linka_users` SET `email` = ?, `phone` = ?, `fullname` = ?, `verify` = ? WHERE `username` = ?")
				if uInfoError == nil {

					_, uErro := uInfo.Exec(rq.UserNewInfo.UEmail, rq.UserNewInfo.UPhone, rq.UserNewInfo.UName, uVerify, eUserName)

					if uErro == nil {
						jsonRes.Success = true
						log.Println("User Profile Update:", eUserName)
					} else {
						log.Println("User Profile Update Err:", uErro)
						jsonRes.Success = false
					}

				} else {
					log.Println(uInfoError, rq.UserNewInfo.UEmail)
				}
				//fmt.Println(uMatch)

				if uVerify == 0 {

					//delete vcode
					emailQ2, err := m.DataBase.Prepare("DELETE FROM `linka_verify_code` WHERE `email` = ?")
					if err != nil {
						log.Println("Delete VCode (Update) Error(s) Email: "+rq.UserNewInfo.UEmail, err)
						jsonRes.Success = false
						json.NewEncoder(w).Encode(&jsonRes)
						return
					}

					_, err = emailQ2.Exec(rq.UserNewInfo.UEmail)
					if err != nil {
						log.Println("Delete VCode (Update) Error Email: "+rq.UserNewInfo.UEmail, err)
						jsonRes.Success = false
						json.NewEncoder(w).Encode(&jsonRes)
						return
					}

					//add vcode
					newVCode, vCodeError := m.DataBase.Prepare("INSERT INTO `linka_verify_code` VALUES(?, ?)")

					if vCodeError != nil {
						log.Println(vCodeError)
					}

					originalCode := time.Now().String() + rq.UserNewInfo.UEmail
					sEnc := base64.StdEncoding.EncodeToString([]byte(originalCode))
					sEnc = base32.StdEncoding.EncodeToString([]byte(sEnc))
					sEnc = tools.SpliteLastN(sEnc, 35)

					_, newCodeError := newVCode.Exec(rq.UserNewInfo.UEmail, sEnc)

					if newCodeError != nil {
						log.Println(newCodeError)
						log.Println("User V Code Not Add (Update)", rq.UserNewInfo.UEmail, newCodeError)
					} else {
						log.Println("User V Code Add (Update)", rq.UserNewInfo.UEmail)
					}
				}

			}

			//fmt.Println(uVerify)

		} else {
			jsonRes.Success = false
		}

		json.NewEncoder(w).Encode(&jsonRes)

	}

}
