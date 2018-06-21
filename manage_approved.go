package main

import (
	"html/template"
	"net/http"
)

// /admin/pending
func (s *server) manageApproved() http.HandlerFunc {
	type file struct {
		Name     string
		Username string
	}
	type data struct {
		NoPending bool
		Files     []file
	}
	t := template.Must(template.ParseFiles("./templates/manage_approved.html"))
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "StatusMethodNotAllowed", http.StatusMethodNotAllowed)
			return
		}

		Data := new(data)
		Data.Files = make([]file, 0)
		for key, _ := range s.Fs.Approved {
			Data.Files = append(Data.Files, file{
				Name:     "/api/approved/" + key,
				Username: key,
			})
		}

		t.Execute(w, Data)

	}
}
