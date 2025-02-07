package main

import (
	"log"
	"net/http"
)

type Middleware func(http.HandlerFunc) http.HandlerFunc

var middlewares = []Middleware{
	TokenAuthMiddleware,
}

func main() {
	var handler http.HandlerFunc = handleClientProfile
	for _, middleware := range middlewares {
		handler = middleware(handler)
	}

	http.HandleFunc("/user/profile", handleClientProfile)

	log.Println("Server started at :8010")
	log.Fatal(http.ListenAndServe(":8010", nil))
}
