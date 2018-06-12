package main

import (
	"html/template"
	"log"
	"net/http"
)

type PersonalData struct {
	HasApprovedAvatar bool
	ApprovedAvatar    string
	HasPendingAvatar  bool
	PendingAvatar     string
}

func (s *server) getSessionId(r *http.Request) string {
	cookie, err := r.Cookie(s.LoginCookieName)
	if err != nil {
		return ""
	}
	return cookie.Value
}

func (s *server) personalHandler(tempAddr string) http.HandlerFunc {
	t := template.Must(template.ParseFiles(tempAddr))
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO get username, and isManager from context variables
		username, _ /*isManager*/, err := s.sessions.GetInfo(s.getSessionId(r))
		if err != nil {
			log.Println("!!! error personalHandler !!!")
		}
		data := &PersonalData{
			HasApprovedAvatar: s.Fs.HasApproved(username),
			ApprovedAvatar:    "/api/approved/" + username,

			HasPendingAvatar: s.Fs.HasPending(username),
			PendingAvatar:    "/api/pending/" + username,
		}
		t.Execute(w, data)
	}
}
