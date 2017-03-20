package main

import (
	"log"
	"net/http"
	"SimpleChat/routers"
)

func main() {
	router := routers.NewRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}
