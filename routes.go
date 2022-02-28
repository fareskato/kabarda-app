package main

import (
	"github.com/go-chi/chi/v5"
	"net/http"
)

// routes contains all application routes and uses Kabarda Routes of type *chi.Mux
func (a *application) routes() *chi.Mux {
	// middlewares go first
	// remember me middleware
	a.use(a.Middleware.CheckRemember)

	// routes go here
	a.get("/", a.Handlers.Home)
	

	// static routes
	fileServer := http.FileServer(http.Dir("./public"))
	a.handle("/public/*", http.StripPrefix("/public", fileServer))
	// return the mux
	return a.App.Routes
}

///////////////////////////////////////////////
// Helpers functions to make the routing easier
///////////////////////////////////////////////

// get helper
func (a *application) get(path string, h http.HandlerFunc) {
	a.App.Routes.Get(path, h)
}

// post helper
func (a *application) post(path string, h http.HandlerFunc) {
	a.App.Routes.Post(path, h)
}

// patch helper
func (a *application) patch(path string, h http.HandlerFunc) {
	a.App.Routes.Patch(path, h)
}

// put helper
func (a *application) put(path string, h http.HandlerFunc) {
	a.App.Routes.Put(path, h)
}

// delete helper
func (a *application) delete(path string, h http.HandlerFunc) {
	a.App.Routes.Delete(path, h)
}

// handle helper
func (a *application) handle(path string, h http.Handler) {
	a.App.Routes.Handle(path, h)
}

// use helper: for middlewares
func (a *application) use(m ...func(handler http.Handler) http.Handler) {
	a.App.Routes.Use(m...)
}
