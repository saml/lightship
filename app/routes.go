package app

func configureRoutes() {
	app.HandleFunc("/articles/", articleListHandler, "articles")
}
