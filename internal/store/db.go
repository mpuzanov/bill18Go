package store

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/mpuzanov/bill18Go/internal/config"
	"github.com/mpuzanov/bill18Go/internal/errors"
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
	ctx context.Context
	db  *sqlx.DB
}

// NewStorageDB .
func NewStorageDB(cfg config.DBConf) (Datastore, error) {
	var err error
	var ds Datastore

	if cfg.Name == "" || cfg.User == "" || cfg.Host == "" || cfg.Password == "" || cfg.Database == "" {
		return nil, errors.ErrBadDBConfiguration
	}
	connString := fmt.Sprintf("%s://%s:%s@%s:%s?database=%s", cfg.Name, cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Database)

	ds, err = GetInitDB(connString)
	if err != nil {
		return nil, err
	}

	return ds, err
}

//GetInitDB Создание соединения с базой данных
func GetInitDB(connString string) (*DB, error) {

	db, err := sqlx.Open("sqlserver", connString)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	ctx := context.Background()
	return &DB{ctx: ctx, db: db}, err
}
