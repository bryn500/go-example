package resource

import (
	"api/app"
	"net/http"
)

type TestType struct {
    A string `json:"a"`
    B string `json:"b"`
    C int    `json:"c"`
	D string `json:"-"`
}

func TestRouter(a *app.App) {
	middlewareChain := middleware1(middleware2(testHandler))
	a.Get("/test", middlewareChain)
}

func TestRouterChain(a *app.App) {
	a.Get("/test-chain", middlewareChain1, middlewareChain2, testHandler)
}

func middleware1(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-Bryn", "middleware1")
		next(w, r)
	}
}

func middleware2(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		next(w, r)
	}
}

func testHandler(w http.ResponseWriter, r *http.Request) {
	test := TestType{
		A: "John",
		B: "Doe",
		C: 30,
		D: "other",
	}

	app.JSONResponse(w, r, test)
}

func middlewareChain1(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("X-Bryn", "middlewareChain1")
}

func middlewareChain2(w http.ResponseWriter, r *http.Request) {
}
