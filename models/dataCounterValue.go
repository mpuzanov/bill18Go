package models

import (
	"database/sql"
	"log"
)

//CounterValue Показание прибора учёта
type CounterValue struct {
	Occ            int     `json:"occ"`
	CounterID      int     `json:"counter_id"`
	InspectorDate  string  `json:"inspector_date"`
	InspectorValue float64 `json:"inspector_value"`
	ActualValue    float64 `json:"actual_value"`
	FinStr         string  `json:"fin_str"`
	ID             int     `json:"id"`
	SerialNumber   string  `json:"serial_number"`
	ServName       string  `json:"serv_name"`
	FinID          int     `json:"fin_id"`
	Sysuser        string  `json:"sysuser,omitempty"`
}

//GetCounterValueByOcc Показания ПУ по лицевому счёту
func (db *DB) GetCounterValueByOcc(occ int) ([]*CounterValue, error) {
	const querySQL = `
	k_show_counters_value @occ=@occ1, @counter_id=null, @row1=@kolval
`
	kolval := 6
	rows, err := db.Query(querySQL,
		sql.Named("occ1", occ),
		sql.Named("kolval", kolval),
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	slice := []*CounterValue{}
	for rows.Next() {
		zap := CounterValue{}
		err := rows.Scan(&zap.Occ,
			&zap.CounterID,
			&zap.InspectorDate,
			&zap.InspectorValue,
			&zap.ActualValue,
			&zap.FinStr,
			&zap.ID,
			&zap.SerialNumber,
			&zap.ServName,
			&zap.FinID,
			&zap.Sysuser,
		)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
		slice = append(slice, &zap)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
		return nil, err
	}
	return slice, nil
}
