package models

import (
	"fmt"
)

//DataPaym Платежи
type DataPaym struct {
	FinStr          string  `json:"fin_str,omitempty" db:"fin_str"`
	Occ             int     `json:"lic,omitempty" db:"lic"`
	Date            string  `json:"date,omitempty" db:"date"`
	Summa           float64 `json:"summa,omitempty" db:"summa"`
	PaymaccountPeny float64 `json:"paymaccount_peny,omitempty" db:"paymaccount_peny"`
	SupName         string  `json:"sup_name,omitempty" db:"sup_name"`
}

//ToString Строковое представление Платежа
func (zap *DataPaym) String() string {
	return fmt.Sprintf("Период: %15s, Лицевой: %9d, Дата: %10s, Сумма: %9g",
		zap.FinStr, zap.Occ, zap.Date, zap.Summa)
}
