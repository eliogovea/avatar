package database

import (
	"encoding/json"
	"errors"
	"os"

	"github.com/eliogovea/avatar/app/model"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type avatarsDAO struct {
	C          *mgo.Collection
	Server     string `json:"server"`
	Database   string `json:"database"`
	Collection string `json:"collection"`
}

func (dao *avatarsDAO) ReadConfig(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	err = json.NewDecoder(file).Decode(dao)
	return dao.Connect()
}

func (dao *avatarsDAO) Connect() error {
	session, err := mgo.Dial(dao.Server)
	if err != nil {
		return err
	}
	dao.C = session.DB(dao.Database).C(dao.Collection)

	return nil
}

func (dao *avatarsDAO) FindAll() ([]model.Avatar, error) {
	var avatars []model.Avatar
	err := dao.C.Find(bson.M{}).All(&avatars)
	return avatars, err
}

func (dao *avatarsDAO) FindById(id string) (model.Avatar, error) {
	var avatar model.Avatar
	err := dao.C.FindId(bson.ObjectIdHex(id)).One(&avatar)
	return avatar, err
}

func (dao *avatarsDAO) FindAllApproved() ([]model.Avatar, error) {
	var avatars []model.Avatar
	err := dao.C.Find(bson.M{"status": "approved"}).All(&avatars)
	return avatars, err
}

func (dao *avatarsDAO) FindApproved(username string) (model.Avatar, error) {
	var avatar model.Avatar
	err := dao.C.Find(bson.M{"username": "approved"}).Select(bson.M{"status": "approved"}).One(&avatar)
	return avatar, err
}

func (dao *avatarsDAO) FindAllPending() ([]model.Avatar, error) {
	var avatars []model.Avatar
	err := dao.C.Find(bson.M{"status": "pending"}).All(&avatars)
	return avatars, err
}

func (dao *avatarsDAO) FindPending(username string) (model.Avatar, error) {
	var avatar model.Avatar
	err := dao.C.Find(bson.M{"username": username}).Select(bson.M{"status": "pending"}).One(&avatar)
	return avatar, err
}

// Insert
func (dao *avatarsDAO) Insert(avatar *model.Avatar) error {
	return dao.C.Insert(avatar)
}

func (dao *avatarsDAO) Remove(avatar *model.Avatar) error {
	return dao.C.Remove(avatar)
}

func (dao *avatarsDAO) Update(avatar *model.Avatar) error {
	return dao.C.UpdateId(avatar.ID, &avatar)
}

func (dao *avatarsDAO) Approve(avatar *model.Avatar) error {
	if avatar.Status != "pending" {
		return errors.New("incorrect")
	}
	// TODO
	return nil
}

// TODO
// approved avatar
// if the approved exists delete it
