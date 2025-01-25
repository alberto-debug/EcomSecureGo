package db

import (
	"database/sql"
	"log"

	"github.com/go-sql-driver/mysql"
)

func MySQLStorage(cfg mysql.Config) (*sql.DB, error) {
	db, err := sql.Open("mysql", cfg.FormatDSN())
	
		log.Fatal(db)
	}

	return db, nil
}
