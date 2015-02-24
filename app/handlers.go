package app

import (
	"net/http"
)

func articleListHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(app.Url("articles")))
}
