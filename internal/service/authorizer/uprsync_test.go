package authorizer

import (
  "testing"
)

func TestUPRSyncWrong(t *testing.T) {
  auth := newUprsync()
  err := auth.CheckCredentials("123", "234")
  if err == nil {
    t.Error("bad credentials accepted")
  }
}
