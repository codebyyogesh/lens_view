package main

import (
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/codebyyogesh/lens_view/internal/actions"
	"github.com/codebyyogesh/lens_view/internal/views"
	"github.com/go-chi/chi/v5"
)

func main() {
	mux := chi.NewRouter()
	// home handler, first parse and then execute
	tpl, err := views.Parse(filepath.Join("assets", "templates", "pages", "home.tmpl"))
	if err != nil {
		panic(err)
	}
	mux.Get("/", actions.StaticHandler(tpl))

	// contact handler, first parse and then execute
	tpl, err = views.Parse(filepath.Join("assets", "templates", "pages", "contact.tmpl"))
	if err != nil {
		panic(err)
	}
	mux.Get("/contact", actions.StaticHandler(tpl))

	// faq handler, first parse and then execute
	tpl, err = views.Parse(filepath.Join("assets", "templates", "pages", "faq.tmpl"))
	if err != nil {
		panic(err)
	}
	mux.Get("/faq", actions.StaticHandler(tpl))

	mux.NotFound(func(w http.ResponseWriter, r *http.Request) { http.Error(w, "Page not found", http.StatusNotFound) })
	fmt.Println("Server listening on port :4444")
	http.ListenAndServe(":4444", mux)
}
