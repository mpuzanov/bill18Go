package models

import (
	"database/sql"
)

//Lic Лицевой
type Lic struct {
	Occ int `json:"occ"`
}

//Lics Лицевые
type Lics struct {
	StreetName string `json:"street_name,omitempty"`
	NomDom     string `json:"nom_dom,omitempty"`
	NomKvr     string `json:"nom_kvr,omitempty"`
	DataKvrLic []Lic  `json:"dataKvrLic,omitempty"`
}

//GetKvrLic Выдаём список лицевых по заданному адресу(улица,дом,квартира)
func (db *DB) GetKvrLic(streetName, nomDom, nomKvr string) (*Lics, error) {
	rows, err := db.Query("k_show_occ_adres @street_name1, @nom_dom1, @nom_kvr1",
		sql.Named("street_name1", streetName),
		sql.Named("nom_dom1", nomDom),
		sql.Named("nom_kvr1", nomKvr))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	lics := new(Lics)
	lics.StreetName = streetName
	lics.NomDom = nomDom
	lics.NomKvr = nomKvr
	for rows.Next() {
		var lic Lic
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
