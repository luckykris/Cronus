package db

import "database/sql"
import _ "github.com/go-sql-driver/mysql"

func startMysql()
	db, err := sql.Open("mysql", "user:password@/dbname")