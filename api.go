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

func hasPrefix(p, s string) bool {
	if len(s) < len(p) {
		return false
	}
	for i := 0; i < len(p); i++ {
		if s[i] != p[i] {
			return false
		}
	}
	return true
}

func (s *server) APIApproved(method string) http.HandlerFunc {
	switch method {
	case http.MethodGet:
		return s.getApproved()
	case http.MethodPost:
		// TODO
	case http.MethodDelete:
		return s.denyApproved()
	default:
		// TODO report error
	}
	return nil
}

func (s *server) APIPending(method string) http.HandlerFunc {
	switch method {
	case http.MethodGet:
		return s.getPending()
	case http.MethodPost:
		// TODO
	case http.MethodDelete:
		return s.denyPending()
	default:
		// TODO
	}
	return nil
}

// /api/approved/{id}
func (s *server) getApproved() http.HandlerFunc {
	prefix := "/api/approved/"
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}
		if !hasPrefix(prefix, r.URL.RequestURI()) {
			http.Error(w, "bad request", http.StatusBadRequest)
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
func (s *server) getPending() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)

		}
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
			http.Error(w, err.Error(), 500)
			return
		}
		defer file.Close()
		w.Header().Set("Content-Type", "image/jpeg")
		io.Copy(w, file)
	}
}

// /api/pending/{id}
// method ???
func (s *server) approve() http.HandlerFunc {
	root := "/api/approve/"
	return func(w http.ResponseWriter, r *http.Request) {
		username := r.URL.RequestURI()[len(root):]
		log.Println("approve pending", username)
		err := s.Fs.approvePending(username)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}

// /api/pending/{id}
// method delete
func (s *server) denyPending() http.HandlerFunc {
	root := "/api/deny/pending/"
	return func(w http.ResponseWriter, r *http.Request) {
		username := r.URL.RequestURI()[len(root):]
		err := s.Fs.denyPending(username)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}

// /api/approved/{id}
// method delete
func (s *server) denyApproved() http.HandlerFunc {
	root := "/api/deny/approved/"
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			http.Error(w, "bad request", http.StatusMethodNotAllowed)
			return
		}
		username := r.URL.RequestURI()[len(root):]
		err := s.Fs.denyApproved(username)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}
