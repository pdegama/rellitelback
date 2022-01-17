package handlers

import (
	"encoding/base32"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"linka/tools"
	"log"
	"net/http"
	"strings"
	"time"
)

func (m *Repository) NewUser(w http.ResponseWriter, r *http.Request) {

	type UserData struct {
		UserName  string `json:"username"`
		UserEmail string `json:"email"`
		UserPass  string `json:"pass"`
		UserPage  string `json:"page"`
	}

	type Req struct {
		BrowserCode string   `json:"browser_code"`
		UserData    UserData `json:"userdata"`
	}

	var rq Req

	tools.StorRequest(r, &rq)

	SetHeaders(m, w)

	type resStruct struct {
		ErrorCode int32 `json:"error_code"`
		Already   struct {
			UserName  bool `json:"username"`
			UserEmail bool `json:"email"`
		} `json:"already_exiest"`
		Success bool `json:"success"`
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

		rq.UserData.UserName = strings.ToLower(rq.UserData.UserName)
		rq.UserData.UserEmail = strings.ToLower(rq.UserData.UserEmail)

		//check username
		cuCheck, cuError := m.DataBase.Query("SELECT count(*) FROM `linka_users` WHERE `username` = '" + rq.UserData.UserName + "'")
		if cuError != nil {
			fmt.Println(cuError)
		}

		defer cuCheck.Close()

		var existUsername int

		for cuCheck.Next() {
			cuCheck.Scan(&existUsername)
		}

		//log.Println(existUsername)

		if existUsername != 0 {
			jsonRes.Already.UserName = true
		}

		//email username
		umCheck, cuError := m.DataBase.Query("SELECT count(*) FROM `linka_users` WHERE `email` = '" + rq.UserData.UserEmail + "'")
		if cuError != nil {
			fmt.Println(cuError)
		}

		defer umCheck.Close()

		var existUseremail int

		for umCheck.Next() {
			umCheck.Scan(&existUseremail)
		}

		if existUseremail != 0 {
			jsonRes.Already.UserEmail = true
		}

		//if username and user email is exist that run condition
		if !jsonRes.Already.UserName && !jsonRes.Already.UserEmail {

			nowTime := time.Now()
			nowT := nowTime.Format("2006-01-02 15:04:05")

			//add new user
			newSqlStat, newUserAddError := m.DataBase.Prepare("INSERT INTO `linka_users` VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?)")

			if newUserAddError != nil {
				fmt.Println(newUserAddError)
			}

			//enc pass
			encPass := tools.AesEncrypt(rq.UserData.UserPass, m.SecKey)
			//fmt.Println(encPass)

			_, newUserAddError = newSqlStat.Exec(nil, rq.UserData.UserName, rq.UserData.UserEmail, encPass, rq.UserData.UserName, 0, nil, nil, nowT)

			if newUserAddError == nil {

				jsonRes.Success = true
				log.Println("User Register User: ", rq.UserData.UserName)

			} else {

				log.Println(newUserAddError)

			}

			//add page
			addPage := m.NewPage(rq.UserData.UserName, rq.UserData.UserPage)

			if !addPage {
				log.Println("User Page Not Add User: " + rq.UserData.UserPage)
			} else {
				log.Println("User Page Add User:", rq.UserData.UserName)
			}

			//add wallet
			newWallet, err := m.DataBase.Prepare("INSERT INTO `linka_wallet` VALUES(?, ?, ?)")
			if err != nil {
				log.Println("Add Wallet Error User (1): "+rq.UserData.UserName, err)
			}

			_, err = newWallet.Exec(nil, rq.UserData.UserName, 0.00)
			if err != nil {
				log.Println("Add Wallet Error User (2): "+rq.UserData.UserName, err)
			} else {
				log.Println("Add Wallet User: "+rq.UserData.UserName)
			}

			//delete vcode
			emailQ2, err := m.DataBase.Prepare("DELETE FROM `linka_verify_code` WHERE `email` = ?")
			if err != nil {
				log.Println("Delete VCode Error(s) Email: "+rq.UserData.UserEmail, err)
			}

			_, err = emailQ2.Exec(rq.UserData.UserEmail)
			if err != nil {
				log.Println("Delete VCode Error Email: " + rq.UserData.UserEmail)
			}

			//add verify code
			newVCode, vCodeError := m.DataBase.Prepare("INSERT INTO `linka_verify_code` VALUES(?, ?)")

			if vCodeError != nil {
				log.Println(vCodeError)
			}

			originalCode := time.Now().String() + rq.UserData.UserEmail
			sEnc := base64.StdEncoding.EncodeToString([]byte(originalCode))
			sEnc = base32.StdEncoding.EncodeToString([]byte(sEnc))
			sEnc = tools.SpliteLastN(sEnc, 35)

			_, newCodeError := newVCode.Exec(rq.UserData.UserEmail, sEnc)

			if newCodeError != nil {
				log.Println(newCodeError)
				log.Println("User V Code Not Add User:", rq.UserData.UserName, newCodeError)
			} else {
				log.Println("User V Code Add User:", rq.UserData.UserName)
			}

		}

		json.NewEncoder(w).Encode(&jsonRes)
	}

}
