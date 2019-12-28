package models

import (
	"fmt"
)

//DataValue Начисления
type DataValue struct {
	FinStr          string  `json:"fin_str,omitempty" db:"fin_str"`
	Occ             int     `json:"lic,omitempty" db:"lic"`
	Saldo           float64 `json:"saldo,omitempty" db:"saldo"`
	Value           float64 `json:"value,omitempty" db:"value"`
	Added           float64 `json:"added,omitempty" db:"added"`
	Paid            float64 `json:"paid,omitempty" db:"paid"`
	Paymaccount     float64 `json:"paymaccount,omitempty" db:"paymaccount"`
	PaymaccountPeny float64 `json:"paymaccount_peny,omitempty" db:"paymaccount_peny"`
	PaymaccountServ float64 `json:"paymaccount_serv,omitempty" db:"paymaccount_serv"`
	Debt            float64 `json:"debt,omitempty" db:"debt"`
	SupName         string  `json:"sup_name,omitempty" db:"sup_name"`
}

//ToString Строковое представление Начисления
func (zap *DataValue) String() string {
	return fmt.Sprintf("Период: %15s, Лицевой: %9d, Saldo: %9g, Paid: %9g",
		zap.FinStr, zap.Occ, zap.Saldo, zap.Paid)
}
