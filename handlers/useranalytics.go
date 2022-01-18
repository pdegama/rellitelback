package handlers

import (
	"encoding/json"
	"linka/tools"
	"linka/tools/auth"
	"net/http"
	"time"
)

func (m *Repository) UserAnalytics(w http.ResponseWriter, r *http.Request) {

	type Req struct {
		BrowserCode string `json:"browser_code"`
		Token       string `json:"token"`
	}

	var rq Req

	tools.StorRequest(r, &rq)

	SetHeaders(m, w)

	type resStruct struct {
		ErrorCode     int32 `json:"error_code"`
		Success       bool  `json:"success"`
		UserAnalytics struct {
			UserBalance   float32 `json:"user_balance"`
			TodayEarn     float32 `json:"today_earn"`
			YesterdayEarn float32 `json:"yesterday_earn"`
			MonthEarn     float32 `json:"month_earn"`
			LastMEarn     float32 `json:"lastm_earn"`
			LastSeven     float32 `json:"last_7"`
			PrevSeven     float32 `json:"prev_7"`
			UserPhone     string  `json:"user_phone"`
			UserPhoto     string  `json:"user_photo"`
			UserVerify    int     `json:"user_verify"`
		} `json:"user_analytics"`
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

			//user balance
			uBal, err := m.DataBase.Query("SELECT `balance` FROM `linka_wallet` WHERE `username` = '" + eUserName + "'")

			if err != nil {
				jsonRes.Success = false
				json.NewEncoder(w).Encode(&jsonRes)
				return
			}

			for uBal.Next() {
				uBal.Scan(&jsonRes.UserAnalytics.UserBalance)
			}

			//today earning
			tT := time.Now().Format("2006-01-02") //today time
			tEarn, err := m.DataBase.Query("SELECT sum(`por`) FROM `linka_analytics` WHERE `username` = '" + eUserName + "' AND `time` >= '" + tT + "'")
			if err != nil {
				jsonRes.Success = false
				json.NewEncoder(w).Encode(&jsonRes)
				return
			}

			for tEarn.Next() {
				tEarn.Scan(&jsonRes.UserAnalytics.TodayEarn)
			}

			//yesterday earning
			yT := time.Now().AddDate(0, 0, -1).Format("2006-01-02") //yesterday time
			yEarn, err := m.DataBase.Query("SELECT sum(`por`) FROM `linka_analytics` WHERE `username` = '" + eUserName + "' AND `time` >= '" + yT + "' AND `time` <= '" + tT + "'")
			if err != nil {
				jsonRes.Success = false
				json.NewEncoder(w).Encode(&jsonRes)
				return
			}

			for yEarn.Next() {
				yEarn.Scan(&jsonRes.UserAnalytics.YesterdayEarn)
			}

			//month earning
			mStart := time.Date(time.Now().Year(), time.Now().Month(), 1, 0, 0, 0, 0, time.UTC).Format("2006-01-02") //start month time
			mEarn, err := m.DataBase.Query("SELECT sum(`por`) FROM `linka_analytics` WHERE `username` = '" + eUserName + "' AND `time` >= '" + mStart + "'")
			if err != nil {
				jsonRes.Success = false
				json.NewEncoder(w).Encode(&jsonRes)
				return
			}

			for mEarn.Next() {
				mEarn.Scan(&jsonRes.UserAnalytics.MonthEarn)
			}

			//last month earning
			lMStart := time.Date(time.Now().Year(), time.Now().Month() - 1, 1, 0, 0, 0, 0, time.UTC).Format("2006-01-02") //start month time
			lMEarn, err := m.DataBase.Query("SELECT sum(`por`) FROM `linka_analytics` WHERE `username` = '" + eUserName + "' AND `time` >= '" + lMStart + "' AND `time` <= '" + mStart + "'")
			if err != nil {
				jsonRes.Success = false
				json.NewEncoder(w).Encode(&jsonRes)
				return
			}

			for lMEarn.Next() {
				lMEarn.Scan(&jsonRes.UserAnalytics.LastMEarn)
			}

			//las 7 days earning
			d7Start := time.Now().AddDate(0, 0, -6).Format("2006-01-02") //7days start time
			d7Earn, err := m.DataBase.Query("SELECT sum(`por`) FROM `linka_analytics` WHERE `username` = '" + eUserName + "' AND `time` >= '" + d7Start + "'")
			if err != nil {
				jsonRes.Success = false
				json.NewEncoder(w).Encode(&jsonRes)
				return
			}

			for d7Earn.Next() {
				d7Earn.Scan(&jsonRes.UserAnalytics.LastSeven)
			}

			//previous 7 days earning\
			p7Start := time.Now().AddDate(0, 0, -7-6).Format("2006-01-02") //7days start time
			p7Earn, err := m.DataBase.Query("SELECT sum(`por`) FROM `linka_analytics` WHERE `username` = '" + eUserName + "' AND `time` >= '" + p7Start + "' AND `time` <= '" + d7Start + "'")
			if err != nil {
				jsonRes.Success = false
				json.NewEncoder(w).Encode(&jsonRes)
				return
			}

			for p7Earn.Next() {
				p7Earn.Scan(&jsonRes.UserAnalytics.PrevSeven)
			}

			jsonRes.Success = true
			/* profileD, profileError := m.DataBase.Query("SELECT `username`, `email`, `fullname`, `verify`, `phone`, `photo` FROM `linka_users` WHERE `username` = '" + userName + "'")
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


			} else {

				fmt.Println(profileError)

			}
			*/
		} else {

			jsonRes.Success = false

		}

		json.NewEncoder(w).Encode(&jsonRes)

	}

}
