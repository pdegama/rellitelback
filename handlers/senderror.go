package handlers

import "net/http"

func (m *Repository) Error404(w http.ResponseWriter, r *http.Request){

	SetHeaders(m, w)

	w.Write([]byte("{\"error_code\":404}"))

}