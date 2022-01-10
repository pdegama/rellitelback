package tools

import (
	"encoding/json"
	/* "fmt"
	"io/ioutil"
	"log" */
	"net/http"
)

func StorRequest(r *http.Request, i interface{}) {

	//r.ParseMultipartForm(2 << 20)

	/* var payload interface{} // The interface where we will save the converted JSON data.
    buffer, err := ioutil.ReadAll(r.Body)
    if err != nil {
        log.Panic(err)
    }
    r.Body.Close() // Close this
    json.Unmarshal(buffer, &payload) // Convert JSON data into interface{} type
    m := payload.(map[string]interface{})

	fmt.Println(m) */
	
	json.NewDecoder(r.Body).Decode(&i)
	
	

}

func SpliteLastN(s string, n int) string {
	if len(s) > n {
		return s[len(s)-n:]
	}
	return s
}
