package controllers

import (
	"database/sql"

	"github.com/mpuzanov/bill18Go/models"
	log "github.com/sirupsen/logrus"
)

//GetAllStreets Возвращаем список улиц
func (db *DB) GetAllStreets() ([]*models.Street, error) {
	const querySQL = `
	k_show_streets
	`
	rows, err := db.Query(querySQL)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	log.Traceln(querySQL)

	slice := []*models.Street{}
	for rows.Next() {
		zap := models.Street{}
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

//GetBuilds Возвращаем список домов по заданной улице
func (db *DB) GetBuilds(streetName string) (*models.Builds, error) {
	rows, err := db.Query("k_show_build @street_name1",
		sql.Named("street_name1", streetName))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	builds := models.Builds{}
	builds.StreetName = streetName
	for rows.Next() {
		var build models.Build
		build.Street.StreetName = streetName
		err := rows.Scan(&build.NomDom, &build.NomDomSort)
		if err != nil {
			return nil, err
		}
		builds.DataBuilds = append(builds.DataBuilds, build)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return &builds, nil
}

//GetFlats Возвращаем с писок квартир по заданной улице
func (db *DB) GetFlats(streetName, nomDom string) (*models.Flats, error) {
	rows, err := db.Query("k_show_kvr @street_name1,@nom_dom1",
		sql.Named("street_name1", streetName),
		sql.Named("nom_dom1", nomDom))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	flats := new(models.Flats)
	flats.StreetName = streetName
	flats.NomDom = nomDom
	for rows.Next() {
		var flat models.Flat
		flat.StreetName = streetName
		flat.NomDom = nomDom
		err := rows.Scan(&flat.NomKvr, &flat.NomKvrSort)
		if err != nil {
			return nil, err
		}
		flats.DataKvr = append(flats.DataKvr, flat)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return flats, nil
}

//GetKvrLic Выдаём список лицевых по заданному адресу(улица,дом,квартира)
func (db *DB) GetKvrLic(streetName, nomDom, nomKvr string) (*models.Lics, error) {
	rows, err := db.Query("k_show_occ_adres @street_name1, @nom_dom1, @nom_kvr1",
		sql.Named("street_name1", streetName),
		sql.Named("nom_dom1", nomDom),
		sql.Named("nom_kvr1", nomKvr))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	lics := new(models.Lics)
	lics.StreetName = streetName
	lics.NomDom = nomDom
	lics.NomKvr = nomKvr
	for rows.Next() {
		var lic models.Lic
		err := rows.Scan(&lic.Occ)
		if err != nil {
			return nil, err
		}
		lics.DataKvrLic = append(lics.DataKvrLic, lic)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return lics, nil
}

//GetCounterByOcc Выводим все ПУ по заданному лицевому счёту
func (db *DB) GetCounterByOcc(occ int) ([]*models.DataCounter, error) {

	//k_show_counters @occ=:occ1
	const querySQL = `
		k_show_counters @occ
		`
	rows, err := db.Query(querySQL,
		sql.Named("occ", occ),
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	slice := []*models.DataCounter{}
	for rows.Next() {
		zap := models.DataCounter{}
		err := rows.Scan(&zap.Occ,
			&zap.CounterID,
			&zap.ServName,
			&zap.SerialNumber,
			&zap.CounterType,
			&zap.MaxValue,
			&zap.UnitID,
			&zap.CountValue,
			&zap.DateCreate,
			&zap.PeriodCheck,
			&zap.ValueDate,
			&zap.LastValue,
			&zap.ActualValue,
			&zap.AvgMonth,
			&zap.Tarif,
			&zap.NormaSingle,
			&zap.AvgItog,
			&zap.KolNorma,
		)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
		//fmt.Println(zap.ToString())
		slice = append(slice, &zap)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
		return nil, err
	}
	return slice, nil
}

//GetDataValueByOcc Начисления по лицевому счёту
func (db *DB) GetDataValueByOcc(occ int) ([]*models.DataValue, error) {
	const querySQL = `
	k_show_values_occ @occ=@occ1, @row1=12
	`
	rows, err := db.Query(querySQL,
		sql.Named("occ1", occ),
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	slice := []*models.DataValue{}
	for rows.Next() {
		zap := models.DataValue{}
		err := rows.Scan(&zap.FinStr,
			&zap.Occ,
			&zap.Saldo,
			&zap.Value,
			&zap.Added,
			&zap.Paid,
			&zap.Paymaccount,
			&zap.PaymaccountPeny,
			&zap.PaymaccountServ,
			&zap.Debt,
			&zap.SupName,
		)
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

//GetDataPaymByOcc Список платежей по лицевому счёту
func (db *DB) GetDataPaymByOcc(occ int) ([]*models.DataPaym, error) {
	const querySQL = `
	k_show_payings @occ=@occ1
`
	rows, err := db.Query(querySQL,
		sql.Named("occ1", occ),
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	slice := []*models.DataPaym{}
	for rows.Next() {
		zap := models.DataPaym{}
		err := rows.Scan(&zap.FinStr,
			&zap.Occ,
			&zap.Date,
			&zap.Summa,
			&zap.PaymaccountPeny,
			&zap.SupName,
		)
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

// GetDataOcc Выдаём информацию по заданному лицевому счёту
func (db *DB) GetDataOcc(occ int) (*models.DataOcc, error) {
	row := db.QueryRow("k_show_occ @occ=@occ1",
		sql.Named("occ1", occ))

	data := models.DataOcc{}
	err := row.Scan(&data.Occ,
		&data.BasaName,
		&data.Address,
		&data.TipName,
		&data.TotalSq,
		&data.OccSup,
		&data.FinCurrent,
		&data.FinCurrentStr,
		&data.KolPeople,
		&data.CV1,
		&data.CV2,
		&data.Rejim)
	if err == sql.ErrNoRows {
		return &data, nil
	} else if err != nil {
		log.Fatal(err)
	}
	return &data, nil
}

//GetCounterValueByOcc Показания ПУ по лицевому счёту
func (db *DB) GetCounterValueByOcc(occ int) ([]*models.CounterValue, error) {
	const querySQL = `
	k_show_counters_value @occ=@occ1, @counter_id=null, @row1=@kolval
`
	kolval := 6
	rows, err := db.Query(querySQL,
		sql.Named("occ1", occ),
		sql.Named("kolval", kolval),
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	slice := []*models.CounterValue{}
	for rows.Next() {
		zap := models.CounterValue{}
		err := rows.Scan(&zap.Occ,
			&zap.CounterID,
			&zap.InspectorDate,
			&zap.InspectorValue,
			&zap.ActualValue,
			&zap.FinStr,
			&zap.ID,
			&zap.SerialNumber,
			&zap.ServName,
			&zap.FinID,
			&zap.Sysuser,
		)
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
