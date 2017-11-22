package service

import (
	"net/http"

	"github.com/unrolled/render"
)

func homeHandler(formatter *render.Render) http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		formatter.HTML(w, http.StatusOK, "index", struct {
			ID      string `json:"id"`
			Name     string `json:"name"`
			Content string `json:"content"`

		}{ID: "15331176", Name:"liziqiao", Content: "welcome to Go!"})
	}
}
