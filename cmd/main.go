package main

import "github.com/alberto-debug/EcomSecureGo/cmd/api"

func main() {
	server := api.NEWAPIServer(":8080", nil)
	server.Run()
}
