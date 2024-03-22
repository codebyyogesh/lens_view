package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/go-chi/chi/v5"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	tplPath := filepath.Join("assets", "templates", "pages", "home.tmpl")
	tpl, err := template.ParseFiles(tplPath)
	if err != nil {
		log.Printf("template parsing error: %v", err)
		http.Error(w, "Error in template parsing", http.StatusInternalServerError)
		return
	}
	err = tpl.Execute(w, nil)
	if err != nil {
		log.Printf("template executing error: %v", err)
		http.Error(w, "Error in template executing", http.StatusInternalServerError)
	}
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Contact Page</h1><p>To get in touch, email me at <a href=\"mailto:yogidk@gmail.com\">yogidk@gmail</a>!</p>") // nolint:forbidh
}

func faqHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, `<h1>FAQ Page</h1> 
	<ul> 
		<li>
			<b>Is there a free version?</b>
			Yes! We offer a free trial for 30 days on any paid plans. 
		</li> 
		<li>
			<b>What are your support hours?</b> 
			We have support staff answering emails 24/7, though response times may be a bit slower on weekends. 
		</li> 
		<li> 
			<b>How do I contact support?</b> 
			Email us - <a href="mailto:support@lensview.com">support@lensview.com</a> 
		</li>
	</ul>`)
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
