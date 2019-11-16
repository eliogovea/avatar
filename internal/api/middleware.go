package api

import (
  "context"
  "net/http"
  "strings"
  "time"
)

func (h *handler) authorizedOnly(f http.HandlerFunc) http.HandlerFunc {
  return func(w http.ResponseWriter, r *http.Request) {
    authorizationHeader := r.Header.Get("authorization")
    token := strings.Split(authorizationHeader, " ")
    if len(token) != 2 || h.srv.Authorizer.CheckToken(token[1]) != nil {
      http.Error(w, "unauthorized", http.StatusUnauthorized)
      return
    }
    username, err := h.srv.Authorizer.GetUsername(token[1])
    if err != nil {
      http.Error(w, "unauthorized", http.StatusUnauthorized)
      return
    }
    ctx := context.WithValue(r.Context(), "username", username)
		f(w, r.WithContext(ctx))
  }
}

func (h *handler) adminsOnly(f http.HandlerFunc) http.HandlerFunc {
  return func(w http.ResponseWriter, r *http.Request) {
    // TODO
    f(w, r)
  }
}

func (h *handler) withoutCache(next http.Handler) http.Handler {
  var epoch = time.Unix(0, 0).Format(time.RFC1123)
  var noCacheHeaders = map[string]string{
	  "Expires":         epoch,
	  "Cache-Control":   "no-cache, no-store, no-transform, must-revalidate, private, max-age=0",
	  "Pragma":          "no-cache",
	  "X-Accel-Expires": "0",
  }
  var etagHeaders = []string{
	  "ETag",
	  "If-Modified-Since",
	  "If-Match",
	  "If-None-Match",
	  "If-Range",
	  "If-Unmodified-Since",
  }
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		for _, v := range etagHeaders {
			r.Header.Del(v)
		}

		wh := w.Header()
		for k, v := range noCacheHeaders {
			wh.Set(k, v)
		}
		next.ServeHTTP(w, r)
	})
}
