package app

import (
	"encoding/json"
	"log"
	"net/http"
)

type App struct {
	Router *http.ServeMux
}

func Init() *App {
	mux := http.NewServeMux()
	return &App{Router: mux}
}

func (a *App) Run(host string) {
	log.Fatal(http.ListenAndServe(host, a.Router))
}

func JSONResponse(w http.ResponseWriter, r *http.Request, data interface{}) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	_, err = w.Write(jsonData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (a *App) AddRoutes(routerFunc func(*App)) {
	routerFunc(a)
}

func (a *App) Get(pattern string, handlers ...http.HandlerFunc) {
	handlerChain := chain(handlers)
	a.addRoute(pattern, http.MethodGet, handlerChain)
}

func (a *App) addRoute(pattern string, method string, handlerFunc http.HandlerFunc) {
	// Define a base handler function that checks the HTTP method.
	baseHandler := func(w http.ResponseWriter, r *http.Request) {
		if r.Method == method {
			handlerFunc(w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	}

	// Create a chain of middleware.
	middlewareChain := http.HandlerFunc(baseHandler)

 	// Add the securityHeadersMiddleware to the middleware chain.
	middlewareChain = securityHeadersMiddleware(middlewareChain)

	// Register the route with the middleware chain.
	a.Router.HandleFunc(pattern, middlewareChain)
}

func securityHeadersMiddleware(next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Security-Policy", "default-src 'self'")
        w.Header().Set("X-Frame-Options", "DENY")
        w.Header().Set("X-XSS-Protection", "0")
        
        // Call the next handler in the chain
        next(w, r)
    }
}

func chain(handlers []http.HandlerFunc) http.HandlerFunc {
	// Create a chain of middleware handlers.
	return func(w http.ResponseWriter, r *http.Request) {
		// Iterate through the list of handlers and call them in order.
		for _, handler := range handlers {
			handler(w, r)
		}
	}
}