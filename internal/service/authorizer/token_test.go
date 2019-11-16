package authorizer

import (
  "testing"
  "log"
)

func TestTokens(t *testing.T) {
  tokenString := GenerateToken("123", "admin")
  log.Println("token: ", tokenString)
  err := CheckToken(tokenString)
  if err != nil {
    t.Error("unexpected error: ", err.Error())
  }
  tokenString = tokenString + "a"
  err = CheckToken(tokenString)
  if err == nil {
    t.Error("expected an error")
  }
}
