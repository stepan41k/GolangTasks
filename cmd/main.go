package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

func GetHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var req string

		err := render.DecodeJSON(r.Body, &req)
		if err != nil {
			panic(err)
		}
		
		render.JSON(w, r, req)
	}
}

func main() {
	router := chi.NewRouter()

	router.Route("/service1", func(r chi.Router) {
		r.Get("/get", GetHandler())
	})

	newServer := http.Server{
		Addr: "localhost:8041",
		Handler: router,
	}

	err := newServer.ListenAndServe()
	panic(err)
}
