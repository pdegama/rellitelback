package handlers

import (
	"encoding/json"
	"fmt"
	"linka/tools"
	"linka/tools/auth"
	"log"
	"math"
	"net/http"
)

func (m *Repository) AllLink(w http.ResponseWriter, r *http.Request) {

	type Req struct {
		BrowserCode string `json:"browser_code"`
		Token       string `json:"token"`
		LinksPage   int    `json:"links_page"`
		LinksQuery  string `json:"links_query"`
	}

	type LinkDataStruct struct {
		LinkName    string `json:"l_name"`
		LinkURI     string `json:"l_uri"`
		LinkSlug    string `json:"l_slug"`
		LinkType    int    `json:"l_type"`
		LinkTime    string `json:"l_time"`
		LinkVisible int    `json:"l_visible"`
		PageId      int    `json:"l_pid"`
	}

	type resStruct struct {
		ErrorCode  int32            `json:"error_code"`
		Success    bool             `json:"success"`
		TotalPage  int              `json:"total_page"`
		PageName   string           `json:"page_name"`
		LinksCount int              `json:"link_count"`
		LinksData  []LinkDataStruct `json:"links_data"`
	}

	var rq Req

	tools.StorRequest(r, &rq)

	SetHeaders(m, w)

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

			pageIdQ, err := m.DataBase.Query("SELECT `id`, `slug` FROM `linka_pages` WHERE `username` = '" + eUserName + "'")

			if err != nil {
				log.Println(err)
				jsonRes.Success = false
				json.NewEncoder(w).Encode(&jsonRes)
				return
			}

			var pageId int
			var pageSlug string

			for pageIdQ.Next() {
				pageIdQ.Scan(&pageId, &pageSlug)
			}

			//fmt.Println(eUserName, pageId, pageSlug, rq.LinksPage)
			//fmt.Println("SELECT `slug`, `link`, `name`, `type`, `time`, `active` FROM `linka_links` WHERE `username` = '" + eUserName + "' AND `pageid` = " + fmt.Sprint(pageId) + " AND `active` IN (1, 2) ORDER BY `id` DESC LIMIT " + fmt.Sprint((rq.LinksPage-1)*15) + ", 15")

			//var infoQ1 *sql.Rows
			var sqlQ1s string

			if rq.LinksQuery == "" {
				sqlQ1s = "FROM `linka_links` WHERE `username` = '" + eUserName + "' AND `pageid` = " + fmt.Sprint(pageId) + " AND `active` IN (1, 2, 3) ORDER BY `id` DESC"
			} else {
				sqlQ1s = "FROM `linka_links` WHERE `username` = '" + eUserName + "' AND `pageid` = " + fmt.Sprint(pageId) + " AND `active` IN (1, 2, 3) AND (`name` LIKE '%" + rq.LinksQuery + "%' OR `slug` LIKE '%" + rq.LinksQuery + "%') ORDER BY `id` DESC"
				//sqlQ1s = "FROM `linka_links` WHERE `username` = '" + eUserName + "' AND `pageid` = " + fmt.Sprint(pageId) + " AND `active` IN (1, 2, 3) AND (`link` LIKE '%" + rq.LinksQuery + "%' OR `name` LIKE '%" + rq.LinksQuery + "%' OR `slug` LIKE '%" + rq.LinksQuery + "%') ORDER BY `id` DESC"
			}

			//fmt.Println(sqlQ1s)
			if rq.LinksPage == 0 {
				rq.LinksPage = 1
			}

			infoQ1, err := m.DataBase.Query("SELECT `pageid`, `slug`, `link`, `name`, `type`, `time`, `active` " + sqlQ1s + " LIMIT " + fmt.Sprint((rq.LinksPage-1)*15) + ", 15")

			if err != nil {
				//log.Println(err)
				fmt.Println(err)
				return
			}

			i := 0
			j := 0

			var lTmpx LinkDataStruct
			Tmpxx := make([]LinkDataStruct, 15)

			for infoQ1.Next() {
				if i == 15 {
					break
				}
				j++
				infoQ1.Scan(&lTmpx.PageId, &lTmpx.LinkSlug, &lTmpx.LinkURI, &lTmpx.LinkName, &lTmpx.LinkType, &lTmpx.LinkTime, &lTmpx.LinkVisible)
				Tmpxx[i] = lTmpx
				i++
			}

			jsonRes.LinksData = Tmpxx
			jsonRes.LinksCount = j
			jsonRes.PageName = pageSlug

			infoQ2, err := m.DataBase.Query("SELECT COUNT(*)" + sqlQ1s)
			if err != nil {
				log.Println(err)
				return
			}

			p := 1

			for infoQ2.Next() {
				infoQ2.Scan(&p)
			}

			var totalPages float64

			totalPages = float64(p) / 15

			totalPages = math.Ceil(totalPages)

			jsonRes.TotalPage = int(totalPages)

			//fmt.Println(p, float64(p) / 15, float32(p), totalPages, int(totalPages))

			jsonRes.Success = true

		} else {
			jsonRes.Success = false
		}

		json.NewEncoder(w).Encode(&jsonRes)

	}

	SetHeaders(m, w)

}
