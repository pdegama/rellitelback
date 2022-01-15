package handlers

import (
	"encoding/json"
	"fmt"
	"linka/tools"
	"linka/tools/auth"
	"log"
	"net/http"
	"time"
)

func (m *Repository) NewLink(w http.ResponseWriter, r *http.Request) {

	type LinkData struct {
		LinkName    string `json:"l_name"`
		LinkURI     string `json:"l_uri"`
		LinkDes     string `json:"l_des"`
		LinkType    string `json:"l_type"`
		PreExt      string `json:"pre_e"`
		PreImg      string `json:"pre_i"`
		ThuExt      string `json:"thu_e"`
		ThuImg      string `json:"thu_i"`
		LinkVisible string `json:"l_visible"`
		LinkCon     string `json:"l_con"`
	}

	type Req struct {
		BrowserCode string   `json:"browser_code"`
		Token       string   `json:"token"`
		LinkData    LinkData `json:"link_info"`
	}

	var rq Req

	tools.StorRequest(r, &rq)

	SetHeaders(m, w)

	type resStruct struct {
		ErrorCode int32  `json:"error_code"`
		Success   bool   `json:"success"`
		LinkSlug  string `json:"link_slug"`
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

			linkSlug := tools.RandomString(8)

			for {
				if tools.SlugChecker(m.DataBase, pageId, linkSlug) {
					linkSlug = tools.RandomString(8)
				} else {
					break
				}
			}

			var linkName string

			if rq.LinkData.LinkName == "" {
				linkName = "untitled"
			} else {
				linkName = rq.LinkData.LinkName
			}

			linkType := 1

			switch rq.LinkData.LinkType {
			case "1":
				linkType = 1
			case "2":
				linkType = 2
			case "3":
				linkType = 3
			case "4":
				linkType = 4
			default:
				linkType = 1
			}

			linkVisible := 1
			switch rq.LinkData.LinkVisible {
			case "1":
				linkVisible = 1
			case "2":
				linkVisible = 2
			case "3":
				linkVisible = 3
			default:
				linkVisible = 1
			}

			nowTime := time.Now()
			nowT := nowTime.Format("2006-01-02 15:04:05")

			thuImgName := ""
			if rq.LinkData.ThuExt != "" {

				thuImgName = "thu_" + fmt.Sprint(pageId) + "_" + linkSlug + "." + rq.LinkData.ThuExt
				tools.UploadImage(rq.LinkData.ThuImg, m.StaticURL2+thuImgName)

			}

			preImgName := ""
			if rq.LinkData.PreExt != "" {

				preImgName = "pre_" + fmt.Sprint(pageId) + "_" + linkSlug + "." + rq.LinkData.PreExt
				tools.UploadImage(rq.LinkData.PreImg, m.StaticURL2+preImgName)

			}

			inserQ, err := m.DataBase.Prepare("INSERT INTO `linka_links` VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")

			if err != nil {
				log.Println(err, "Link No Insert User: "+eUserName)
				jsonRes.Success = false
				json.NewEncoder(w).Encode(&jsonRes)
				return
			}

			_, err = inserQ.Exec(nil, linkSlug, rq.LinkData.LinkURI, linkName, pageId, eUserName, nowT, linkType, rq.LinkData.LinkDes, thuImgName, preImgName, linkVisible, rq.LinkData.LinkCon)

			if err != nil {
				log.Println(err, "Link No Insert User: "+eUserName)
				jsonRes.Success = false
				json.NewEncoder(w).Encode(&jsonRes)
				return
			}

			jsonRes.Success = true
			jsonRes.LinkSlug = pageSlug + "/" + linkSlug

			log.Println("New Link Create User: " + eUserName + " Page: " + fmt.Sprint(pageId) + " Slug: " + linkSlug)

		} else {
			jsonRes.Success = false
			jsonRes.LinkSlug = "a23d8bdk39"
		}

		json.NewEncoder(w).Encode(&jsonRes)

	}

}
