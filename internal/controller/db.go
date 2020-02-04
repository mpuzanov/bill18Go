package controller

import (
	"github.com/jmoiron/sqlx"
	"github.com/mpuzanov/bill18Go/internal/model"
)

//Datastore  интерфейс Datastore реализовывает некоторые методы, в нашем типе DB
type Datastore interface {
	GetAllStreets() (*[]model.Street, error)
	GetBuilds(streetName string) (*model.Builds, error)
	GetFlats(streetName, nomDom string) (*model.Flats, error)
	GetKvrLic(streetName, nomDom, nomKvr string) (*model.Lics, error)
	GetCounterByOcc(occ int) (*[]model.DataCounter, error)
	GetCounterValueByOcc(occ int) (*[]model.CounterValue, error)
	GetDataValueByOcc(occ int) (*[]model.DataValue, error)
	GetDataPaymByOcc(occ int) (*[]model.DataPaym, error)
	GetDataOcc(occ int) (*model.DataOcc, error)
	PuAddValue(puID int, value int) (*model.Result, error)
	PuDelValue(puID int, id int) (*model.Result, error)
	GetCounterValueByTip(tipID int) (*[]model.CounterValueTip, error)
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
