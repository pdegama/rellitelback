package handlers

import (
	"encoding/json"
	"linka/tools"
	"linka/tools/auth"
	"net/http"
)

func (m *Repository) Wallet(w http.ResponseWriter, r *http.Request) {

	type Req struct {
		BrowserCode string `json:"browser_code"`
		Token       string `json:"token"`
	}

	var rq Req

	tools.StorRequest(r, &rq)

	SetHeaders(m, w)

	type withTab struct {
		Date   string `json:"date"`
		POR    string `json:"por"`
		POM    string `json:"pom"`
		POA    string `json:"poa"`
		Status string `json:"status"`
	}

	type resStruct struct {
		ErrorCode   int32     `json:"error_code"`
		Success     bool      `json:"success"`
		UserBalance float32   `json:"user_balance"`
		Panding     int       `json:"user_panding"`
		WithTable   []withTab `json:"user_withtable"`
		WithCon     int       `json:"user_withcon"`
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

		userE, eUserName := auth.UserAuth(m.DataBase, rq.Token)

		if userE {

			uBalQ, err := m.DataBase.Query("SELECT `balance` FROM `linka_wallet` WHERE `username` = '" + eUserName + "'")
			if err != nil {
				jsonRes.Success = false
				json.NewEncoder(w).Encode(&jsonRes)
				return
			}

			for uBalQ.Next() {
				uBalQ.Scan(&jsonRes.UserBalance)
			}

			uPanQ, err := m.DataBase.Query("SELECT SUM(`por`) FROM `linka_withdraw` WHERE `username` = '" + eUserName + "' AND `status` = 1")
			if err != nil {
				jsonRes.Success = false
				json.NewEncoder(w).Encode(&jsonRes)
				return
			}

			for uPanQ.Next() {
				uPanQ.Scan(&jsonRes.Panding)
			}

			getTQ, err := m.DataBase.Query("SELECT `date`, `por`, `pom`, `poa`, `status` FROM `linka_withdraw` WHERE `username` = '" + eUserName + "' ORDER BY `id` DESC LIMIT 0, 20")
			if err != nil {
				jsonRes.Success = false
				json.NewEncoder(w).Encode(&jsonRes)
				return
			}

			withTabData := make([]withTab, 20)
			var tJ withTab
			tI := 0
			for getTQ.Next() {
				getTQ.Scan(&tJ.Date, &tJ.POR, &tJ.POM, &tJ.POA, &tJ.Status)
				withTabData[tI] = tJ
				tI++
			}
			jsonRes.WithCon = tI
			jsonRes.WithTable = withTabData
			jsonRes.Success = true

		} else {

			jsonRes.Success = false

		}

		json.NewEncoder(w).Encode(&jsonRes)

	}

}
