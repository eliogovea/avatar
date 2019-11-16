package api

import (
  "net/http"
//  "log"
  "io/ioutil"
  "encoding/json"
  "bytes"
)

func (h *handler) uploadHandler() http.HandlerFunc {
  type request struct {
    Image string `json:"image"`
  }
  return func(w http.ResponseWriter, r *http.Request) {
    body, err := ioutil.ReadAll(r.Body)
    if err != nil {
      http.Error(w, err.Error(), http.StatusBadRequest)
      return
    }
    var data request
    decoder := json.NewDecoder(bytes.NewBuffer(body))
    err = decoder.Decode(&data)
    if err != nil {
      http.Error(w, err.Error(), 401)
      return
    }
    username := r.Context().Value("username").(string)
    err = h.srv.Storage.SaveImage(username, data.Image)
    if err != nil {
      http.Error(w, err.Error(), http.StatusInternalServerError)
      return
    }
    w.WriteHeader(http.StatusOK)
  }
}
