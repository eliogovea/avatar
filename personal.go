package main

import (
	"log"
	"net/http"
	"html/template"
)

type PersonalData struct {
	HasApprovedAvatar bool
	ApprovedAvatar 	  string

	HasPendingAvatar	bool
	PendingAvatar		string
}

func (s *server)getSessionId(r *http.Request) string {
	cookie, err := r.Cookie(s.LoginCookieName)
	if err != nil {
		return ""
	}
	return cookie.Value
}

func (s *server)personalHandler(tempAddr string) http.HandlerFunc {
	t := template.Must(template.ParseFiles(tempAddr))
	return func(w http.ResponseWriter, r *http.Request) {
		username, _ /*isManager*/, err := s.sessions.GetInfo(s.getSessionId(r))
		if err != nil {
			log.Println("!!! error personalHandler !!!")
			// TODO internal error
		}
		data := &PersonalData{
			ApprovedAvatar: "/api/approved/" + username,
		}
		t.Execute(w, data)
	}
}
