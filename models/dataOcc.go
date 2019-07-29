package models

import (
	"database/sql"
	"fmt"
)

//DataOcc Информация по лицевому счёту
type DataOcc struct {
	Occ           int     `json:"occ"`
	BasaName      string  `json:"basa_name"`
	Address       string  `json:"address"`
	TipName       string  `json:"tip_name"`
	TotalSq       float64 `json:"total_sq"`
	OccSup        int     `json:"occ_sup"`
	FinCurrent    int     `json:"fin_current"`
	FinCurrentStr string  `json:"fin_current_str"`
	KolPeople     int     `json:"kol_people"`
	CV1           int
	CV2           int
	Rejim         string `json:"Rejim,omitempty"`
}

// GetDataOcc Выдаём информацию по заданному лицевому счёту
func GetDataOcc(occ string) (*DataOcc, error) {
	rows, err := db.Query("k_show_occ @occ",
		sql.Named("occ", occ))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	DataOcc1 := new(DataOcc)
	for rows.Next() {
		err := rows.Scan(&DataOcc1.Occ, &DataOcc1.BasaName, &DataOcc1.Address, &DataOcc1.TipName,
			&DataOcc1.TotalSq, &DataOcc1.OccSup, &DataOcc1.FinCurrent, &DataOcc1.FinCurrentStr,
			&DataOcc1.KolPeople, &DataOcc1.CV1, &DataOcc1.CV2)
		//, &DataOcc1.Rejim)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return DataOcc1, nil
}
