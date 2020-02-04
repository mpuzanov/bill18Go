package model

import (
	"github.com/mpuzanov/bill18Go/internal/sqltype"
)

//CounterValueTip Показания приборов учёта по типу фонда(список)
type CounterValueTip struct {
	Occ                int                 `json:"occ" db:"occ"`
	StreetName         sqltype.NullString  `json:"street_name" db:"street_name"`
	NomDom             sqltype.NullString  `json:"nom_dom" db:"nom_dom"`
	NomKvr             sqltype.NullString  `json:"nom_kvr" db:"nom_kvr"`
	FinStr             string              `json:"fin_str" db:"fin_str"`
	FinID              int                 `json:"fin_id" db:"fin_id"`
	CounterID          int                 `json:"counter_id" db:"counter_id"`
	SerialNumber       string              `json:"serial_number" db:"serial_number"`
	ServName           string              `json:"serv_name" db:"serv_name"`
	DateCreate         sqltype.NullTime    `json:"date_create" db:"date_create"`
	PeriodCheck        sqltype.NullTime    `json:"period_check" db:"period_check"`
	InspectorDate      sqltype.NullTime    `json:"inspector_date" db:"inspector_date"`
	InspectorValue     sqltype.NullFloat64 `json:"inspector_value" db:"inspector_value"`
	ActualValue        sqltype.NullFloat64 `json:"actual_value" db:"actual_value"`
	PredInspectorDate  sqltype.NullTime    `json:"pred_inspector_date" db:"pred_inspector_date"`
	PredInspectorValue sqltype.NullFloat64 `json:"pred_inspector_value" db:"pred_inspector_value"`
	PredActualValue    sqltype.NullFloat64 `json:"pred_actual_value" db:"pred_actual_value"`
}
