package main

import (
	"net/http"
)

func (s *server) rootHandler() http.HandlerFunc {
	// TODO something to load ???
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Location", "/personal")
		w.WriteHeader(http.StatusMovedPermanently)
	}
}
