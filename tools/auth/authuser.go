package auth

import (
	"database/sql"
	"log"
)

func UserAuth(db *sql.DB, userToken string) (bool, string){
	
	umCheck, cuError := db.Query("SELECT count(*), `username` FROM `linka_user_token` WHERE token = '" + userToken + "'")

	if cuError == nil {

		userCount, userName := 0, "0"

		for umCheck.Next() {
			umCheck.Scan(&userCount, &userName)
		}

		var userBool bool

		if userCount >= 1{
			userBool = true
		}

		return userBool, userName

	} else {
		log.Panicln(cuError)
	}

	return false, "0"

}