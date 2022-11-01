package config

import (
	"database/sql"
	"time"
)

func DB() *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/bs_pg1")
	if err != nil {
		panic(err)
	}
	//defer db.close()

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)

	return db
}
