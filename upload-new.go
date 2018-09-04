package main

import (
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

func (s *server) uploadNewHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if r.Method != http.MethodPost {
			http.Error(w, "only post allowed", http.StatusMethodNotAllowed)
			return
		}

		if r.ContentLength > int64(s.MaxUploadSize) {
			log.Println(r.ContentLength, s.MaxUploadSize, s.Fs.MaxUploadSize)
			http.Error(w, "request too large", http.StatusExpectationFailed)
			return
		}

		r.Body = http.MaxBytesReader(w, r.Body, int64(s.MaxUploadSize))

		username := r.Context().Value("username").(string)
		// log.Println("context variable username: ", username)

		r.ParseForm()
		// log.Println(r.FormValue("imagebase64"))

		image := r.FormValue("imagebase64")

		comma := strings.Index(string(image), ",")

		info := string(image)[:comma]
		rawImage := string(image)[comma+1:]

		log.Println(info)
		log.Println(rawImage)

		imageOK, err := base64.StdEncoding.DecodeString(rawImage)

		if err != nil {
			log.Println(err.Error())
		}

		var file *os.File

		if info == fmt.Sprintf("data:image/%s;base64", "jpeg") {
			file, err = os.Create("avatars/tmp/" + username + ".jpeg")

			if err != nil {
				log.Println("error !!!!!!", err.Error())
				// TODO report error
				panic(err)
			}

			file.Write(imageOK)

		}

		if info == fmt.Sprintf("data:image/%s;base64", "png") {
			file, err := os.Create("avatars/tmp/" + username + ".jpeg")

			if err != nil {
				log.Println("error !!!!!!", err.Error())
				// TODO report error
				panic(err)
			}
			file.Write(imageOK)
		}

		file.Close()

		if err != nil {
			log.Println("!!!error creating pending: ", err)
		}

		file, err = os.Open("avatars/tmp/" + username + ".jpeg")
		defer file.Close()

		err = s.Fs.CreatePending(file, username)

		if err != nil {
			log.Println("error creating pending: ", err)
		}

		w.Header().Set("Location", s.PersonalPath)
		w.WriteHeader(http.StatusSeeOther)

	}
}
