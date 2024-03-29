package model

//Build Дом
type Build struct {
	Street     `json:"-"`
	NomDom     string `json:"nom_dom" db:"nom_dom"`
	NomDomSort string `json:"nom_dom_sort" db:"nom_dom_sort"`
}

//Builds Дома
type Builds struct {
	StreetName string  `json:"street_name,omitempty"`
	DataBuilds []Build `json:"dataBuilds,omitempty"`
}
