package models

import (
	"fmt"
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
