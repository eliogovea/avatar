package main

import (
  "time"
  "log"
  "net/http"

  "github.com/eliogovea/upr-profile-pic/internal/service"
  "github.com/eliogovea/upr-profile-pic/internal/api"
)

func main() {
  service := service.New()
  handler := api.New(service, "web")
  server := http.Server {
    Addr:              ":8080",
    Handler:           handler,
    ReadHeaderTimeout: time.Second * 5,
    ReadTimeout:       time.Second * 10,
  }
  log.Println("listening on port 8080")
  err := server.ListenAndServe()
  log.Fatal(err.Error())
}
