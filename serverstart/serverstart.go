package serverstart

import (
	"database/sql"
	"fmt"
	"linka/config"
	"linka/handlers"
	"linka/router"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/gorilla/securecookie"
)

var hashKey = []byte("helloworld")
var blockKey = []byte{176, 122, 95, 133, 55, 187, 179, 0, 6, 210, 42, 228, 122, 0, 246, 219, 200, 97, 76, 170, 107, 154, 34, 217, 3, 102, 164, 199, 26, 66, 70, 206}
var secureCookie = securecookie.New(hashKey, blockKey)

func ServerStart(ListenAddress string){
	
	sqliteDatabase, dbError := sql.Open("mysql", config.AppMainConfig.DataBase.User + ":" + config.AppMainConfig.DataBase.Pass + "@tcp(" + config.AppMainConfig.DataBase.Host + ":3306)/" + config.AppMainConfig.DataBase.Data)

	if dbError != nil {
		log.Println(dbError)
	}

	defer sqliteDatabase.Close()

	//server start
	file, err := os.OpenFile(config.AppMainConfig.LogFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	log.SetOutput(file)	

	AppHeaders := make(map[string]string)
	AppHeaders["Content-Type"] = "application/json"
	AppHeaders["Access-Control-Allow-Origin"] = config.AppMainConfig.ACAO
	AppHeaders["Access-Control-Allow-Headers"] = config.AppMainConfig.ACAH

	handlers.Repo.DataBase = sqliteDatabase
	handlers.Repo.Headers = &AppHeaders
	handlers.Repo.SecureCookie = secureCookie

	handlers.Repo.StaticURL = config.AppMainConfig.StaticPath
	handlers.Repo.StaticURL2 = config.AppMainConfig.StaticPath2
	handlers.Repo.FrontURL = config.AppMainConfig.FrontURL
	handlers.Repo.SecKey = config.AppMainConfig.SecKey

	//fmt.Println(config.AppMainConfig.StaticPath)
	//fmt.Println(config.AppMainConfig.StaticPath2)
	//fmt.Println(config.AppMainConfig.FrontURL)
	//fmt.Println(config.AppMainConfig.SecKey)

	rh := mux.NewRouter()
	router.HTTPRouter(rh)

	//server start
	log.Println(fmt.Sprintf("Staring application on port %s", ListenAddress))
	_ = http.ListenAndServe(ListenAddress, rh)

}