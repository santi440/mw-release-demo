package main

import (
	"log"
	"net/http"
)

func main() {
	handler := http.HandlerFunc(MikrowaysDemoServer)
	log.Fatal(http.ListenAndServe(":8080", handler))
}
