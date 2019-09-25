package models

import (
	"fmt"

	log "github.com/sirupsen/logrus"
)

//Street Улица
type Street struct {
	StreetName string `json:"name"`
}

//ToString Строковое представление Платежа
func (zap *Street) ToString() string {
	return fmt.Sprintf("%s", zap.StreetName)
}

//GetAllStreets Возвращаем список улиц
func (db *DB) GetAllStreets() ([]*Street, error) {
	const querySQL = `
	k_show_streets
	`
	rows, err := db.Query(querySQL)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	log.Traceln(querySQL)

	slice := []*Street{}
	for rows.Next() {
		zap := Street{}
		err := rows.Scan(&zap.StreetName)
		if err != nil {
			return nil, err
		}
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
		slice = append(slice, &zap)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
		return nil, err
	}
	return slice, nil
}
