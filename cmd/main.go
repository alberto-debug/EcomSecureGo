package main

import (
	"log"

	"github.com/alberto-debug/EcomSecureGo/cmd/api"
)

func main() {
	server := api.NEWAPIServer(":8080", nil)

	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
