package handlers

import (
	"log"
	"strconv"
	"time"

	"github.com/gosimple/slug"
)

func (m *Repository) NewPage(uName string, pName string) bool {

	sName := slug.Make(pName)

	//log.Println(pName, sName, m.slugChecker(sName))
	mainName := sName
	l := 0
	for {
		if m.slugChecker(mainName) {
			l = l + 1
			mainName = sName + "-" + strconv.Itoa(l)
		} else {
			break
		}
	}

	newPage, err2 := m.DataBase.Prepare("INSERT INTO `linka_pages` VALUES (?, ?, ?, ?, ?, ?, ?, ?)")

	if err2 == nil {

		nowTime := time.Now()
		nowT := nowTime.Format("2006-01-02 15:04:05")
		_, err := newPage.Exec(nil, pName, mainName, uName, nowT, "{}", "{\"logo\":\"/assets/img/brand/g61.png\",\"logo2\":\"/assets/img/brand/g61.png\",\"des\":\"\",\"seo\":true,\"social\":{\"yt\":\"#\",\"tg\":\"#\",\"fb\":\"#\",\"ig\":\"#\"}}", 1)

		if err == nil {
			return true
		} else {
			log.Println(1, err)
			return false
		}

	} else {
		log.Println(err2)
		return false
	}

}

func (m *Repository) slugChecker(slug string) bool {

	res1, err := m.DataBase.Query("SELECT COUNT(*) FROM `linka_pages` WHERE `slug` = '" + slug + "'")

	if err == nil {

		var slugChecker int

		for res1.Next() {
			res1.Scan(&slugChecker)
		}

		if slugChecker > 0 {
			return true
		} else {
			return false
		}

	} else {

		return true

	}

}
