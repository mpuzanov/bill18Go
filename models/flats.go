package models

import (
	"database/sql"
)

//Flat Квартира
type Flat struct {
	Build      `json:"-"`
	NomKvr     string `json:"nom_kvr"`
	NomKvrSort string `json:"nom_kvr_sort"`
}

//Flats Квартиры
type Flats struct {
	StreetName string `json:"street_name,omitempty"`
	NomDom     string `json:"nom_dom,omitempty"`
	DataKvr    []Flat `json:"dataKvr,omitempty"`
}

//GetFlats Возвращаем с писок квартир по заданной улице
func (db *DB) GetFlats(streetName, nomDom string) (*Flats, error) {
	rows, err := db.Query("k_show_kvr @street_name1,@nom_dom1",
		sql.Named("street_name1", streetName),
		sql.Named("nom_dom1", nomDom))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	flats := new(Flats)
	flats.StreetName = streetName
	flats.NomDom = nomDom
	for rows.Next() {
		var flat Flat
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
