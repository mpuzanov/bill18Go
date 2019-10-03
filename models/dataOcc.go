package models

import (
	"fmt"
)

//DataOcc Информация по лицевому счёту
type DataOcc struct {
	Occ           int     `json:"occ,omitempty"`
	BasaName      string  `json:"basa_name,omitempty"`
	Address       string  `json:"address,omitempty"`
	TipName       string  `json:"tip_name,omitempty"`
	TotalSq       float64 `json:"total_sq,omitempty"`
	OccSup        int     `json:"occ_sup,omitempty"`
	FinCurrent    int     `json:"fin_current,omitempty"`
	FinCurrentStr string  `json:"fin_current_str,omitempty"`
	KolPeople     int     `json:"kol_people,omitempty"`
	CV1           int     `json:"CV1,omitempty"`
	CV2           int     `json:"CV2,omitempty"`
	Rejim         string  `json:"rejim,omitempty"`
}

//ToString Строковое представление Начисления
func (zap *DataOcc) ToString() (result string) {
	if zap.FinCurrentStr != "" {
		result = fmt.Sprintf("Период: %15s, Лицевой: %9d, Адрес: %s, Тип фонда: %s",
			zap.FinCurrentStr, zap.Occ, zap.Address, zap.TipName)
	}
	return
}
