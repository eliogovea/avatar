package api

import (
  "net/http"
  "encoding/json"
  "io/ioutil"
  "log"
  "bytes"
)

func (h *handler) loginHandler() http.HandlerFunc {
  type request struct {
    Username string `json:"username"`
    Password string `json:"password"`
  }
  type response struct {
    Token string `json:"token"`
    Role  string `json:"role"`
  }
  return func(w http.ResponseWriter, r *http.Request) {
//    log.Println("request /login")
    var err error
    cred := request{}

    body, err := ioutil.ReadAll(r.Body)
    if err != nil {
//      log.Println("error reading body", err)
      http.Error(w, err.Error(), http.StatusBadRequest)
    }
//    log.Println("body: ", string(body))

    decoder := json.NewDecoder(bytes.NewBuffer(body))
    err = decoder.Decode(&cred)
    if err != nil {
//      log.Println("error decoding: ", err.Error())
      http.Error(w, err.Error(), 401)
      return
    }

//    log.Println("username: ", cred.Username)
//    log.Println("password: ", cred.Password)

    err = h.srv.Authorizer.CheckCredentials(cred.Username, cred.Password)

    if err != nil {
      log.Println("!!! login error: ", err.Error())
      http.Error(w, err.Error(), http.StatusUnauthorized)
      return
    }

    log.Println("authentication OK")

    w.Header().Set("Content-Type", "application/json; charset=utf-8")
    w.WriteHeader(http.StatusOK)
    encoder := json.NewEncoder(w);
    encoder.Encode(response{
      Token: h.srv.Authorizer.GenerateToken(cred.Username, ""),
      Role: "admin",
    });
  }
}


