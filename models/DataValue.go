package models

import (
	"database/sql"
	"fmt"
	"log"
)

//DataValue Начисления
type DataValue struct {
	FinStr          string  `json:"fin_str,omitempty"`
	Occ             int     `json:"lic,omitempty"`
	Saldo           float64 `json:"saldo,omitempty"`
	Value           float64 `json:"value,omitempty"`
	Added           float64 `json:"added,omitempty"`
	Paid            float64 `json:"paid,omitempty"`
	Paymaccount     float64 `json:"paymaccount,omitempty"`
	PaymaccountPeny float64 `json:"paymaccount_peny,omitempty"`
	PaymaccountServ float64 `json:"paymaccount_serv,omitempty"`
	Debt            float64 `json:"debt,omitempty"`
	SupName         string  `json:"sup_name,omitempty"`
}

//ToString Строковое представление Начисления
func (zap *DataValue) ToString() string {
	return fmt.Sprintf("Период: %15s, Лицевой: %9d, Saldo: %9g, Paid: %9g",
		zap.FinStr, zap.Occ, zap.Saldo, zap.Paid)
}

//GetDataValueByOcc Начисления по лицевому счёту
func (db *DB) GetDataValueByOcc(occ int) ([]*DataValue, error) {
	const querySQL = `
	k_show_values_occ @occ=@occ1, @row1=12
	`
	rows, err := db.Query(querySQL,
		sql.Named("occ1", occ),
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	slice := []*DataValue{}
	for rows.Next() {
		zap := DataValue{}
		err := rows.Scan(&zap.FinStr,
			&zap.Occ,
			&zap.Saldo,
			&zap.Value,
			&zap.Added,
			&zap.Paid,
			&zap.Paymaccount,
			&zap.PaymaccountPeny,
			&zap.PaymaccountServ,
			&zap.Debt,
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
