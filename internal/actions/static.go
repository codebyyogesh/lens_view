package actions

// This package is nothing but the controller package of MVC.
//  It acts as intermediaries between the user's requests (handled through HTTP handlers) and the application's models and views.
// Can handle Receiving/Processing Requests, Interacting with Models/Views, Rendering Views etc
import (
	"net/http"

	"github.com/codebyyogesh/lens_view/internal/views"
)

func StaticHandler(t views.Template) http.HandlerFunc {
	// we return a closure
	return func(w http.ResponseWriter, r *http.Request) {
		t.Execute(w, nil)
	}
}
