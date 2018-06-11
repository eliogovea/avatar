package main

import (
	"net/http"
	"log"
	"html/template"
	"time"
)

func (s *server)loginHandler(tempAddr string) http.HandlerFunc {
	t := template.Must(template.ParseFiles(tempAddr))
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {	
			t.Execute(w, nil)
			return
		}

		username := r.FormValue("username")
		password := r.FormValue("password")

		err, isManager := s.AuthMan.CheckUserAndPass(username, password)

		if err == nil {
			log.Println("authentication accepted " + username)

			newId, err := s.sessions.AddSession(username, isManager) // TODO handle error

			if err != nil {
				log.Println("error creating session:", err)
			}

			cookie := &http.Cookie {
				Name: s.LoginCookieName,
				Value: newId,
				Expires: time.Now().Add(s.SessionDuration),
			}

			http.SetCookie(w, cookie)

			w.Header().Set("Location", s.PersonalPath)
			w.WriteHeader(http.StatusMovedPermanently)
			return
		}

//		log.Println("authentication rejected: ", err)

		// TODO show login again + report error
		w.Header().Set("Location", s.LoginPath)
		w.WriteHeader(http.StatusMovedPermanently)
	}
}
