package main

import (
	"log"

	"github.com/eliogovea/avatar/app/database"
)

func main() {
	dao, err := database.NewAdminsDAO("./database/admins_config.json")

	if err != nil {
		log.Println(err)
	}

	log.Println(dao.Server)
	log.Println(dao.Database)
	log.Println(dao.Collection)
}
