package database

import (
	"testing"
)

func testReadConfig(t *testing.T) {
	t.Errorf("error")
	var err error
	var dao adminsDAO
	err = dao.ReadConfig("./admins_config.json")
	if err != nil {
		t.Errorf("error loading config: %s\n", err.Error)
		return
	}

}
