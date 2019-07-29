package models

import (
	"database/sql"
	"log"
)

var db *sql.DB

// InitDB Создание соединения с базой данных
func InitDB(dataSourceName string) {
	var err error
	db, err = sql.Open("sqlserver", dataSourceName)
	if err != nil {
		log.Panic(err)
	}

	if err = db.Ping(); err != nil {
		log.Panic(err)
	}
}
