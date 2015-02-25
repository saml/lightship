package main

import (
	"flag"
	"fmt"
	"net/http"
)

import (
	"github.com/saml/lightship/app"
)

func main() {
	port := flag.Int("port", 8080, "server port to listen")
	flag.Parse()

	http.Handle("/", app.Init())

	addr := fmt.Sprintf(":%d", *port)
	fmt.Printf("Listening http://localhost%s\n", addr)
	http.ListenAndServe(addr, nil)
}
