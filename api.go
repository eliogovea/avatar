package main

import (
	"io"
	"os"
	"net/http"
	"log"
)

// /api/approved/{id}
func (s *server)getApprovedAvatar() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			// TODO
			return
		}
		log.Println(">>>> ", r.URL.RequestURI())
		username := r.URL.RequestURI()[len("/api/approved/"):]
		log.Println("get approved avatar -> username: ", username)
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
