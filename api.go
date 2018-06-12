package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func (s *server) getDefaultAvatar() http.HandlerFunc {
	file, err := os.Open(s.Fs.DefaultAvatar)
	if err != nil {
		log.Panic(err)
	}
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "image/jpeg")
		io.Copy(w, file)
	}
}

// /api/approved/{id}
func (s *server) getApprovedAvatar() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "only get", http.StatusMethodNotAllowed)
			// TODO
			return
		}
		username := r.URL.RequestURI()[len("/api/approved/"):]
		path := s.Fs.getApproved(username)
		file, err := os.Open(path)
		if err != nil {
			log.Println("error get approved: ", err)
			// error
			return
		}
		defer file.Close()
		w.Header().Set("Content-Type", "image/jpeg")
		io.Copy(w, file)
	}
}

// /api/pending/{id}
func (s *server) getPendingAvatar() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		username := r.Context().Value("username").(string)
		isManager := r.Context().Value("isManager").(bool)

		if !isManager {
			if username != r.URL.RequestURI()[len("/api/pending/"):] {
				http.Error(w, "not allowed", http.StatusForbidden)
				return
			}
		}

		if r.Method != http.MethodGet {
			http.Error(w, "only get", http.StatusMethodNotAllowed)
			return
		}

		path := s.Fs.getPending(username)
		file, err := os.Open(path)
		if err != nil {
			log.Println("error get pending: ", err)
			// TODO handle error
			return
		}
		defer file.Close()
		w.Header().Set("Content-Type", "image/jpeg")
		io.Copy(w, file)
	}
}
