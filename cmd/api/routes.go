package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func (app *Application) routes() http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Logger)
	mux.Use(middleware.Recoverer)
	mux.Use(app.enableCORS)

	mux.Post("/authenticate", app.Authenticate)
	mux.Get("/refresh", app.RefreshToken)
	mux.Get("/logout", app.Logout)

	mux.Get("/questions", app.AllQuestions)
	mux.Post("/questions", app.InsertQuestion)

	mux.Route("/admin", func(mux chi.Router) {
		mux.Use(app.authRequired)

		mux.Patch("/questions/{id}", app.UpdatedQuestion)
		mux.Delete("/questions/{id}", app.DeleteQuestion)
	})

	return mux
}
