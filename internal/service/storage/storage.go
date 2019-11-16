package storage

import (
  "errors"
)

var ErrNotFound = errors.New("not found")

type Storage interface {
  GetAll() ([][]byte, error)
  GetImage(/* username */ string) ([]byte, error)
  SaveImage(/* username */ string, /* base64Image */ string) (error)
}

func New(rootPath string) Storage {
  return newFileSystem(rootPath)
}
