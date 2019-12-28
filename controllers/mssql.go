package controllers

import (
	"database/sql"

	"github.com/mpuzanov/bill18Go/models"
	log "github.com/sirupsen/logrus"
)

//GetAllStreets Возвращаем список улиц
func (db *DB) GetAllStreets() (*[]models.Street, error) {
	const querySQL = `k_show_streets`
	data := []models.Street{}
	err := db.Select(&data, querySQL)
	if err == sql.ErrNoRows {
		return &data, nil
	}
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return &data, nil
}

//GetBuilds Возвращаем список домов по заданной улице
func (db *DB) GetBuilds(streetName string) (*models.Builds, error) {

	builds := models.Builds{}
	builds.StreetName = streetName
	err := db.Select(&builds.DataBuilds, "k_show_build @street_name1", sql.Named("street_name1", streetName))
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return &builds, nil
}

//GetFlats Возвращаем список квартир по заданной улице
func (db *DB) GetFlats(streetName, nomDom string) (*models.Flats, error) {

	flats := models.Flats{}
	flats.StreetName = streetName
	flats.NomDom = nomDom
	err := db.Select(&flats.DataKvr, "k_show_kvr @street_name1,@nom_dom1",
		sql.Named("street_name1", streetName),
		sql.Named("nom_dom1", nomDom),
	)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return &flats, nil
}

//GetKvrLic Выдаём список лицевых по заданному адресу(улица,дом,квартира)
func (db *DB) GetKvrLic(streetName, nomDom, nomKvr string) (*models.Lics, error) {

	lics := models.Lics{}
	lics.StreetName = streetName
	lics.NomDom = nomDom
	lics.NomKvr = nomKvr
	err := db.Select(&lics.DataKvrLic, "k_show_occ_adres @street_name1, @nom_dom1, @nom_kvr1",
		sql.Named("street_name1", streetName),
		sql.Named("nom_dom1", nomDom),
		sql.Named("nom_kvr1", nomKvr),
	)
	if err == sql.ErrNoRows {
		return &lics, nil
	}
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return &lics, nil
}

//GetCounterByOcc Выводим все ПУ по заданному лицевому счёту
func (db *DB) GetCounterByOcc(occ int) (*[]models.DataCounter, error) {
	const querySQL = `k_show_counters @occ`

	data := []models.DataCounter{}
	err := db.Select(&data, querySQL,
		sql.Named("occ", occ),
	)
	if err == sql.ErrNoRows {
		return &data, nil
	}
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return &data, nil
}

//GetDataValueByOcc Начисления по лицевому счёту
func (db *DB) GetDataValueByOcc(occ int) (*[]models.DataValue, error) {
	const querySQL = `k_show_values_occ @occ=@occ1, @row1=12`

	data := []models.DataValue{}
	err := db.Select(&data, querySQL,
		sql.Named("occ1", occ),
	)
	if err == sql.ErrNoRows {
		return &data, nil
	}
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return &data, nil
}

//GetDataPaymByOcc Список платежей по лицевому счёту
func (db *DB) GetDataPaymByOcc(occ int) (*[]models.DataPaym, error) {
	const querySQL = `k_show_payings @occ=@occ1`

	data := []models.DataPaym{}
	err := db.Select(&data, querySQL,
		sql.Named("occ1", occ),
	)
	if err == sql.ErrNoRows {
		return &data, nil
	}
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return &data, nil
}

// GetDataOcc Выдаём информацию по заданному лицевому счёту
func (db *DB) GetDataOcc(occ int) (*models.DataOcc, error) {
	const querySQL = `k_show_occ @occ=@occ1`

	data := models.DataOcc{}
	err := db.Get(&data, querySQL, sql.Named("occ1", occ))
	if err == sql.ErrNoRows {
		return &data, nil
	}
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return &data, nil
}

//GetCounterValueByOcc Показания ПУ по лицевому счёту
func (db *DB) GetCounterValueByOcc(occ int) (*[]models.CounterValue, error) {
	const querySQL = `k_show_counters_value @occ=@occ1, @counter_id=null, @row1=@kolval`
	kolval := 6

	data := []models.CounterValue{}
	err := db.Select(&data, querySQL,
		sql.Named("occ1", occ),
		sql.Named("kolval", kolval),
	)
	if err == sql.ErrNoRows {
		return &data, nil
	}
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return &data, nil
}
