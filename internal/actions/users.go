package actions

import (
	"fmt"
	"net/http"

	"github.com/codebyyogesh/lens_view/internal/database"
)

type Users struct {
	// club all templates in an anonymous struct
	Templates struct {
		New    Template
		SignIn Template
	}
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
	u.Templates.New.Execute(w, data)
}

// This gets called when the signup form is filled and submitted (ie.e POST /signup)
func (u Users) Create(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	password := r.FormValue("password")

	user, err := u.UserStore.Create(email, password)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Something went wrong:", http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "User created: %+v\n", user)
}

// user sign in handler (ie. GET /signin)
func (u Users) SignInHandler(w http.ResponseWriter, r *http.Request) {
	// anonymous struct
	// This is used to pre fill data in the sign in page
	// ie when you type in the browser http://localhost:4444/signin?email=bi@bi.io
	// email was already pre filled and this data is passed to the signin template
	//page before rendering it
	data := struct {
		Email string
	}{
		Email: r.FormValue("email"),
	}
	u.Templates.SignIn.Execute(w, data)
}

// This gets called when the signin form is filled and submitted (ie. POST /signin)
func (u Users) ProcessSignInHandler(w http.ResponseWriter, r *http.Request) {
	// anonymous struct for convenience
	data := struct {
		Email    string
		Password string
	}{
		Email:    r.FormValue("email"),
		Password: r.FormValue("password"),
	}
	user, err := u.UserStore.Authenticate(data.Email, data.Password)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Something went wrong:", http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "User authenticated: %+v\n", user)
}
