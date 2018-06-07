package main

import (
	"net/http"
)

func (s *server)uploadHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "only post allowed", http.StatusMethodNotAllowed)
			return
		}
		if r.ContentLength > int64(s.MaxUploadSize) {
			http.Error(w, "request too large", http.StatusExpectationFailed)
			return
		}
		r.Body = http.MaxBytesReader(w, r.Body, int64(s.MaxUploadSize))


	}
}
