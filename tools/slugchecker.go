package tools

import (
	"database/sql"
	"fmt"
)

func SlugChecker(db *sql.DB, pageId int, linkSlug string) bool{
	que1, err := db.Query("SELECT COUNT(*) FROM `linka_links` WHERE `slug` = '" + linkSlug + "' AND `pageid` = '" + fmt.Sprint(pageId) + "'")

	if err != nil {
		fmt.Println(err)
		return false
	}

	var n int
	for que1.Next() {
		que1.Scan(&n)
	}

	if n == 0 {
		return false
	} else {
		return true
	}

}