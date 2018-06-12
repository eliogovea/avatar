package main

import (
	"html/template"
	"net/http"
)

// /admin/pending
func (s *server) managePending() http.HandlerFunc {
	type data struct {
		files []string
	}
	t := template.Must(template.ParseFiles("./templates/manage_pending.html"))
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "StatusMethodNotAllowed", http.StatusMethodNotAllowed)
			return
		}

		t.Execute(w, nil)

	}
}
