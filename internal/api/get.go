package api


import (
  "log"
  "fmt"
  "net/http"
  "io/ioutil"
)

func (h *handler) getApprovedHandler() http.HandlerFunc {
  prefix := "/avatars/approved/"
  return func(w http.ResponseWriter, r *http.Request) {
    log.Println("get approved")

    if r.Method != http.MethodGet {
      http.Error(w, "forbidden method", http.StatusBadRequest)
      return
    }

    username := r.URL.RequestURI()[len(prefix):]

    url :=  fmt.Sprintf("http://sync.upr.edu.cu/api/user/%s", username)
    request, err := http.NewRequest("GET", url, nil)
    if err != nil {
      http.Error(w, err.Error(), http.StatusInternalServerError)
      return
    }

    request.Header.Set("accept", "application/json")
    response, err := http.DefaultClient.Do(request)

    if err != nil {
      log.Println("could not get avatar from sync ", err.Error())
      http.Redirect(w, r, "/web/static/img/default-avatar.jpeg", http.StatusSeeOther)
      return
    }

    defer response.Body.Close()

    if response.StatusCode != http.StatusOK {
      log.Println("could not get avatar from sync, status", response.StatusCode)
      http.Redirect(w, r, "/web/static/img/default-avatar.jpeg", http.StatusSeeOther)
      return
    }

    body, err := ioutil.ReadAll(response.Body)
    if err != nil {
      http.Error(w, err.Error(), http.StatusInternalServerError)
      return
    }

    w.Write(body)
  }
}

func (h *handler) getPendingHandler() http.HandlerFunc {
  prefix := "/avatars/pending/"
  return func(w http.ResponseWriter, r *http.Request) {
    log.Println("get pending")
    if r.Method != http.MethodGet {
      http.Error(w, "forbidden method", http.StatusBadRequest)
      return
    }

    username := r.URL.RequestURI()[len(prefix):]

    log.Println("username", username)

    image, err := h.srv.Storage.GetImage(username)

    log.Println("error", err)

    if err != nil {
      log.Println("could not get pending avatar ", username, err.Error())
      http.Redirect(w, r, "/web/static/img/default-avatar.jpeg", http.StatusSeeOther)
      return
    }

    w.Header().Set("Content-Type", "image/jpeg")
    w.Write(image)
  }
}
