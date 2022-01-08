package handlers

import (
	"encoding/json"
	"linka/tools"
	"linka/tools/atools"
	"net/http"
)

func (m *Repository) ContantChecker(w http.ResponseWriter, r *http.Request) {

	type Req struct {
		BrowserCode string `json:"browser_code"`
		URI         string `json:"uri"`
	}

	var rq Req

	tools.StorRequest(r, &rq)

	SetHeaders(m, w)

	type resStruct struct {
		ErrorCode int32  `json:"error_code"`
		Exist     bool   `json:"exist"`
		Contant   string `json:"contant"`
	}

	if rq.BrowserCode == "" {

		json.NewEncoder(w).Encode(&struct {
			ErrorCode int32 `json:"error_code"`
		}{
			404,
		})

	} else {

		var jsonRes resStruct

		jsonRes.ErrorCode = 200

		ConType, ExistBool := atools.CChecker(rq.URI)
		jsonRes.Exist = ExistBool
		jsonRes.Contant = ConType

		json.NewEncoder(w).Encode(&jsonRes)

	}

}
