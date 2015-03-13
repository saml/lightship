package application

import (
	"net/http"
)

func (app *webApp) articleListHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(app.Url("articles")))
}
