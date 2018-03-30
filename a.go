package main

import (
	"net/http"

	"github.com/unrolled/render"
)

func createMatch(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		formatter.JSON(w, http.StatusOK, struct{ Test string }{"some val"})
	}
}
