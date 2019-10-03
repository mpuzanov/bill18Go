package models

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
