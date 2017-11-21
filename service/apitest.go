package service

import (
	"net/http"
	"github.com/unrolled/render"
)

func apiTestHandler(formatter *render.Render) http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		formatter.JSON(w, http.StatusOK, struct {
			ID      string `json:"id"`
			AGE     string `json:"age"`
			Content string `json:"content"`
			
		}{ID: "15331151", AGE:"20", Content: "Nice to meet you!"})
	}
}
