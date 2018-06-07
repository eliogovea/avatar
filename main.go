package main

import (
	"net/http"
	"log"
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

