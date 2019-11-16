package api

import (
  "net/http"
  "html/template"
)

func (h *handler) rootHandler(templatePath string) http.HandlerFunc {
  t := template.Must(template.ParseFiles(templatePath))
  return func(w http.ResponseWriter, r *http.Request) {
    t.Execute(w, nil)
  }
}
