package model

//Flat Квартира
type Flat struct {
	Build      `json:"-"`
	NomKvr     string `json:"nom_kvr" db:"nom_kvr"`
	NomKvrSort string `json:"nom_kvr_sort" db:"nom_kvr_sort"`
}

//Flats Квартиры
type Flats struct {
	StreetName string `json:"street_name,omitempty" db:"street_name"`
	NomDom     string `json:"nom_dom,omitempty" db:"nom_dom"`
	DataKvr    []Flat `json:"dataKvr,omitempty"`
}
