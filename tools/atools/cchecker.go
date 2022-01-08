package atools

import (
	"net/http"
	"strings"
)


func CChecker(url string) (string, bool) {


	r, e:= http.Get(url)

	if e == nil {
		return strings.Split(r.Header["Content-Type"][0], ";")[0], true
	}
	
	return "", false
	
}