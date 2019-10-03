package models

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
