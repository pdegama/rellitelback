package handlers

import (
	"encoding/base32"
	"encoding/base64"
	"encoding/json"
	"linka/tools"
	"linka/tools/htmltemplates"
	"log"
	"net/http"
	"time"
)

func (m *Repository) ForgotPass(w http.ResponseWriter, r *http.Request) {

	type Req struct {
		BrowserCode string `json:"browser_code"`
		UserEmail   string `json:"user_email"`
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

		//fmt.Println(rq)

		countUser, err := m.DataBase.Query("SELECT COUNT(*) FROM `linka_users` WHERE `email` = '" + rq.UserEmail + "'")
		if err != nil {
			log.Println("User Check Error(Forgot Link Send 1) Email: ", rq.UserEmail)
			jsonRes.Success = false
			json.NewEncoder(w).Encode(&jsonRes)
			return
		}

		userCount := 0

		for countUser.Next() {
			countUser.Scan(&userCount)
		}

		if userCount != 0 {

			addCodeQ1, err := m.DataBase.Prepare("INSERT INTO `linka_forgot_code` VALUES (?, ?)")
			if err != nil {
				log.Println("User Check Error(Forgot Link Send 2) Email: ", rq.UserEmail)
				jsonRes.Success = false
				json.NewEncoder(w).Encode(&jsonRes)
				return
			}

			//make code
			originalCode := time.Now().String() + rq.UserEmail
			sEnc := base64.StdEncoding.EncodeToString([]byte(originalCode))
			sEnc = base32.StdEncoding.EncodeToString([]byte(sEnc))
			sEnc = tools.SpliteLastN(sEnc, 35)

			_, err = addCodeQ1.Exec(rq.UserEmail, sEnc)
			if err != nil {
				log.Println("Add Forgot Code Email: ", rq.UserEmail)
				jsonRes.Success = false
				json.NewEncoder(w).Encode(&jsonRes)
				return
			}

			tools.SendMail(rq.UserEmail, "Reset Password", htmltemplates.Forgot("", m.FrontURL+"fp/"+rq.UserEmail+"/"+sEnc))

			jsonRes.Success = true

		} else {
			jsonRes.Success = true
		}

		json.NewEncoder(w).Encode(&jsonRes)

	}

}
