package application

// Default routes.
func (app *webApp) configureRoutes() {
	app.HandleFunc("/articles/", app.articleListHandler, "articles")
}
