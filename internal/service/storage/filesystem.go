package storage

import (
  "os"
  "log"
  "errors"
  "encoding/base64"
  "strings"
  "io/ioutil"
)

type fileSystem struct {
  rootPath string `json:"root_path"`
}

func newFileSystem(rootPath string) *fileSystem {
  return &fileSystem{
    rootPath: rootPath,
  }
}

func (fs *fileSystem) GetAll() ([][]byte, error) {
  // files, err = ioutil.ReadDir(fs.rootPath)
  return [][]byte{}, errors.New("not implemented")
}

func (fs *fileSystem) GetImage(username string) ([]byte, error) {
  file, err := os.Open("avatars/" + username)
  if err != nil {
    return []byte{}, err
  }
  defer file.Close()
  data, err := ioutil.ReadAll(file)
  return data, err
}

// image -> base64
func (fs *fileSystem) SaveImage(username string, imageBase64 string) (error) {
  comma := strings.Index(string(imageBase64), ",")
  info := string(imageBase64)[:comma]
  imageRaw := string(imageBase64)[comma+1:]
  imageOK, err := base64.StdEncoding.DecodeString(imageRaw)

  log.Println("info: ", info)
  log.Println("data: ", imageRaw)

  if err != nil {
    log.Println("error decoding: ", err.Error())
    return err
  }

  var file *os.File
  defer file.Close()

  path := fs.rootPath + "/" + username
  if _, err = os.Stat(path); err != nil {
    file, err = os.Create(path)
    if err != nil {
      log.Println("error opening file: ", err.Error())
      return err
    }
  } else {
    file, err = os.OpenFile(path, os.O_RDWR, 0777)
    if err != nil {
      log.Println("error opening file: ", err.Error())
      return err
    }
  }

  file.Write(imageOK)

  return err
}
