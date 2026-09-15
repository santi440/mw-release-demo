package main

import (
	"fmt"
	"net/http"
)

func helloworld() string {
	return "Hello World!!"
}

func healthcheck() string {
	return "Health OK!"
}

func livenesscheck() string {
	return "I am alive!!!"
}

func MikrowaysDemoServer(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/healthz" {
		fmt.Fprint(w, healthcheck())
	}else if r.URL.Path == "/liveness" {
		fmt.Fprint(w, livenesscheck())
	} else {
		fmt.Fprint(w, helloworld())
	}
}
