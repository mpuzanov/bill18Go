package models

import (
	"fmt"
)

//DataOcc Информация по лицевому счёту
type DataOcc struct {
	Occ           int     `json:"occ,omitempty" db:"occ"`
	BasaName      string  `json:"basa_name,omitempty" db:"basa_name"`
	Address       string  `json:"address,omitempty" db:"adres"`
	TipName       string  `json:"tip_name,omitempty" db:"tip_name"`
	TotalSq       float64 `json:"total_sq,omitempty" db:"total_sq"`
	OccSup        int     `json:"occ_sup,omitempty" db:"occ_sup"`
	FinCurrent    int     `json:"fin_current,omitempty" db:"fin_current"`
	FinCurrentStr string  `json:"fin_current_str,omitempty" db:"fin_current_str"`
	KolPeople     int     `json:"kol_people,omitempty" db:"kol_people"`
	CV1           int     `json:"CV1,omitempty" db:"CV1"`
	CV2           int     `json:"CV2,omitempty" db:"CV2"`
	Rejim         string  `json:"rejim,omitempty" db:"Rejim"`
}

//ToString Строковое представление
func (zap *DataOcc) String() string {
	result := ""
	if zap.FinCurrentStr != "" {
		result = fmt.Sprintf("Период: %15s, Лицевой: %9d, Адрес: %s, Тип фонда: %s",
			zap.FinCurrentStr, zap.Occ, zap.Address, zap.TipName)
	}
	return result
}
