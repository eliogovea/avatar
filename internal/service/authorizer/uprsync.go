package authorizer

import (
  "fmt"
  "errors"
  "net/http"
)

type uprsync struct {
}

func newUprsync() *uprsync {
  return &uprsync{}
}

func (auth *uprsync) CheckCredentials(username string, password string) (error) {
  url := fmt.Sprintf("http://sync.upr.edu.cu/api/apilogin/%s/%s/undefined", username, password)
  response, err := http.Get(url)
  if err != nil {
    msg := fmt.Sprintf("http get error: %s", err.Error())
    return errors.New(msg)
  }
  if response.StatusCode == 400 {
    return errors.New("unauthorized 400")
  }
  if response.StatusCode != 200 {
    msg := fmt.Sprintf("http error: %d", response.StatusCode)
    return errors.New(msg)
  }
  return nil
}

func (auth *uprsync) CheckToken(token string) (error) {
  return CheckToken(token)
}

func (auth *uprsync) GenerateToken(username string, role string) (string) {
  return GenerateToken(username, role)
}

func (auth *uprsync) GetUsername(token string) (string, error) {
  return GetUsername(token)
}
