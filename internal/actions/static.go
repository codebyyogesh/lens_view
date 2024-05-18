package actions

// This package is nothing but the controller package of MVC.
//  It acts as intermediaries between the user's requests (handled through HTTP handlers) and the application's models and views.
// Can handle Receiving/Processing Requests, Interacting with Models/Views, Rendering Views etc
import (
	"html/template"
	"net/http"
)

func StaticHandler(t Template) http.HandlerFunc {
	// we return a closure
	return func(w http.ResponseWriter, r *http.Request) {
		t.Execute(w, nil)
	}
}

func FAQ(t Template) http.HandlerFunc {
	questions := []struct {
		Question string
		Answer   template.HTML
	}{
		{
			Question: "Is there a free version?",
			Answer:   "Yes! We offer a free trial for 30 days on any paid plans.",
		},

		{
			Question: "What are your support hours?",
			Answer:   "We have support staff answering emails 24/7, though response times may be a bit slower on weekends.",
		},

		{
			Question: "How do I contact support?",
			Answer:   `Email us - <a href="mailto:support@lensview.in">support@lensview.in</a>`,
		},

		{
			Question: "Can I cancel my subscription?",
			Answer:   "Yes, you can. Simply send an email to <a href='mailto:support@lensview.in'>support@lensview.in</a> and we'll process your cancellation request.",
		},
	}

	return func(w http.ResponseWriter, r *http.Request) {
		t.Execute(w, questions)
	}
}
