package main

import (
	"net/http"
)

func Handler() http.HandlerFunc {
	return func(w http.ResponseWriter,r *http.Request) {
		_, _ = w.Write([]byte("hello\n"))
	}
}

func main() {
	mux := http.NewServeMux()

	mux.Handle("/health", Handler())

	server := http.Server{
		Addr: ":8089",
		Handler: mux,
	}

	if err := server.ListenAndServe(); err != nil {
		panic("failed to start server")
	}
}