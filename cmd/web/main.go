package main

import (
	"fmt"
	"net/http"

	"github.com/codebyyogesh/lens_view/assets"
	"github.com/codebyyogesh/lens_view/internal/actions"
	"github.com/codebyyogesh/lens_view/internal/database"
	"github.com/codebyyogesh/lens_view/internal/database/migrations"
	"github.com/codebyyogesh/lens_view/internal/views"

	"github.com/gorilla/csrf"

	"github.com/go-chi/chi/v5"
)

func main() {
	mux := chi.NewRouter()
	// home handler, first parse and then execute
	// Must() already handles the panic during the start
	// Register your API route handlers of mux using Get, Post, Put and Delete methods
	tpl := views.Must(views.ParseFS(assets.EmbeddedFiles,
		"templates/pages/home.tmpl",
		"templates/pages/tailwind.tmpl"))

	mux.Get("/", actions.StaticHandler(tpl))

	// contact handler, first parse and then execute
	tpl = views.Must(views.ParseFS(assets.EmbeddedFiles,
		"templates/pages/contact.tmpl",
		"templates/pages/tailwind.tmpl"))
	mux.Get("/contact", actions.StaticHandler(tpl))

	// faq handler, first parse and then execute
	tpl = views.Must(views.ParseFS(assets.EmbeddedFiles,
		"templates/pages/faq.tmpl",
		"templates/pages/tailwind.tmpl"))
	mux.Get("/faq", actions.FAQ(tpl))

	// Setup the DB connection
	cfg := database.DefaultPostgresConfig()
	db, err := database.Open(cfg)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	// db migrations with Goose
	err = database.MigrateFS(db, migrations.FS, ".")
	if err != nil {
		panic(err)
	}
	// setup the db(db acts here as a model of mvc pattern) store
	userStore := database.UserStore{
		DB: db,
	}
	sessionStore := database.SessionStore{
		DB: db,
	}
	// contact handler, first parse and then execute
	user := actions.Users{
		UserStore:    &userStore,
		SessionStore: &sessionStore,
	}

	// SignUp Page creation
	user.Templates.New = views.Must(views.ParseFS(assets.EmbeddedFiles,
		"templates/pages/signup.tmpl",
		"templates/pages/tailwind.tmpl"))

	mux.Get("/signup", user.NewHandler)
	// POST API route handler for form /signup
	mux.Post("/signup", user.Create)

	// SignIn Page Creation
	user.Templates.SignIn = views.Must(views.ParseFS(assets.EmbeddedFiles,
		"templates/pages/signin.tmpl",
		"templates/pages/tailwind.tmpl"))

	mux.Get("/signin", user.SignInHandler)
	mux.Post("/signin", user.ProcessSignInHandler)
	mux.Post("/signout", user.ProcessSignOutHandler)

	mux.Get("/users/me", user.CurrentUser)
	mux.NotFound(func(w http.ResponseWriter, r *http.Request) { http.Error(w, "Page not found", http.StatusNotFound) })
	fmt.Println("Server listening on port :4444")
	csrfKey := "8jRMm5SZweHXO0ngTy80E7uG7t0hO6LX"
	csrfMiddleware := csrf.Protect(
		[]byte(csrfKey),
		csrf.Secure(false), //ToDO: Fix this before deploying
	)

	http.ListenAndServe(":4444", csrfMiddleware(mux))
}
