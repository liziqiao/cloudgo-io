package service

import (
	"net/http"
	"github.com/unrolled/render"
)

func apiTestHandler(formatter *render.Render) http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		formatter.JSON(w, http.StatusOK, struct {
			ID      string `json:"id"`
			Name     string `json:"name"`
			Content string `json:"content"`
			
		}{ID: "15331176", Name: "liziqiao", Content: "welcome to Go!"})
	}
}
