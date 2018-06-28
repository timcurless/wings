package flightservice

import (
	"log"
	"net/http"
)

func StartService(port string) {

	r := NewRouter()
	http.Handle("/", r)

	log.Println("Starting Flight Service on port 8080")
	err := http.ListenAndServe(":"+port, nil)

	if err != nil {
		panic(err)
	}
}
