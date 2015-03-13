// CMS Application
//
// Usage:
//	import "net/http"
//	import "lightship/app"
//	cms := app.Init()
//	http.ListenAndServe(":8080", cms.Handler())
package application

import (
	"net/http"
)

import (
	"github.com/gorilla/mux"
)

type App interface {
	// Given name of a handler and handler's path args, returns url (path).
	// In case of error, returns empty string.
	Url(name string, args ...string) string

	// Registers handler at path. Names the handler as well.
	HandleFunc(path string, handler http.HandlerFunc, name string) *mux.Route

	// Returns main http handler.
	Handler() http.Handler
}

type webApp struct {
	router *mux.Router
}

func (app *webApp) Url(name string, args ...string) string {
	url, err := app.router.Get(name).URL(args...)
	if err != nil {
		return ""
	}
	return url.String()
}

func (app *webApp) HandleFunc(path string, handler http.HandlerFunc, name string) *mux.Route {
	return app.router.HandleFunc(path, handler).Name(name)
}

func (app *webApp) Handler() http.Handler {
	return app.router
}

// Gives instance of App.
//
// TODO: takes configuration.
func New() App {
	app := new(webApp)
	app.router = mux.NewRouter()
	app.configureRoutes()
	return app
}
