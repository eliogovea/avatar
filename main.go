package main

import (
	"log"
	"net/http"
)

func testWithLoad() {
	s, err := loadFromConfig("./config.json")
	if err != nil {
		log.Println("!!! error loading the configuration", err)
		return
	}
	http.ListenAndServe(s.Address, s.router)
}

func main() {
	// test()
	testWithLoad()
}
