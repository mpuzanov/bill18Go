package models

import (
	"database/sql"
	"fmt"
	"log"
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

// GetDataOcc Выдаём информацию по заданному лицевому счёту
func (db *DB) GetDataOcc(occ int) (*DataOcc, error) {
	row := db.QueryRow("k_show_occ @occ=@occ1",
		sql.Named("occ1", occ))

	data := DataOcc{}
	err := row.Scan(&data.Occ,
		&data.BasaName,
		&data.Address,
		&data.TipName,
		&data.TotalSq,
		&data.OccSup,
		&data.FinCurrent,
		&data.FinCurrentStr,
		&data.KolPeople,
		&data.CV1,
		&data.CV2,
		&data.Rejim)
	if err == sql.ErrNoRows {
		return &data, nil
	} else if err != nil {
		log.Fatal(err)
	}
	return &data, nil
}
