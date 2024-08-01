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
	UserStore    *database.UserStore
	SessionStore *database.SessionStore
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
	u.Templates.New.Execute(w, r, data)
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

	// Create a new session for the user
	session, err := u.SessionStore.Create(user.ID)
	if err != nil {
		fmt.Println(err)
		// TODO: Warn the user about not being able to sign the user in , but user signup was successful
		http.Redirect(w, r, "/signin", http.StatusFound) // redirect to /signin
		return
	}
	// Set the session cookie
	setCookie(w, CookieSession, session.Token)
	http.Redirect(w, r, "/users/me", http.StatusFound)
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
	u.Templates.SignIn.Execute(w, r, data)
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
	// create a session
	session, err := u.SessionStore.Create(user.ID)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Something went wrong:", http.StatusInternalServerError)
		return
	}
	// set a session cookie
	setCookie(w, CookieSession, session.Token)
	http.Redirect(w, r, "/users/me", http.StatusFound)
}

// Called when the user clicks signs out (ie. POST /signout)
func (u Users) ProcessSignOutHandler(w http.ResponseWriter, r *http.Request) {
	token, err := readCookie(r, CookieSession)
	if err != nil {
		fmt.Println(err)
		http.Redirect(w, r, "/signin", http.StatusFound)
		return
	}
	// delete the session for the user
	err = u.SessionStore.Delete(token)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Something went wrong:", http.StatusInternalServerError)
		return
	}
	// delete the session cookie as well
	deleteCookie(w, CookieSession)
	http.Redirect(w, r, "/signin", http.StatusFound)
}

// This gets called for GET /users/me ie know your current user
func (u Users) CurrentUser(w http.ResponseWriter, r *http.Request) {
	token, err := readCookie(r, CookieSession)
	if err != nil {
		fmt.Println(err)
		http.Redirect(w, r, "/signin", http.StatusFound)
		return
	}
	// Get the user from the session using tokenCookie
	user, err := u.SessionStore.UserLookup(token)
	if err != nil {
		fmt.Println(err)
		http.Redirect(w, r, "/signin", http.StatusFound)
		return
	}

	fmt.Fprintf(w, "Current User: %s\n ", user.Email)
}
