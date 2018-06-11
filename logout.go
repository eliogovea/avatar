package main

import (
	"net/http"
//	"log"
	"time"
)

func (s *server)logoutHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie(s.LoginCookieName)
		if err == nil {
			id := cookie.Value
// 			log.Println("deleting session: " + id)
			err = s.sessions.DeleteSession(id)
			if err != nil {
				// TODO report error
				panic(err)
			}
			expiredCookie := &http.Cookie {
				Path: "",
				Name: s.LoginCookieName,
				MaxAge: -1,
				Expires: time.Now().Add(-100 * time.Hour),
			}
			http.SetCookie(w, expiredCookie)	
		}
		w.Header().Set("Location", s.LoginPath)
		w.WriteHeader(http.StatusFound)
	}
}
