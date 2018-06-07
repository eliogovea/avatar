package main

import (
	"net/http"
)

func (s *server)isLogged(r *http.Request) bool {
	cookie, err := r.Cookie(s.LoginCookieName)
	if err != nil {
		return false
	}
	return s.sessions.IsIdActive(cookie.Value)
}


func (s *server)notLoggedOnly(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if s.isLogged(r) {
			w.Header().Set("Location", s.PersonalPath)
			w.WriteHeader(http.StatusMovedPermanently)
			return
		}
		f(w, r)
	}
}

func (s *server)loggedOnly(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !s.isLogged(r) {
			w.Header().Set("Location", s.LoginPath)
			w.WriteHeader(http.StatusMovedPermanently)
			return
		}
		f(w, r)
	}
}


func (s *server)managerOnly(f http.HandlerFunc) http.HandlerFunc {
	// TODO once load facedetectio data
	return func(w http.ResponseWriter, r *http.Request) {
		f = s.loggedOnly(f)
		// TODO
		f(w, r)
	}
}
