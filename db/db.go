package db

import (
	"database/sql"

	"github.com/go-sql-driver/mysql"
)

func MySQLStorage(cfg mysql.Config) (*sql.DB, error) {
}
