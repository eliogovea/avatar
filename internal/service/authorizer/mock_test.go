package authorizer

import (
  "testing"
)

func TestMock(t *testing.T) {
  auth := newMock()
  if err := auth.CheckCredentials("123", "234"); err != nil {
    t.Error("unexpected error: ", err.Error())
  }
  if err := auth.CheckCredentials("123", "123"); err == nil {
    t.Error("expected some error")
  }
}
