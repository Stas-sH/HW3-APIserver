package main

import (
	"Stas-sH/testclientserver/server"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/users", server.WeryffyMiddleware(server.LoggerMiddleware(server.HandleUsers)))
	http.HandleFunc("/users/user", server.WeryffyMiddleware(server.LoggerMiddleware(server.HandleUser)))
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}

}
