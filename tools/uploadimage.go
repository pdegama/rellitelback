package tools

import (
	b64 "encoding/base64"
	"log"
	"os"
)

func UploadImage(datax string, path string) bool {

	sDec, err := b64.StdEncoding.DecodeString(datax)
	if err != nil {
		log.Println(err)
	}

	//log.Println("asdasd")

	err = os.WriteFile(path, sDec, 0644)

	if err == nil {
		//log.Println("yes")
		return true
	} else {
		//log.Println("no")
		log.Println(err)
		return false
	}

}
