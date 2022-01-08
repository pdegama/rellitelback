package handlers

import (
	"encoding/json"
	"linka/tools"
	"linka/tools/auth"
	"log"
	"net/http"
	"strconv"
	"strings"

	_ "github.com/gosimple/slug"
)

func (m *Repository) UpdatePage(w http.ResponseWriter, r *http.Request) {

	type Req struct {
		BrowserCode string `json:"browser_code"`
		Token       string `json:"token"`
		PageInfo    struct {
			PageSlug  string `json:"pageslug_i"`
			PageName  string `json:"pagename_i"`
			PageLogo2 string `json:"pagelogo1_i"`
			PageLogo1 string `json:"pagelogo2_i"`
			PageDes   string `json:"pagedes_i"`
			Theme     string `json:"theme_i"`
			SEO       bool   `json:"seo_i"`
			SFB       string `json:"sfb_i"`
			SIG       string `json:"sig_i"`
			STG       string `json:"stg_i"`
			SYT       string `json:"syt_i"`
		} `json:"page_info"`
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
		Already   bool  `json:"already"`
	}

	if rq.BrowserCode == "" {

		json.NewEncoder(w).Encode(&struct {
			ErrorCode int32 `json:"error_code"`
		}{
			404,
		})

	} else {

		//fmt.Println(rq.PageInfo)

		var jsonRes resStruct

		jsonRes.ErrorCode = 200

		uEBool, eUserName := auth.UserAuth(m.DataBase, rq.Token)

		if uEBool {

			pageInfoF, err := m.DataBase.Query("SELECT `id`, `name`, `slug`, `info`, `theme` FROM `linka_pages` WHERE `username` = '" + eUserName + "'")
			if err == nil {

				type pnStruct struct {
					Id       int
					pageName string
					pageSlug string
					info     string
					theme    int
				}

				var pageInfo pnStruct

				for pageInfoF.Next() {
					pageInfoF.Scan(&pageInfo.Id, &pageInfo.pageName, &pageInfo.pageSlug, &pageInfo.info, &pageInfo.theme)
				}

				pageSlug := pageInfo.pageSlug
				//if slug is already exist
				if rq.PageInfo.PageSlug != "" {

					pageSlug = strings.ToLower(rq.PageInfo.PageSlug)

					if m.slugChecker(pageSlug) {

						jsonRes.Success = false
						jsonRes.Already = true

						json.NewEncoder(w).Encode(&jsonRes)
						return

					} else {

						jsonRes.Success = true
						jsonRes.Already = false

					}
				}

				//log.Println(pageInfo)

				var pInfoX pInfoD

				json.Unmarshal([]byte(pageInfo.info), &pInfoX)

				if rq.PageInfo.PageLogo1 != "" {
					tools.UploadImage(rq.PageInfo.PageLogo1, m.StaticURL+strconv.Itoa(pageInfo.Id)+"_1")
					pInfoX.Logo = strconv.Itoa(pageInfo.Id) + "_1"
				}
				if rq.PageInfo.PageLogo2 != "" {
					tools.UploadImage(rq.PageInfo.PageLogo2, m.StaticURL+strconv.Itoa(pageInfo.Id)+"_2")
					pInfoX.Logo2 = strconv.Itoa(pageInfo.Id) + "_2"
				}
				if rq.PageInfo.PageDes != "" {
					pInfoX.Des = rq.PageInfo.PageDes
				}
				if rq.PageInfo.SEO != pInfoX.SEO {
					pInfoX.SEO = rq.PageInfo.SEO
				}
				if rq.PageInfo.SYT != "" {
					pInfoX.Social.YT = rq.PageInfo.SYT
				}
				if rq.PageInfo.STG != "" {
					pInfoX.Social.TG = rq.PageInfo.STG
				}
				if rq.PageInfo.SFB != "" {
					pInfoX.Social.FB = rq.PageInfo.SFB
				}
				if rq.PageInfo.SIG != "" {
					pInfoX.Social.IG = rq.PageInfo.SIG
				}

				pageInfoJson, err := json.Marshal(pInfoX)

				if err != nil {
					log.Println(err)
				}

				//fmt.Println(string(b))
				

				pageTheme := pageInfo.theme

				if rq.PageInfo.Theme != "" {
					pageTheme, err = strconv.Atoi(rq.PageInfo.Theme)
					if err != nil {
						log.Println("Fack Page Edit Error User", eUserName, err)
						return
					}
				}

				pageName := pageInfo.pageName

				if rq.PageInfo.PageName != "" {
					pageName = rq.PageInfo.PageName
				}

				updatePage, err := m.DataBase.Prepare("UPDATE `linka_pages` SET `name` = ?, `slug` = ?, `info` = ?, `theme` = ? WHERE `username` = ? AND `id` = ?")
				if err != nil {

					log.Println("Page Not Update User", eUserName, err)
					return

				} else {

					_, err = updatePage.Exec(pageName, pageSlug, pageInfoJson, pageTheme, eUserName, pageInfo.Id)

					if err != nil {
						log.Println("Page Not Update User", eUserName, err)
						return
					} else {

						jsonRes.Success = true
						log.Println("Page Update User: ", eUserName)

					}

				}

			} else {

				log.Println(err, "page info not fetch")

			}

		} else {
			jsonRes.Success = false
		}

		json.NewEncoder(w).Encode(&jsonRes)

	}

}
