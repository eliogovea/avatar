package main

import (
	"context"
	"log"
	"net/http"
)

func (s *server) isLogged(r *http.Request) (bool, string, bool) {
	cookie, err := r.Cookie(s.LoginCookieName)
	if err != nil {
		return false, "", false
	}
	if !s.sessions.IsIdActive(cookie.Value) {
		return false, "", false
	}
	username, isManager, err := s.sessions.GetInfo(cookie.Value)

	if err != nil {
		// unexpected error
		log.Println("error (get session info): ", cookie.Value, err)
		return false, "", false
	}

	return true, username, isManager
}

func (s *server) notLoggedOnly(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ok, _, _ := s.isLogged(r)
		if ok {
			w.Header().Set("Location", s.PersonalPath)
			w.WriteHeader(http.StatusMovedPermanently)
			return
		}
		f(w, r)
	}
}

func (s *server) loggedOnly(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ok, username, isManager := s.isLogged(r)
		if !ok {
			w.Header().Set("Location", s.LoginPath)
			w.WriteHeader(http.StatusMovedPermanently)
			return
		}
		ctx := context.WithValue(r.Context(), "username", username)
		ctx = context.WithValue(ctx, "isManager", isManager)
		f(w, r.WithContext(ctx))
	}
}

func (s *server) managerOnly(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ok, username, isManager := s.isLogged(r)
		if !ok || !isManager {
			w.Header().Set("Location", s.PersonalPath)
			w.WriteHeader(http.StatusMovedPermanently)
			return
		}
		ctx := context.WithValue(r.Context(), "username", username)
		ctx = context.WithValue(ctx, "isManager", isManager)
		f(w, r.WithContext(ctx))
	}
}
