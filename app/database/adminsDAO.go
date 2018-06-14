package database

import (
	"os"

	"github.com/eliogovea/avatar/app/model"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type adminsDAO struct {
	C          *mgo.Collection
	Server     string `json:"server"`
	Database   string `json:"database"`
	Collection string `json:"collection"`
}

func (dao *adminsDAO) ReadConfig(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	return dao.Connect()
}

func (dao *adminsDAO) Connect() error {
	session, err := mgo.Dial(dao.Server)
	if err != nil {
		return err
	}
	dao.C = session.DB(dao.Database).C(dao.Collection)
	return nil
}

func (dao *adminsDAO) FindAll() ([]model.Admin, error) {
	var admins []model.Admin
	err := dao.C.Find(bson.M{}).All(&admins)
	return admins, err
}

func (dao *adminDAO) FindByUsername(username string) (model.Admin, error) {
	var admin model.Admin
	err := dao.C.Find(bson.M{"username": username}).One(&admin)
	return admin, err
}

func (dao *adminsDAO) Insert(admin *model.Admin) error {
	return dao.C.Insert(admin)
}

func (dao *adminsDAO) Remove(admin *model.Admin) error {
	return dao.C.Remove(avatar)
}

func (dao *adminsDAO) Remove(username string) error {
	admin, err := dao.FindByUsername(username)
	if err != nil {
		return err
	}
	return dao.Remove(&admin)
}
