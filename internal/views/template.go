/*
Views generates a user interface for the user. Views are created by the data which  is collected by the model component but these data aren’t taken directly but through the controller (action). It only interacts with the controller(action).
*/
package views

import (
	"fmt"
	"html/template"
	"io/fs"
	"log"
	"net/http"
	"path/filepath"
)

func Must(tpl Template, err error) Template {
	if err != nil {
		panic(err)
	}
	return tpl
}

func csrfField() template.HTML {
	return template.HTML(`<input type="hidden" />`)
}

// Use ParseFS instead of Parse() so that it embeds the templates(tmpl files) into
// the final binary. Parse() cannot handle the case of running the app from a different
// directory due to relative paths of the template files
func ParseFS(fs fs.FS, patterns ...string) (Template, error) {
	// because our tmpl files are in the templates/pages folder, we need to
	// use filepath.Base(patterns[0]) to get the name of the tmpl file, else it gives
	// an error
	filename := filepath.Base(patterns[0])

	tpl := template.New(filename)
	tpl = tpl.Funcs(
		template.FuncMap{
			"csrfField": csrfField,
		},
	)

	tpl, err := tpl.ParseFS(fs, patterns...)
	if err != nil {
		return Template{}, fmt.Errorf("parseFS template: %w", err)
	}
	return Template{htmlTpl: tpl}, nil
}

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

// Execute implements actions.Template.
// receivers
func (t Template) Execute(w http.ResponseWriter, data any) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	err := t.htmlTpl.Execute(w, data)
	if err != nil {
		log.Printf("template executing error: %v", err)
		http.Error(w, "Error in template executing", http.StatusInternalServerError)
	}
}
