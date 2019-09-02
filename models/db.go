package models

import (
	"database/sql"
)

//Datastore  интерфейс Datastore реализовывает некоторые методы, в нашем типе DB
type Datastore interface {
	GetAllStreets() ([]*Street, error)
	GetBuilds(streetName string) (*Builds, error)
	GetFlats(streetName, nomDom string) (*Flats, error)
	GetKvrLic(streetName, nomDom, nomKvr string) (*Lics, error)
	GetCounterByOcc(occ int) ([]*DataCounter, error)
	GetCounterValueByOcc(occ int) ([]*CounterValue, error)
	GetDataValueByOcc(occ int) ([]*DataValue, error)
	GetDataPaymByOcc(occ int) ([]*DataPaym, error)
	GetDataOcc(occ int) (*DataOcc, error)
}

//DB пользовательский тип БД
type DB struct {
	*sql.DB
}

//GetInitDB Создание соединения с базой данных
func GetInitDB(connString string) (*DB, error) {

	db, err := sql.Open("sqlserver", connString)
	if err != nil {
		return nil, err
	}
	return &DB{db}, err
}
