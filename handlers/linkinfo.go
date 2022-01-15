package handlers

import (
	"encoding/json"
	"fmt"
	"linka/tools"
	"linka/tools/auth"
	"log"
	"net/http"
)

func (m *Repository) LinkInfo(w http.ResponseWriter, r *http.Request) {

	type Req struct {
		BrowserCode string `json:"browser_code"`
		Token       string `json:"token"`
		LinkSlug    string `json:"link_slug"`
	}

	type resStruct struct {
		ErrorCode int32 `json:"error_code"`
		Success   bool  `json:"success"`
		LinkInfo  struct {
			LinkCount   int    `json:"l_count"`
			LinkName    string `json:"l_name"`
			LinkURI     string `json:"l_uri"`
			LinkDes     string `json:"l_des"`
			LinkType    int    `json:"l_type"`
			PreImg      string `json:"pre_i"`
			ThuImg      string `json:"thu_i"`
			LinkVisible int    `json:"l_visible"`
			LinkTime    string `json:"l_time"`
			LinkCon     string `json:"l_con"`
		} `json:"link_info"`
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

			//fmt.Println(pageId, pageSlug)

			infoQ1, err := m.DataBase.Query("SELECT COUNT(*), `link`, `name`, `type`, `des`, `thuimg`, `preimg`, `active`, `time`, `con` FROM `linka_links` WHERE `username` = '" + eUserName + "' AND `pageid` = " + fmt.Sprint(pageId) + " AND `slug` = '" + rq.LinkSlug + "' AND `active` IN (1, 2, 3)")
			if err != nil {
				log.Println(err)
				jsonRes.Success = false
				json.NewEncoder(w).Encode(&jsonRes)
				return
			}

			for infoQ1.Next() {
				infoQ1.Scan(&jsonRes.LinkInfo.LinkCount, &jsonRes.LinkInfo.LinkURI, &jsonRes.LinkInfo.LinkName, &jsonRes.LinkInfo.LinkType, &jsonRes.LinkInfo.LinkDes, &jsonRes.LinkInfo.ThuImg, &jsonRes.LinkInfo.PreImg, &jsonRes.LinkInfo.LinkVisible, &jsonRes.LinkInfo.LinkTime, &jsonRes.LinkInfo.LinkCon)
			}

			jsonRes.Success = true

		} else {
			jsonRes.Success = false
		}

		json.NewEncoder(w).Encode(&jsonRes)

	}

	SetHeaders(m, w)

}
