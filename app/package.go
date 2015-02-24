package app

import (
	"github.com/gorilla/mux"
)

type WebApp struct {
	router *mux.Router
}

func newApp() WebApp {
	return WebApp{}
}

func (app WebApp) Url(name string, args ...string) string {
	url, err := app.router.Get(name).URL(args...)
	if err != nil {
		return ""
	}
	return url.String()
}

var app = newApp()

func Init() {
	app.router = newRouter()
}

func Router() *mux.Router {
	return app.router
}
