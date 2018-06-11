package main

import (
	"io"
	"os"
	"log"
 	"net/http"
)

type storage struct {
	DefaultAvatar		string	`json:"default_avatar"`
	AvatarsDirectory 	string	`json:"avatars_directory"`
	ApprovedDirectory	string	`json:"approved_directory"`
	PendingDirectory	string	`json:"pending_directory"`
	MaxUploadSize		int32	`json:"max_upload_size"`
	AllowedTypes		[]string `json:"allowed_types"`
}



func (fs *storage)copyFile(from io.Reader, path string) error {
	// TODO check if the path is valid

	// TODO check if the file type is valid


	buffer := make([]byte, fs.MaxUploadSize)
	_, err := from.Read(buffer)

	if err != nil {
		return err
	}

	contentType := http.DetectContentType(buffer)

	ok := false
	for _, v := range(fs.AllowedTypes) {
		if v == contentType {
			ok = true
			break
		}
	}

	if !ok {
		log.Println("file type not allowed: ", err)
		return err
	}

	log.Println(contentType, "OK")

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

func (fs *storage)HasPending(username string) bool {
	_, err := os.Stat(fs.PendingDirectory + "/" + username)
	if err != nil {
		return false;
	}
	return true
}

func (fs *storage)HasApproved(username string) bool {
	_, err := os.Stat(fs.ApprovedDirectory + "/" + username)
	if err != nil {
		return false
	}
	return true
}

func (fs *storage)CreatePending(from io.Reader, username string) error {
	return fs.copyFile(from, fs.PendingDirectory + "/" + username)
}

func (fs *storage)exists(fileName string) error {
	_, err := os.Stat(fileName)
	return err
}

func (fs *storage)getPending(username string) string {
	fileName := fs.PendingDirectory + "/" + username
	if err := fs.exists(fileName); err != nil {
		if !os.IsNotExist(err) {
			log.Println("unknown error: ", err)
		}
		return fs.DefaultAvatar
	}
	return fileName
}

func (fs *storage)getApproved(username string) string {
	fileName := fs.ApprovedDirectory + "/" + username
	if err := fs.exists(fileName); err != nil {
		if !os.IsNotExist(err) {
			log.Println("unknown error: ", err)
		}
		return fs.DefaultAvatar
	}
	return fileName
}

func deletePending(username string) error {
	// TODO
	return nil
}

func deleteApproved(username string) error {
	// TODO
	return nil
}

func movePendignToApproved(username string) error {
	// TODO
	return nil
}
