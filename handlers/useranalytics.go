package handlers

import (
	"encoding/json"
	"fmt"
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

	type evStruct struct {
		Date string  `json:"date"`
		Earn float32 `json:"earn"`
		View int     `json:"view"`
	}

	type linkD struct {
		Slug  string `json:"l_slug"`
		Name  string `json:"l_name"`
		Count int    `json:"l_count"`
	}

	type userNStr struct {
		Title string `json:"title"`
		URL   string `json:"url"`
	}

	type resStruct struct {
		ErrorCode     int32 `json:"error_code"`
		Success       bool  `json:"success"`
		UserAnalytics struct {
			UserBalance   float32    `json:"user_balance"`
			TodayEarn     float32    `json:"today_earn"`
			YesterdayEarn float32    `json:"yesterday_earn"`
			MonthEarn     float32    `json:"month_earn"`
			LastMEarn     float32    `json:"lastm_earn"`
			LastSeven     float32    `json:"last_7"`
			PrevSeven     float32    `json:"prev_7"`
			ChartData     []evStruct `json:"chart_data"`
			TopLink       []linkD    `json:"top_link"`
			PageSlug      string     `json:"p_slug"`
			UserNews      []userNStr `json:"user_news"`
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

			//user page
			uPageQ, err := m.DataBase.Query("SELECT `slug` FROM `linka_pages` WHERE `username` = '" + eUserName + "'")
			if err != nil {
				jsonRes.Success = false
				json.NewEncoder(w).Encode(&jsonRes)
				return
			}
			for uPageQ.Next() {
				uPageQ.Scan(&jsonRes.UserAnalytics.PageSlug)
			}
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
			lMStart := time.Date(time.Now().Year(), time.Now().Month()-1, 1, 0, 0, 0, 0, time.UTC).Format("2006-01-02") //start month time
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

			//get chart deta
			cData := make([]evStruct, 31)
			monthDay := tools.DaysIn(time.Now().Month(), time.Now().Year())
			dynDate := time.Now().AddDate(0, 0, -monthDay+1)
			var tEV evStruct
			for sInd := 0; sInd < monthDay; {
				//log.Println(dynDate.Format("2006-01-02"), dynDate.AddDate(0, 0, 1).Format("2006-01-02"))

				date1 := dynDate.Format("2006-01-02")
				date2 := dynDate.AddDate(0, 0, 1).Format("2006-01-02")

				evQ, err := m.DataBase.Query("SELECT COUNT(*), sum(`por`) FROM `linka_analytics` WHERE `username` = '" + eUserName + "' AND `time` >= '" + date1 + "' AND `time` <= '" + date2 + "'")
				if err != nil {
					jsonRes.Success = false
					json.NewEncoder(w).Encode(&jsonRes)
					return
				}

				tEV.Date = ""
				tEV.Earn = 0
				tEV.View = 0

				for evQ.Next() {
					tEV.Date = date1
					evQ.Scan(&tEV.View, &tEV.Earn)
				}

				cData[sInd] = tEV

				dynDate = dynDate.AddDate(0, 0, 1)
				sInd++
			}

			jsonRes.UserAnalytics.ChartData = cData

			//top links data
			tLiID, err := m.DataBase.Query("SELECT DISTINCT `linkid` FROM `linka_analytics` WHERE `username` = '" + eUserName + "' AND `time` >= '" + d7Start + "'")
			if err != nil {
				jsonRes.Success = false
				json.NewEncoder(w).Encode(&jsonRes)
				return
			}

			tLId1 := 0
			lCount := make(map[int]int)
			for tLiID.Next() {
				tLiID.Scan(&tLId1)
				lCount[tLId1] = 0
			}

			tLId2 := 0
			for lID := range lCount {
				tLiQ2, err := m.DataBase.Query("SELECT COUNT(`linkid`) FROM `linka_analytics` WHERE `username` = '" + eUserName + "' AND `linkid` = " + fmt.Sprint(lID) + " AND `time` >= '" + d7Start + "'")
				if err != nil {
					jsonRes.Success = false
					json.NewEncoder(w).Encode(&jsonRes)
					return
				}
				for tLiQ2.Next() {
					tLiQ2.Scan(&tLId2)
					lCount[lID] = tLId2
				}
			}

			lCOU := len(lCount)
			linkList := make([]linkD, lCOU)

			var tLId3 linkD
			tLId4 := 0
			for lID, lCO := range lCount {
				tLiQ3, err := m.DataBase.Query("SELECT `slug`, `name` FROM `linka_links` WHERE `id` = " + fmt.Sprint(lID) + " AND `username` = '" + eUserName + "'")
				if err != nil {
					jsonRes.Success = false
					json.NewEncoder(w).Encode(&jsonRes)
					return
				}
				for tLiQ3.Next() {
					tLId3.Count = lCO
					tLiQ3.Scan(&tLId3.Slug, &tLId3.Name)
				}

				linkList[tLId4] = tLId3

				tLId4++

			}

			jsonRes.UserAnalytics.TopLink = linkList

			//get user news
			userNewsQ, err := m.DataBase.Query("SELECT * FROM `linka_news` LIMIT 0, 10")
			if err != nil {
				jsonRes.Success = false
				json.NewEncoder(w).Encode(&jsonRes)
				return
			}
			
			var tNews userNStr
			jsonRes.UserAnalytics.UserNews = make([]userNStr, 10)
			tV4 := 0
			for userNewsQ.Next() {
				if tV4 == 10 {
					break
				}
				userNewsQ.Scan(&tNews.Title, &tNews.URL)
				jsonRes.UserAnalytics.UserNews[tV4] = tNews
				tV4 ++
			}

			jsonRes.Success = true

		} else {

			jsonRes.Success = false

		}

		json.NewEncoder(w).Encode(&jsonRes)

	}

}
