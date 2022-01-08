package handlers

import (
	"encoding/json"
	"linka/tools"
	"linka/tools/auth"
	"net/http"
)

func (m *Repository) MainAuth(w http.ResponseWriter, r *http.Request) {

	type Req struct {
		BrowserCode string `json:"browser_code"`
		Token    string `json:"token"`
	}

	var rq Req

	tools.StorRequest(r, &rq)

	SetHeaders(m, w)

	type resStruct struct {
		ErrorCode int32 `json:"error_code"`
		Auth      bool  `json:"user_auth"`
	}

	if rq.BrowserCode == "" {
		json.NewEncoder(w).Encode(&struct {
			ErrorCode int32 `json:"error_code"`
		}{
			404,
		})
	} else {

		var jsonRes resStruct

		userBool, _ := auth.UserAuth(m.DataBase, rq.Token)

		if userBool {
			jsonRes.Auth = true
			jsonRes.ErrorCode = 200
		}

		json.NewEncoder(w).Encode(&jsonRes)
	}

}
