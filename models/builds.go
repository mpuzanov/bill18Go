package models

import (
	"database/sql"
)

//Build Дом
type Build struct {
	Street     `json:"-"`
	NomDom     string `json:"nom_dom"`
	NomDomSort string `json:"nom_dom_sort"`
}

//Builds Дома
type Builds struct {
	StreetName string  `json:"street_name,omitempty"`
	DataBuilds []Build `json:"dataBuilds,omitempty"`
}

//GetBuilds Возвращаем список домов по заданной улице
func GetBuilds(streetName string) (*Builds, error) {
	rows, err := db.Query("k_show_build @street_name1",
		sql.Named("street_name1", streetName))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	builds := new(Builds)
	builds.StreetName = streetName
	for rows.Next() {
		var build Build
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
	return builds, nil
}
