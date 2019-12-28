package controllers

import (
	"github.com/jmoiron/sqlx"
	"github.com/mpuzanov/bill18Go/models"
)

//Datastore  интерфейс Datastore реализовывает некоторые методы, в нашем типе DB
type Datastore interface {
	GetAllStreets() (*[]models.Street, error)
	GetBuilds(streetName string) (*models.Builds, error)
	GetFlats(streetName, nomDom string) (*models.Flats, error)
	GetKvrLic(streetName, nomDom, nomKvr string) (*models.Lics, error)
	GetCounterByOcc(occ int) (*[]models.DataCounter, error)
	GetCounterValueByOcc(occ int) (*[]models.CounterValue, error)
	GetDataValueByOcc(occ int) (*[]models.DataValue, error)
	GetDataPaymByOcc(occ int) (*[]models.DataPaym, error)
	GetDataOcc(occ int) (*models.DataOcc, error)
}

//DB пользовательский тип БД
type DB struct {
	//*sql.DB
	*sqlx.DB
}

//GetInitDB Создание соединения с базой данных
func GetInitDB(connString string) (*DB, error) {

	//db, err := sql.Open("sqlserver", connString)
	db, err := sqlx.Open("sqlserver", connString)
	if err != nil {
		return nil, err
	}
	return &DB{db}, err
}
