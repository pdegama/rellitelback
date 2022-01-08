package serverstart

import (
	"database/sql"
	"fmt"
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
	
	//server start
	file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	log.SetOutput(file)

	sqliteDatabase, dbError := sql.Open("mysql", "root:@/linka")

	if dbError != nil {
		log.Println(dbError)
	}

	defer sqliteDatabase.Close()

	AppHeaders := make(map[string]string)
	AppHeaders["Content-Type"] = "application/json"
	AppHeaders["Access-Control-Allow-Origin"] = "*"
	AppHeaders["Access-Control-Allow-Headers"] = "*"

	handlers.Repo.DataBase = sqliteDatabase
	handlers.Repo.Headers = &AppHeaders
	handlers.Repo.SecureCookie = secureCookie

	handlers.Repo.StaticURL = "/home/parthka/Project/linkyflix-master/linka-front/public/pagelogo/"
	handlers.Repo.StaticURL2 = "/home/parthka/Project/linkyflix-master/linka-front/public/images/"
	handlers.Repo.FrontURL = "http://app.rellitel.ink/"
	handlers.Repo.SecKey = "17813736645r73er8t4w2c8t"
	
	/* 
	orig := "Parth@123"
    fmt.Println("Original text:", orig)

    encryptCode := tools.AesEncrypt(orig, handlers.Repo.SecKey)
    fmt.Println("Ciphertext:" , encryptCode)
	*/
    
	 /* 
	decryptCode := tools.AesDecrypt("RIR8M6ZEELyxNsmgvk8Edw", handlers.Repo.SecKey)
    fmt.Println("Decryption result:", decryptCode)
	*/

	rh := mux.NewRouter()
	router.HTTPRouter(rh)

	//server start
	log.Println(fmt.Sprintf("Staring application on port %s", ListenAddress))
	_ = http.ListenAndServe(ListenAddress, rh)

}