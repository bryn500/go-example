package main

import (
	"api/app"
	"api/config"
	"api/resource"
	"fmt"
	"net/http"
)

func main() {
	host, port := config.GetHostAndPort()
	server := app.Init()
	server.Get("/", home)
	server.AddRoutes(resource.TestRouter)
	server.AddRoutes(resource.TestRouterChain)
	server.Run(fmt.Sprintf("%s:%d", host, port))
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home: %s!", r.URL.Path[1:])
}