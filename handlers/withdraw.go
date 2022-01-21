package handlers

import (
	"encoding/json"
	"linka/tools"
	"linka/tools/auth"
	"log"
	"net/http"
)

func (m *Repository) Withdraw(w http.ResponseWriter, r *http.Request) {

	type Req struct {
		BrowserCode string  `json:"browser_code"`
		Token       string  `json:"token"`
		Withdraw    float64 `json:"wdamount"`
		POM         string  `json:"paym"`
		POA         string  `json:"paya"`
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

		userE, eUserName := auth.UserAuth(m.DataBase, rq.Token)

		if userE {

			uBalQ, err := m.DataBase.Query("SELECT `balance` FROM `linka_wallet` WHERE `username` = '" + eUserName + "'")
			if err != nil {
				jsonRes.Success = false
				json.NewEncoder(w).Encode(&jsonRes)
				return
			}

			var UserBalance float64

			for uBalQ.Next() {
				uBalQ.Scan(&UserBalance)
			}

			//fmt.Println(rq, UserBalance)

			if UserBalance >= rq.Withdraw && rq.Withdraw != 0 {
				// add withdraw
				addWithQ, err := m.DataBase.Prepare("INSERT INTO `linka_withdraw` VALUES (?, ?, ?, ?, ?, ?, DEFAULT, ?)")
				if err != nil {
					jsonRes.Success = false
					json.NewEncoder(w).Encode(&jsonRes)
					return
				}

				_, err = addWithQ.Exec(nil, eUserName, UserBalance, rq.Withdraw, rq.POM, rq.POA, 1)
				if err != nil {
					log.Println("Withdraw Error User: ", eUserName)
					jsonRes.Success = false
					json.NewEncoder(w).Encode(&jsonRes)
					return
				}
				log.Println("Withdraw Add User: ", eUserName, " por: ", rq.Withdraw)

				//minus use walltel balance
				minWallQ, err := m.DataBase.Prepare("UPDATE `linka_wallet` SET `balance` = `balance` - ? WHERE `username` = ?")
				if err != nil {
					log.Println("Wallet Minus Error User: ", eUserName)
					jsonRes.Success = false
					json.NewEncoder(w).Encode(&jsonRes)
					return
				}

				_, err = minWallQ.Exec(rq.Withdraw, eUserName)
				if err != nil {
					log.Println("Wallet Minus Error User: ", eUserName)
					jsonRes.Success = false
					json.NewEncoder(w).Encode(&jsonRes)
					return
				}

			} else {
				jsonRes.Success = false
				json.NewEncoder(w).Encode(&jsonRes)
				return
			}

			jsonRes.Success = true

		} else {

			jsonRes.Success = false

		}

		json.NewEncoder(w).Encode(&jsonRes)

	}

}
