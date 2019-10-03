package models

import (
	"fmt"
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
