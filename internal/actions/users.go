package actions

import (
	"net/http"
)

type Users struct {
	New Template
}

func (u Users) NewHandler(w http.ResponseWriter, r *http.Request) {
	u.New.Execute(w, nil)
}
