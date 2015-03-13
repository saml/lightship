package main

import (
	"flag"
	"fmt"
	"net/http"
)

import (
	"github.com/saml/lightship/application"
)

func main() {
	port := flag.Int("port", 8080, "server port to listen")
	flag.Parse()

	app := application.New()

	addr := fmt.Sprintf(":%d", *port)
	fmt.Printf("Listening http://localhost%s\n", addr)
	http.ListenAndServe(addr, app.Handler())
}
