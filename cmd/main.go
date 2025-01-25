package main

import (
	"log"

	"github.com/go-sql-driver/mysql"

	"github.com/alberto-debug/EcomSecureGo/cmd/api"
	"github.com/alberto-debug/EcomSecureGo/db"
)

func main() {
	db, err := db.MySQLStorage(mysql.Config{
		User:                 "root",
		Passwd:               "",
		Addr:                 "127.0.1:3306",
		DBName:               "ecom",
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	})

	server := api.NEWAPIServer(":8080", nil)

	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
