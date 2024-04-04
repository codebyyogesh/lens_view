package views

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

// returns the Template struct
func Parse(tplPath string) (Template, error) {
	tpl, err := template.ParseFiles(tplPath)
	if err != nil {
		return Template{}, fmt.Errorf("parsing template: %w", err)
	}
	return Template{htmlTpl: tpl}, nil
}

type Template struct {
	htmlTpl *template.Template
}

// receivers
func (t *Template) Execute(w http.ResponseWriter, data any) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	err := t.htmlTpl.Execute(w, data)
	if err != nil {
		log.Printf("template executing error: %v", err)
		http.Error(w, "Error in template executing", http.StatusInternalServerError)
	}
}
