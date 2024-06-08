package actions

import (
	"fmt"
	"net/http"

	"github.com/codebyyogesh/lens_view/internal/database"
)

type Users struct {
	New       Template
	UserStore *database.UserStore
}

// Renders the signup Page (ie. GET /signup)
func (u Users) NewHandler(w http.ResponseWriter, r *http.Request) {
	// anonymous struct
	// This is used to pre fill data in the sign up page
	// ie when you type in the browser http://localhost:4444/signup?email=bi@bi.io
	// email was already pre filled and this data is passed to the signup template
	//page before rendering it
	data := struct {
		Email string
	}{
		Email: r.FormValue("email"),
	}
	u.New.Execute(w, data)
}

// This gets called when the signup form is filled and submitted (ie.e POST /signup)
func (u Users) Create(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Email:  %v\n", r.FormValue("email"))
	fmt.Fprintf(w, "Password:  %v\n", r.FormValue("password"))
}
