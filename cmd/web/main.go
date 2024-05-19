package main

import (
	"fmt"
	"net/http"

	"github.com/codebyyogesh/lens_view/assets"
	"github.com/codebyyogesh/lens_view/internal/actions"
	"github.com/codebyyogesh/lens_view/internal/views"

	"github.com/go-chi/chi/v5"
)

func main() {
	mux := chi.NewRouter()
	// home handler, first parse and then execute
	// Must() already handles the panic during the start
	// Register your API route handlers of mux using Get, Post, Put and Delete methods
	tpl := views.Must(views.ParseFS(assets.EmbeddedFiles,
		"templates/pages/home.tmpl",
		"templates/pages/tailwind.tmpl"))

	mux.Get("/", actions.StaticHandler(tpl))

	// contact handler, first parse and then execute
	tpl = views.Must(views.ParseFS(assets.EmbeddedFiles,
		"templates/pages/contact.tmpl",
		"templates/pages/tailwind.tmpl"))
	mux.Get("/contact", actions.StaticHandler(tpl))

	// faq handler, first parse and then execute
	tpl = views.Must(views.ParseFS(assets.EmbeddedFiles,
		"templates/pages/faq.tmpl",
		"templates/pages/tailwind.tmpl"))
	mux.Get("/faq", actions.FAQ(tpl))

	// contact handler, first parse and then execute
	userSignUp := actions.Users{}
	userSignUp.New = views.Must(views.ParseFS(assets.EmbeddedFiles,
		"templates/pages/signup.tmpl",
		"templates/pages/tailwind.tmpl"))

	mux.Get("/signup", userSignUp.NewHandler)
	// POST API route handler for form /signup
	mux.Post("/signup", userSignUp.Create)

	mux.NotFound(func(w http.ResponseWriter, r *http.Request) { http.Error(w, "Page not found", http.StatusNotFound) })
	fmt.Println("Server listening on port :4444")
	http.ListenAndServe(":4444", mux)
}
