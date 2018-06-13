package main

import (
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type storage struct {
	DefaultAvatar     string   `json:"default_avatar"`
	AvatarsDirectory  string   `json:"avatars_directory"`
	ApprovedDirectory string   `json:"approved_directory"`
	PendingDirectory  string   `json:"pending_directory"`
	MaxUploadSize     int32    `json:"max_upload_size"`
	AllowedTypes      []string `json:"allowed_types"`

	Approved map[string]bool
	Pending  map[string]bool
}

func NewStorage(config string) *storage {
	file, err := os.Open(config)
	if err != nil {
		log.Fatalf("can't open storage config")
	}
	fs := new(storage)
	err = json.NewDecoder(file).Decode(fs)
	if err != nil {
		log.Fatalf("error reading storage config")
	}

	fs.Approved = make(map[string]bool)
	fs.Pending = make(map[string]bool)

	fs.loadAvatars()

	return fs
}

func (fs *storage) loadAvatars() error {
	fs.Approved = make(map[string]bool)
	fs.Pending = make(map[string]bool)
	files, err := ioutil.ReadDir(fs.ApprovedDirectory)
	if err != nil {
		return errors.New("can't read files on " + fs.ApprovedDirectory)
	}
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		// log.Println("approved", file.Name())
		fs.Approved[file.Name()] = true
	}
	files, err = ioutil.ReadDir(fs.PendingDirectory)
	if err != nil {
		return errors.New("can't read files on " + fs.PendingDirectory)
	}
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		// log.Println("pending ", file.Name())
		fs.Pending[file.Name()] = true
	}
	return nil
}

func (fs *storage) copyFile(from io.Reader, path string) error {
	// TODO check if the path is valid

	// TODO check if the file type is valid

	buffer := make([]byte, fs.MaxUploadSize)
	_, err := from.Read(buffer)

	if err != nil {
		return err
	}

	contentType := http.DetectContentType(buffer)

	ok := false
	// TODO user a map
	for _, v := range fs.AllowedTypes {
		if v == contentType {
			ok = true
			break
		}
	}

	if !ok {
		log.Println("file type not allowed: ", err)
		return err
	}

	var to *os.File
	if err := fs.exists(path); err != nil {
		to, err = os.Create(path)
		if err != nil {
			return err
		}
	} else {
		to, err = os.OpenFile(path, os.O_RDWR, 0777)
		if err != nil {
			return err
		}
	}

	_, err = to.Write(buffer)
	return err
}

func (fs *storage) approvePending(username string) error {
	pendingPath := fs.PendingDirectory + "/" + username
	approvedPath := fs.ApprovedDirectory + "/" + username

	if _, ok := fs.Pending[username]; !ok {
		return errors.New("no pending avatar")
	}

	err := os.Rename(pendingPath, approvedPath)

	if err != nil {
		fs.Approved[username] = true
		delete(fs.Pending, username)
	}

	return err
}

func (fs *storage) denyPending(username string) error {
	err := os.Remove(fs.PendingDirectory + "/" + username)
	if err != nil {
		delete(fs.Pending, username)
	}
	return err
}

func (fs *storage) denyApproved(username string) error {
	err := os.Remove(fs.PendingDirectory + "/" + username)
	if err != nil {
		delete(fs.Approved, username)
	}
	return err
}

func (fs *storage) HasPending(username string) bool {
	_, ok := fs.Pending[username]
	return ok
}

func (fs *storage) HasApproved(username string) bool {
	_, ok := fs.Approved[username]
	return ok
}

func (fs *storage) CreatePending(from io.Reader, username string) error {
	err := fs.copyFile(from, fs.PendingDirectory+"/"+username)
	if err != nil {
		fs.Pending[username] = true
		log.Println("added pending", username)
	}
	return err
}

func (fs *storage) exists(fileName string) error {
	_, err := os.Stat(fileName)
	return err
}

func (fs *storage) getPending(username string) string {
	_, ok := fs.Pending[username]
	if !ok {
		return fs.DefaultAvatar
	}
	return fs.PendingDirectory + "/" + username
}

func (fs *storage) getApproved(username string) string {
	_, ok := fs.Approved[username]
	if !ok {
		return fs.DefaultAvatar
	}
	return fs.ApprovedDirectory + "/" + username
}
