package actions

import (
	"net/http"

	"github.com/codebyyogesh/lens_view/internal/views"
)

type Users struct {
	New views.Template
}

func (u Users) NewHandler(w http.ResponseWriter, r *http.Request) {
	u.New.Execute(w, nil)
}
