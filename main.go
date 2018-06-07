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

	http.ListenAndServe(s.Address, s.router)
}

func main() {
	// test()
	testWithLoad()
}

