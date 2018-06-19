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
	go func() {
		http.HandleFunc("/", s.getApproved())
		err := http.ListenAndServe(":8080", nil)
		if err != nil {
			log.Fatal("listen and serve http:", err)
		}
	}()
	/*
		err = http.ListenAndServe(":8080", s.router)
		if err != nil {
			log.Fatal("can't start: ", err)
		}
	*/
	err = http.ListenAndServeTLS(":443", "cert.pem", "key.pem", s.router)
	if err != nil {
		log.Fatal("Listen and serve https: ", err)
	}
}

func main() {
	// test()
	testWithLoad()
}
