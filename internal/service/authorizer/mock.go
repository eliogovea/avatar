package authorizer

import (
  "log"
  "errors"
)

var errUnauthorized error = errors.New("unauthorized")

type mockAuthorizer struct {
}

func newMock() *mockAuthorizer {
  return &mockAuthorizer{}
}

func (auth *mockAuthorizer) CheckCredentials(username string, password string) (error) {
  log.Println("login: ", username, password)
  if (username == "123" && password == "234") {
    return nil
  }
  if (username == "456" && password == "567") {
    return nil
  }
  if (username == "frank.vigil" && password == "123") {
    return nil
  }
  return errUnauthorized
}

func (auth *mockAuthorizer) CheckToken(token string) (error) {
  return CheckToken(token)
}

func (auth *mockAuthorizer) GenerateToken(username string, role string) (string) {
  return GenerateToken(username, role)
}

func (auth *mockAuthorizer) GetUsername(token string) (string, error) {
  return GetUsername(token)
}
