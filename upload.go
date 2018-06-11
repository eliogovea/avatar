package main

import (
	"net/http"
	"log"
)


func (s *server)uploadHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		log.Println("upload file")

		if r.Method != http.MethodPost {
			http.Error(w, "only post allowed", http.StatusMethodNotAllowed)
			return
		}

		if r.ContentLength > int64(s.MaxUploadSize) {
			http.Error(w, "request too large", http.StatusExpectationFailed)
			return
		}


		r.Body = http.MaxBytesReader(w, r.Body, int64(s.MaxUploadSize))

		username := r.Context().Value("username").(string)
		// log.Println("context variable username: ", username)

		file, _, err := r.FormFile("file")

		if err != nil {
			log.Println("error reading uploaded file: ", err)
			// http.Error(w, "error reading file", /* TODO */)
		}
		
		err = s.Fs.CreatePending(file, username)

		if err != nil {
			log.Println("error creating pending: ", err)
		}

		w.Header().Set("Location", s.PersonalPath)
		w.WriteHeader(http.StatusSeeOther)

	}
}
