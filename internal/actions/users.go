package actions

import (
	"fmt"
	"net/http"
)

type Users struct {
	New Template
}

func (u Users) NewHandler(w http.ResponseWriter, r *http.Request) {
	u.New.Execute(w, nil)
}

func (u Users) Create(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Email:  %v\n", r.FormValue("email"))
	fmt.Fprintf(w, "Password:  %v\n", r.FormValue("password"))
}
