package api

import (
  "log"
  "bytes"
  "testing"
  "net/http"
  "net/http/httptest"

  "github.com/eliogovea/upr-profile-pic/internal/service"
)

func TestLoginRoute(t *testing.T) {
  body := []byte(`{"username":"123","password":"234"}`)
  request, err := http.NewRequest("POST", "/login", bytes.NewBuffer(body))
  if err != nil {
    t.Error(err.Error())
  }
  recorder := httptest.NewRecorder()
  service := service.New()
  handler := New(service, "../../web/")
  handler.ServeHTTP(recorder, request)
  if status := recorder.Code; status != http.StatusOK {
    t.Error("expected 200, found ", status)
  }
  log.Println("response: ", recorder.Body.String())
}
