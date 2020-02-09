package controller

import (
	"database/sql"

	"github.com/mpuzanov/bill18Go/internal/model"
	log "github.com/sirupsen/logrus"
)

//GetAllStreets Возвращаем список улиц
func (db *DB) GetAllStreets() (*[]model.Street, error) {
	const querySQL = `k_show_streets`
	data := []model.Street{}
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
func (db *DB) GetBuilds(streetName string) (*model.Builds, error) {

	builds := model.Builds{}
	builds.StreetName = streetName

	params := map[string]interface{}{"street_name1": streetName}
	nstmt, err := db.PrepareNamed(`k_show_build @street_name1=:street_name1`)
	err = nstmt.Select(&builds.DataBuilds, params)
	//err := db.Select(&builds.DataBuilds, "k_show_build @street_name1", sql.Named("street_name1", streetName))
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return &builds, nil
}

//GetFlats Возвращаем список квартир по заданной улице
func (db *DB) GetFlats(streetName, nomDom string) (*model.Flats, error) {

	flats := model.Flats{}
	flats.StreetName = streetName
	flats.NomDom = nomDom

	params := map[string]interface{}{"street_name1": streetName, "nom_dom1": nomDom}
	nstmt, err := db.PrepareNamed(`k_show_kvr @street_name1=:street_name1, @nom_dom1=:nom_dom1`)
	err = nstmt.Select(&flats.DataKvr, params)

	// err := db.Select(&flats.DataKvr, "k_show_kvr @street_name1,@nom_dom1",
	// 	sql.Named("street_name1", streetName),
	// 	sql.Named("nom_dom1", nomDom),
	// )
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return &flats, nil
}

//GetKvrLic Выдаём список лицевых по заданному адресу(улица,дом,квартира)
func (db *DB) GetKvrLic(streetName, nomDom, nomKvr string) (*model.Lics, error) {

	lics := model.Lics{}
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
func (db *DB) GetCounterByOcc(occ int) (*[]model.DataCounter, error) {
	const querySQL = `k_show_counters @occ`

	data := []model.DataCounter{}
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
func (db *DB) GetDataValueByOcc(occ int) (*[]model.DataValue, error) {
	const querySQL = `k_show_values_occ @occ=@occ1, @row1=12`

	data := []model.DataValue{}
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
func (db *DB) GetDataPaymByOcc(occ int) (*[]model.DataPaym, error) {
	const querySQL = `k_show_payings @occ=@occ1`

	data := []model.DataPaym{}
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
func (db *DB) GetDataOcc(occ int) (*model.DataOcc, error) {
	const querySQL = `k_show_occ @occ=@occ1`

	data := model.DataOcc{}
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
func (db *DB) GetCounterValueByOcc(occ int) (*[]model.CounterValue, error) {
	const querySQL = `k_show_counters_value @occ=@occ1, @counter_id=null, @row1=@kolval`
	kolval := 6

	data := []model.CounterValue{}
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

//PuAddValue Ввод показания прибора учёта
func (db *DB) PuAddValue(puID int, value int) (*model.Result, error) {
	const querySQL = `k_pu_add_value @basa_name=@basa_name1, @counter_id1=@puID, @inspector_value1=@value`
	data := model.Result{}
	err := db.Get(&data, querySQL,
		sql.Named("basa_name1", "komp"),
		sql.Named("puID", puID),
		sql.Named("value", value),
	)
	log.Traceln("PuAddValue data:", &data)
	if err == sql.ErrNoRows {
		return &data, nil
	}
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return &data, nil
}

//PuDelValue Удаление заданного показания прибора учёта
func (db *DB) PuDelValue(puID int, id int) (*model.Result, error) {
	const querySQL = `k_pu_del_value @basa_name=@basa_name1, @counter_id1=@puID, @id1=@id`
	data := model.Result{}
	err := db.Get(&data, querySQL,
		sql.Named("basa_name1", "komp"),
		sql.Named("puID", puID),
		sql.Named("id", id),
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

//GetCounterValueByTip Показания ПУ по типу фонда
func (db *DB) GetCounterValueByTip(tipID int) (*[]model.CounterValueTip, error) {
	const querySQL = `k_show_counters_value_tip @basa_name=@basa_name1, @tip_id=@tip_id1`

	data := []model.CounterValueTip{}
	err := db.Select(&data, querySQL,
		sql.Named("basa_name1", "komp"),
		sql.Named("tip_id1", tipID),
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
