package models

//Lic Лицевой
type Lic struct {
	Occ int `json:"occ" db:"occ"`
}

//Lics Лицевые
type Lics struct {
	StreetName string `json:"street_name,omitempty" db:"street_name"`
	NomDom     string `json:"nom_dom,omitempty" db:"nom_dom"`
	NomKvr     string `json:"nom_kvr,omitempty" db:"nom_kvr"`
	DataKvrLic []Lic  `json:"dataKvrLic,omitempty"`
}
