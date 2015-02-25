package app

func configureRoutes() {
	app.Route("/articles/", articleListHandler, "articles")
}
