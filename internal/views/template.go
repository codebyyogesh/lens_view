/*
Views generates a user interface for the user. Views are created by the data which  is collected by the model component but these data aren’t taken directly but through the controller (action). It only interacts with the controller(action).
*/
package views

import (
	"bytes"
	"fmt"
	"html/template"
	"io"
	"io/fs"
	"log"
	"net/http"
	"path/filepath"

	"github.com/gorilla/csrf"
)

func Must(tpl Template, err error) Template {
	if err != nil {
		panic(err)
	}
	return tpl
}

func csrfField() (template.HTML, error) {
	return "", fmt.Errorf("csrfField not implemented")
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
			"csrfField": csrfField, // just a placeholder function for now for parsing. It will be overridden later in Execute
		},
	)

	tpl, err := tpl.ParseFS(fs, patterns...)
	if err != nil {
		return Template{}, fmt.Errorf("parseFS template: %w", err)
	}
	return Template{htmlTpl: tpl}, nil
}

type Template struct {
	htmlTpl *template.Template
}

// Execute implements actions.Template.
// receivers
func (t Template) Execute(w http.ResponseWriter, r *http.Request, data any) {

	//Concurrent web requests in separate goroutines(created by net/http for every web request a goroutine is created) can cause incorrect CSRF tokens due to shared templates. Cloning the template before adding request-specific functions prevents this issue.
	tpl, err := t.htmlTpl.Clone()
	if err != nil {
		log.Printf("template cloning error: %v", err)
		http.Error(w, "There was an error rendering page.", http.StatusInternalServerError)
		return
	}

	// Add custom functions to tpl, creating a new template instance
	tpl = tpl.Funcs(template.FuncMap{
		"csrfField": func() template.HTML {
			return csrf.TemplateField(r) // We use the anonymous function to pass in the template's functions so that we do not have to pass http.Request as a param in every function such as csrfField()
		},
	})
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	// Execute the modified template (tpl), not the original one (t.htmlTpl)
	// To get a proper http.Error response if there is an error we use an in memory buffer, otherwise we get the error as - http: superfluous response.WriteHeader call from views.Template.Execute which may not reflect the correct error condition
	// PS: If the html page rendered is big, then do not use buffer, rather directly use w instead of buf, ie tpl.Execute(w, data), in such a case the http.Error will not be properly handled, but the performance will be much better
	var buf bytes.Buffer

	err = tpl.Execute(&buf, data)
	if err != nil {
		log.Printf("template executing error: %v", err)
		http.Error(w, "Error in template executing", http.StatusInternalServerError)
		return
	}
	io.Copy(w, &buf)
}
