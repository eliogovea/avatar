package api

import (
  "net/http"

  "github.com/eliogovea/upr-avatar/internal/service"
)

type handler struct {
  srv *service.Service
}

func New(srv *service.Service, assetsPath string) http.Handler {
  api := &handler{srv}

  router := http.NewServeMux()

  router.HandleFunc("/", api.rootHandler(assetsPath + "/static/index.html"))

  router.HandleFunc("/login", api.loginHandler())

  router.HandleFunc("/avatars/approved/", api.getApprovedHandler())
  router.HandleFunc("/avatars/pending/", api.getPendingHandler())

  router.HandleFunc("/upload", api.authorizedOnly(api.uploadHandler()))

  fs := http.StripPrefix("/" + assetsPath + "/", http.FileServer(http.Dir(assetsPath)))
  fs = api.withoutCache(fs)
  router.Handle("/" + assetsPath + "/", fs)

  return router
}
