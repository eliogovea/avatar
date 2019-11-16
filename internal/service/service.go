package service

import (
  "github.com/eliogovea/upr-profile-pic/internal/service/authorizer"
  "github.com/eliogovea/upr-profile-pic/internal/service/storage"
)

type Service struct {
  Authorizer authorizer.Authorizer
  Storage    storage.Storage
}

func New() *Service {
  return &Service{
    Authorizer: authorizer.New("ldap"),
    Storage:    storage.New("avatars"),
  }
}
