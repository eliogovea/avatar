package database

import "testing"

func TestReadConfig(t *testing.T) {
	path := "./admins_config.json"
	dao := new(adminsDAO)
	err := dao.ReadConfig(path)
	if err != nil {
		t.Error(err)
	}
	if dao.Server != "localhost" {
		t.Error("wrong server")
	}
	if dao.Database != "database" {
		t.Error("wrong database")
	}
	if dao.Collection != "admins" {
		t.Error("wrong collection")
	}
}

func TestConnect(t *testing.T) {
	path := "./admins_config.json"
	dao := new(adminsDAO)
	dao.ReadConfig(path)
	if dao.Connect() != nil {
		t.Error("error connecting to database")
	}
}
