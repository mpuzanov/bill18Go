package models

import (
	"database/sql"
	"fmt"
	"log"
)

//DataPaym Платежи
type DataPaym struct {
	FinStr          string  `json:"fin_str,omitempty"`
	Occ             int     `json:"lic,omitempty"`
	Date            string  `json:"date,omitempty"`
	Summa           float64 `json:"summa,omitempty"`
	PaymaccountPeny float64 `json:"paymaccount_peny,omitempty"`
	SupName         string  `json:"sup_name,omitempty"`
}

//ToString Строковое представление Платежа
func (zap *DataPaym) ToString() string {
	return fmt.Sprintf("Период: %15s, Лицевой: %9d, Дата: %10s, Сумма: %9g",
		zap.FinStr, zap.Occ, zap.Date, zap.Summa)
}

//GetDataPaymByOcc Список платежей по лицевому счёту
func (db *DB) GetDataPaymByOcc(occ int) ([]*DataPaym, error) {
	const querySQL = `
	k_show_payings @occ=@occ1
`
	rows, err := db.Query(querySQL,
		sql.Named("occ1", occ),
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	slice := []*DataPaym{}
	for rows.Next() {
		zap := DataPaym{}
		err := rows.Scan(&zap.FinStr,
			&zap.Occ,
			&zap.Date,
			&zap.Summa,
			&zap.PaymaccountPeny,
			&zap.SupName,
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
