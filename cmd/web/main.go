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
	tpl := views.Must(views.ParseFS(assets.EmbeddedFiles, "templates/pages/home.tmpl"))
	mux.Get("/", actions.StaticHandler(tpl))

	// contact handler, first parse and then execute
	tpl = views.Must(views.ParseFS(assets.EmbeddedFiles, "templates/pages/contact.tmpl"))
	mux.Get("/contact", actions.StaticHandler(tpl))

	// faq handler, first parse and then execute
	tpl = views.Must(views.ParseFS(assets.EmbeddedFiles, "templates/pages/faq.tmpl"))
	mux.Get("/faq", actions.StaticHandler(tpl))

	mux.NotFound(func(w http.ResponseWriter, r *http.Request) { http.Error(w, "Page not found", http.StatusNotFound) })
	fmt.Println("Server listening on port :4444")
	http.ListenAndServe(":4444", mux)
}
