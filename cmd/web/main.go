package main

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"

	"github.com/codebyyogesh/lens_view/internal/views"
	"github.com/go-chi/chi/v5"
)

func parseAndExecuteTemplate(w http.ResponseWriter, tplPath string, data any) {
	tpl, err := views.Parse(tplPath)
	if err != nil {
		log.Printf("template parsing error: %v", err)
		http.Error(w, "Error in template parsing", http.StatusInternalServerError)
		return
	}
	tpl.Execute(w, data)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	tplPath := filepath.Join("assets", "templates", "pages", "home.tmpl")
	parseAndExecuteTemplate(w, tplPath, nil)
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	tplPath := filepath.Join("assets", "templates", "pages", "contact.tmpl")
	parseAndExecuteTemplate(w, tplPath, nil)
}

func faqHandler(w http.ResponseWriter, r *http.Request) {
	tplPath := filepath.Join("assets", "templates", "pages", "faq.tmpl")
	parseAndExecuteTemplate(w, tplPath, nil)
}

func main() {
	mux := chi.NewRouter()
	mux.Get("/", homeHandler)
	mux.Get("/contact", contactHandler)
	mux.Get("/faq", faqHandler)
	mux.NotFound(func(w http.ResponseWriter, r *http.Request) { http.Error(w, "Page not found", http.StatusNotFound) })
	fmt.Println("Server listening on port :4444")
	http.ListenAndServe(":4444", mux)
}
