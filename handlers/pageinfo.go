package handlers

import (
	"encoding/json"
	"linka/tools"
	"linka/tools/auth"
	"log"
	"net/http"
)

func (m *Repository) PageInfo(w http.ResponseWriter, r *http.Request) {

	type Req struct {
		BrowserCode string `json:"browser_code"`
		Token       string `json:"token"`
	}

	var rq Req

	tools.StorRequest(r, &rq)

	SetHeaders(m, w)

	type pInfoD struct {
		Logo   string `json:"logo"`
		Logo2  string `json:"logo2"`
		Des    string `json:"des"`
		SEO    bool   `json:"seo"`
		Social struct {
			YT string `json:"yt"`
			TG string `json:"tg"`
			FB string `json:"fb"`
			IG string `json:"ig"`
		} `json:"social"`
	}

	type resStruct struct {
		ErrorCode int32 `json:"error_code"`
		Success   bool  `json:"success"`
		PageInfo  struct {
			PageName string `json:"page_name"`
			PageSlug string `json:"page_slug"`
			Theme    int    `json:"page_theme"`
			Logo  string `json:"logo"`
			Logo2 string `json:"logo2"`
			Des   string `json:"des"`
			SEO   bool   `json:"seo"`
			SYT string `json:"s_yt"`
			STG string `json:"s_tg"`
			SFB string `json:"s_fb"`
			SIG string `json:"s_ig"`
		} `json:"page_info"`
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

			pageInfoF, err := m.DataBase.Query("SELECT `name`, `slug`, `info`, `theme` FROM `linka_pages` WHERE `username` = '" + eUserName + "'")
			if err == nil {

				type pnStruct struct {
					pageName string
					pageSlug string
					info     string
					theme    int
				}

				var pageInfo pnStruct

				for pageInfoF.Next() {
					pageInfoF.Scan(&pageInfo.pageName, &pageInfo.pageSlug, &pageInfo.info, &pageInfo.theme)
				}

				//log.Println(pageInfo)

				var pInfoX pInfoD

				json.Unmarshal([]byte(pageInfo.info), &pInfoX)

				jsonRes.PageInfo.Logo = pInfoX.Logo
				jsonRes.PageInfo.Logo2 = pInfoX.Logo2
				jsonRes.PageInfo.Des = pInfoX.Des
				jsonRes.PageInfo.SEO = pInfoX.SEO
				jsonRes.PageInfo.SFB = pInfoX.Social.FB
				jsonRes.PageInfo.SIG = pInfoX.Social.IG
				jsonRes.PageInfo.SYT = pInfoX.Social.YT
				jsonRes.PageInfo.STG = pInfoX.Social.TG
				jsonRes.PageInfo.PageName = pageInfo.pageName
				jsonRes.PageInfo.PageSlug = pageInfo.pageSlug
				jsonRes.PageInfo.Theme = pageInfo.theme

				jsonRes.Success = true

				//log.Println(pInfoX)

			} else {
				log.Println(err, "page info not fetch")
			}

		} else {
			jsonRes.Success = false
		}

		json.NewEncoder(w).Encode(&jsonRes)

	}

}
