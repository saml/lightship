package app

import (
	"net/http"
)

import (
	"github.com/gorilla/mux"
)

type WebApp struct {
	router *mux.Router
}

func (app *WebApp) Url(name string, args ...string) string {
	url, err := app.router.Get(name).URL(args...)
	if err != nil {
		return ""
	}
	return url.String()
}

func (app *WebApp) Route(path string, handler http.HandlerFunc, name string) *mux.Route {
	return app.router.HandleFunc(path, handler).Name(name)
}

var app = WebApp{router: mux.NewRouter()}

func Init() http.Handler {
	configureRoutes()
	return app.router
}
