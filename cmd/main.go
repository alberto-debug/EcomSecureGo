package main

import (
	"log"

	"github.com/go-sql-driver/mysql"

	"github.com/alberto-debug/EcomSecureGo/cmd/api"
	"github.com/alberto-debug/EcomSecureGo/config"
	"github.com/alberto-debug/EcomSecureGo/db"
)

func main() {
	db, err := db.MySQLStorage(mysql.Config{
		User:                 config.Envs.DBName,
		Passwd:               config.Envs.DBPassword,
		Addr:                 config.Envs.DBAddress,
		DBName:               config.Envs.DBName,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	})
	if err != nil {
		log.Fatal(err)
	}

	server := api.NEWAPIServer(":8080", db)

	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
