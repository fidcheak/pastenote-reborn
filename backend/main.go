package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func main() {
	DB, _ = InitDB()
	// if err := DeleteTables(DB); err != nil {
	// 	log.Fatal(err)
	// }
	if err := InitTables(DB); err != nil {
		log.Fatal(err)
	}
	r := chi.NewRouter()
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"*"},
		// AllowedOrigins: []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		// ExposedHeaders:   []string{"Link"},
		// AllowCredentials: false,
		// MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Post("/sign-up", SignUp)
	r.Post("/sign-in", SignIn)
	r.Delete("/user-delete", UserDelete)

	r.Post("/note-create", NoteCreate)
	r.Get("/notes-get", NotesGet)
	r.Put("/note-update", NoteUpdate)
	r.Delete("/note-delete", NoteDelete)

	r.Put("/note-pin", NotePin)
	r.Put("/note-unpin", NoteUnpin)

	r.Put("/note-feature", NoteFeature)
	r.Put("/note-unfeature", NoteUnfeature)

	http.ListenAndServe(":8080", r)
}
