package database

import (
	"encoding/json"
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

func NewAdminsDAO(path string, session *mgo.Session) (*adminsDAO, error) {
	dao := new(adminsDAO)
	err := dao.ReadConfig(path)
	if err != nil {
		return dao, err
	}
	dao.C = session.DB(dao.Database).C(dao.Collection)
	// err = dao.Connect()
	return dao, err
}

func (dao *adminsDAO) ReadConfig(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	return json.NewDecoder(file).Decode(dao)
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

func (dao *adminsDAO) FindByUsername(username string) (model.Admin, error) {
	var admin model.Admin
	err := dao.C.Find(bson.M{"username": username}).One(&admin)
	return admin, err
}

func (dao *adminsDAO) Insert(admin *model.Admin) error {
	return dao.C.Insert(admin)
}

func (dao *adminsDAO) Remove(admin *model.Admin) error {
	return dao.C.Remove(admin)
}

func (dao *adminsDAO) RemoveByUsername(username string) error {
	admin, err := dao.FindByUsername(username)
	if err != nil {
		return err
	}
	return dao.Remove(&admin)
}
