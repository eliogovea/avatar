package authorizer

import (
  "log"
)

type Authorizer interface {
  CheckCredentials(/* username */ string, /* password */ string) (error)
  GenerateToken(/* username */ string, /* role */ string) (string)
  CheckToken(/* token */ string) (error)
  GetUsername(/* token */ string) (string, error)
}

func New(t string) Authorizer {
  if t == "mock" {
    return newMock()
  }
  if t == "ldap" {
    return newLDAP()
  }
  if t == "sync" {
    return newUprsync()
  }
  log.Fatal("wrong authorizer")
  return nil
}
