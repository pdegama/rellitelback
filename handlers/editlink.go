package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"linka/tools"
	"linka/tools/auth"
	"log"
	"net/http"
	"os"
)

func (m *Repository) EditLink(w http.ResponseWriter, r *http.Request) {

	type LinkData struct {
		LinkName    string `json:"l_name"`
		LinkDes     string `json:"l_des"`
		LinkType    string `json:"l_type"`
		PreExt      string `json:"pre_e"`
		PreImg      string `json:"pre_i"`
		PreCle      bool   `json:"pre_c"`
		ThuExt      string `json:"thu_e"`
		ThuImg      string `json:"thu_i"`
		ThuCle      bool   `json:"thu_c"`
		LinkVisible string `json:"l_visible"`
	}

	type Req struct {
		BrowserCode string   `json:"browser_code"`
		Token       string   `json:"token"`
		LinkSlug    string   `json:"link_slug"`
		LinkData    LinkData `json:"link_info"`
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

			//fmt.Println(pageId)

			pageInfoQ, err := m.DataBase.Query("SELECT `name`, `type`, `des`, `thuimg`, `preimg`, `active` FROM `linka_links` WHERE `pageid` = " + fmt.Sprint(pageId) + " AND `slug` = '" + rq.LinkSlug + "'")
			if err != nil {
				log.Println(err)
				jsonRes.Success = false
				json.NewEncoder(w).Encode(&jsonRes)
				return
			}

			linkName := ""
			linkDes := ""
			linkType := 1
			linkThu := ""
			linkPre := ""
			linkVisible := 1

			for pageInfoQ.Next() {
				pageInfoQ.Scan(&linkName, &linkType, &linkDes, &linkThu, &linkPre, &linkVisible)
			}

			//fmt.Println(linkName, linkType, linkDes, linkThu, linkPre)

			if rq.LinkData.LinkName != "" {
				linkName = rq.LinkData.LinkName
			}

			if rq.LinkData.LinkDes != "" {
				linkDes = rq.LinkData.LinkDes
			}

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
				//linkType = linkType
			}

			switch rq.LinkData.LinkVisible{
			case "1":
				linkVisible = 1
			case "2":
				linkVisible = 2
			case "3":
				linkVisible = 3
			default:
				//linkVisible = linkVisible
			}

			if rq.LinkData.ThuCle {
				if linkThu != "" {
					//fmt.Println(m.StaticURL2 + linkThu)
					if _, err := os.Stat(m.StaticURL2 + linkThu); errors.Is(err, os.ErrNotExist) {
					} else {
						os.Remove(m.StaticURL2 + linkThu)
					}
				}
				linkThu = ""
			} else if rq.LinkData.ThuExt != "" {
				tools.UploadImage(rq.LinkData.ThuImg, m.StaticURL2+"thu_"+rq.LinkSlug+"."+rq.LinkData.ThuExt)
				linkThu = "thu_" + rq.LinkSlug + "." + rq.LinkData.ThuExt
			}

			if rq.LinkData.PreCle {
				if linkPre != "" {
					//fmt.Println(m.StaticURL2 + linkPre)
					if _, err := os.Stat(m.StaticURL2 + linkPre); errors.Is(err, os.ErrNotExist) {
					} else {
						os.Remove(m.StaticURL2 + linkPre)
					}
				}
				linkPre = ""
			} else if rq.LinkData.PreExt != "" {
				tools.UploadImage(rq.LinkData.PreImg, m.StaticURL2+"pre_"+rq.LinkSlug+"."+rq.LinkData.PreExt)
				linkPre = "pre_" + rq.LinkSlug + "." + rq.LinkData.PreExt
			}

			//fmt.Println(linkName, linkType, linkDes, linkThu, linkPre)

			linkEditQ, err := m.DataBase.Prepare("UPDATE `linka_links` SET `name`= ?,`type`= ?,`des`= ?,`thuimg`= ?,`preimg`= ?, `active` = ? WHERE `pageid` = ? AND `slug` = ?")
			if err != nil {
				log.Println(err)
				jsonRes.Success = false
				json.NewEncoder(w).Encode(&jsonRes)
				return
			}

			_, err = linkEditQ.Exec(linkName, linkType, linkDes, linkThu, linkPre, linkVisible, pageId, rq.LinkSlug)

			if err != nil {
				log.Println(err)
				jsonRes.Success = false
				json.NewEncoder(w).Encode(&jsonRes)
				return
			}

			jsonRes.Success = true
			log.Println("Link Update User: " + eUserName + " Page: " + fmt.Sprint(pageId) + " Slug: " + rq.LinkSlug)

		} else {

			jsonRes.Success = false

		}

		json.NewEncoder(w).Encode(&jsonRes)

	}

}
