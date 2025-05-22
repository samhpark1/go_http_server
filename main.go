package main

import (
	"log"

	"github.com/samhpark1/go_http_server/server"
)

func main() {
	srv := &server.Server{
		Addr: "5173",
	}

	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
