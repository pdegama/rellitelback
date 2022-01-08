package tools

import (
	"encoding/json"
	"net/http"
)

func StorRequest(r *http.Request, i interface{}) {

	//r.ParseMultipartForm(2 << 20)

	json.NewDecoder(r.Body).Decode(&i)

}

func SpliteLastN(s string, n int) string {
	if len(s) > n {
		 return s[len(s)-n:]
	}
	return s
}