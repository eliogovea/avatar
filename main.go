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

	go startHTTP(s)

	err = http.ListenAndServeTLS(":443", "cert.pem", "key.pem", s.router)
	if err != nil {
		log.Fatal("Listen and serve https: ", err)
	}
}

func startHTTP(s *server) {
	http.HandleFunc("/", s.getApprovedAvatar())
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("listen and serve http:", err)
	}
}

func main() {
	// test()
	testWithLoad()
}
