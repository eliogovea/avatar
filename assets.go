package main

import (
	"io"
	"os"
	"log"
	"net/http"
)

func hasPrefix(s, p string) bool {
	if len(s) < len(p) {
		return false
	}
	for i := 0; i < len(p); i++ {
		if s[i] != p[i] {
			return false
		}
	}
	return true
}

func (s *server)assetsHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!")
		path := r.URL.RequestURI()
		if !hasPrefix(path, "/assets/") {
			log.Println("error!!!!!!!!!!")
			return
		}
		file, err := os.Open(path)
		if err != nil {
			log.Println(" file error")
		}
		defer file.Close()
		io.Copy(w, file)
	}
}
