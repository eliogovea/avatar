package main

import (
	"os"
	"log"
)

type storage struct {
	DefaultAvatar		string	`json:"default_avatar"`
	AvatarsDirectory 	string	`json:"avatars_directory"`
	ApprovedDirectory	string	`json:"approved_directory"`
	PendingDirectory	string	`json:"pending_directory"`
}

func (fs *storage)exists(fileName string) error {
	_, err := os.Stat(fileName)
	return err
}

func (fs *storage)getPending(fileName string) string {
	// TODO
	return ""
}

func (fs *storage)getApproved(username string) string {
	fileName := fs.ApprovedDirectory + "/" + username + ".jpeg"
	log.Println("query: ", fileName)
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
