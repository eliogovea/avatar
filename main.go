package main

import (
	"net/http"
	"time"
	"os"
	"log"

	"github.com/eliogovea/avatar/auth"
	"github.com/eliogovea/avatar/session"
)

func testWithLoad() {
	s, err := loadFromConfig("./config.json")
	if err != nil {
		log.Println("!!! error loading the configuration", err)
		return
	}
	s.buildHandlers()

	s.router.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.ListenAndServe(s.Address, s.router)
}

func main() {
	// test()
	testWithLoad()
}

