package handlers

import (
	"encoding/json"
	"linka/tools"
	"log"
	"net/http"
)

func (m *Repository) Verify(w http.ResponseWriter, r *http.Request) {

	type Req struct {
		BrowserCode string `json:"browser_code"`
		UserEmail   string `json:"user_email"`
		UserCode    string `json:"user_code"`
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

		//fmt.Println(uEmail, uVCode)

		emailVQ1, err := m.DataBase.Query("SELECT count(*) FROM `linka_verify_code` WHERE `email` = '" + rq.UserEmail + "' AND `code` = '" + rq.UserCode + "'")

		if err != nil {
			jsonRes.Success = false
			json.NewEncoder(w).Encode(&jsonRes)
			return
		}

		matchQ := 0
		for emailVQ1.Next() {
			emailVQ1.Scan(&matchQ)
		}

		//fmt.Println("SELECT count(*) FROM `linka_verify_code` WHERE `email` = '" + rq.UserEmail + "' AND `code` = '" + rq.UserCode + "'", matchQ)

		if matchQ == 1 {

			emailVQ2, err := m.DataBase.Prepare("UPDATE `linka_users` SET `verify`= ? WHERE `email` = ?")
			if err != nil {
				jsonRes.Success = false
				json.NewEncoder(w).Encode(&jsonRes)
				return
			}

			_, err = emailVQ2.Exec(1, rq.UserEmail)
			if err != nil {
				jsonRes.Success = false
				json.NewEncoder(w).Encode(&jsonRes)
				return
			}

			jsonRes.Success = true

			//delete vcode
			emailQ2, err := m.DataBase.Prepare("DELETE FROM `linka_verify_code` WHERE `email` = ?")
			if err != nil {
				log.Println("Delete VCode Error(s) Email: " + rq.UserEmail, err)
				jsonRes.Success = false
				json.NewEncoder(w).Encode(&jsonRes)
				return
			}

			_, err = emailQ2.Exec(rq.UserEmail)
			if err != nil {
				log.Println("Delete VCode Error Email: " + rq.UserEmail, err)
				jsonRes.Success = false
				json.NewEncoder(w).Encode(&jsonRes)
				return
			}

		} else {

			jsonRes.Success = false
		
		}


		json.NewEncoder(w).Encode(&jsonRes)

	}

}
