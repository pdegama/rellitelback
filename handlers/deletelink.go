package handlers

import (
	"encoding/json"
	"fmt"
	"linka/tools"
	"linka/tools/auth"
	"log"
	"net/http"
)

func (m *Repository) DeleteLink(w http.ResponseWriter, r *http.Request) {

	type Req struct {
		BrowserCode string   `json:"browser_code"`
		Token       string   `json:"token"`
		LinkSlug    string   `json:"link_slug"`
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

		//fmt.Println(rq)

		var jsonRes resStruct
		jsonRes.ErrorCode = 200

		uEBool, eUserName := auth.UserAuth(m.DataBase, rq.Token)

		if uEBool {

			//fmt.Println(eUserName)

			pageIdQ, err := m.DataBase.Query("SELECT `id`, `slug` FROM `linka_pages` WHERE `username` = '" + eUserName + "'")

			if err != nil {
				log.Println(err)
				return
			}

			var pageId int
			var pageSlug string

			for pageIdQ.Next() {
				pageIdQ.Scan(&pageId, &pageSlug)
			}

			linkEditQ, err := m.DataBase.Prepare("UPDATE `linka_links` SET `active` = ? WHERE `pageid` = ? AND `slug` = ?")
			
			if err != nil {
				log.Println(err)
				jsonRes.Success = false
				json.NewEncoder(w).Encode(&jsonRes)
				return
			}

			_, err = linkEditQ.Exec(0, pageId, rq.LinkSlug)

			if err != nil {
				log.Println(err)
				jsonRes.Success = false
				json.NewEncoder(w).Encode(&jsonRes)
				return
			}

			jsonRes.Success = true
			log.Println("Link Delete User: " + eUserName + " Page: " + fmt.Sprint(pageId) + " Slug: " + rq.LinkSlug)

		} else {

			jsonRes.Success = false

		}

		json.NewEncoder(w).Encode(&jsonRes)

	}

}
