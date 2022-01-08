package handlers

import (
	"encoding/json"
	"linka/tools"
	"log"
	"net/http"
)

func (m *Repository) UserLogOut(w http.ResponseWriter, r *http.Request) {

	type Req struct {
		BrowserCode string `json:"browser_code"`
		Token       string `json:"token"`
	}

	var rq Req

	tools.StorRequest(r, &rq)

	SetHeaders(m, w)

	type resStruct struct {
		ErrorCode int32  `json:"error_code"`
		Success   bool   `json:"success"`
	}

	if rq.BrowserCode == "" {

		json.NewEncoder(w).Encode(&struct {
			ErrorCode int32 `json:"error_code"`
		}{
			404,
		})

	} else {	

		var jsonRes resStruct

		removeToknenSta, removeTokenErr := m.DataBase.Prepare("DELETE FROM `linka_user_token` WHERE `token` = ?")
		if removeTokenErr != nil {
			log.Println("removeTokenErr", removeTokenErr)
		}

		_, removeTokenErrX := removeToknenSta.Exec(rq.Token)

		if removeTokenErrX == nil {
			//set token response
			jsonRes.ErrorCode = 200
			jsonRes.Success = true
		} else {
			log.Println(removeTokenErrX)
		}

		json.NewEncoder(w).Encode(&jsonRes)

	}

}
