package app

import (
	"github.com/gorilla/mux"
)

func newRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/articles/", articleListHandler).Name("articles")
	return router
}
