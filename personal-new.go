package main

import (
	"html/template"
	"log"
	"net/http"
)

func (s *server) personalNewHandler() http.HandlerFunc {
	t := template.Must(template.ParseFiles("templates/personal_new.html"))
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO get username, and isManager from context variables
		username, isManager, err := s.sessions.GetInfo(s.getSessionId(r))
		if err != nil {
			log.Println("!!! error personalHandler !!!")
		}
		data := &PersonalData{
			Username:          username,
			IsManager:         isManager,
			HasApprovedAvatar: s.Fs.HasApproved(username),
			ApprovedAvatar:    "/api/approved/" + username,

			HasPendingAvatar: s.Fs.HasPending(username),
			PendingAvatar:    "/api/pending/" + username,
		}
		t.Execute(w, data)
	}
}
