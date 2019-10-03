package models

import (
	"fmt"
)

//DataCounter Прибор учёта
type DataCounter struct {
	Occ          int     `json:"lic,omitempty"`
	CounterID    int     `json:"counter_id,omitempty"`
	ServName     string  `json:"serv_name,omitempty"`
	SerialNumber string  `json:"serial_number,omitempty"`
	CounterType  string  `json:"counter_type,omitempty"`
	MaxValue     int     `json:"max_value,omitempty"`
	UnitID       string  `json:"unit_id,omitempty"`
	CountValue   float64 `json:"count_value,omitempty"`
	DateCreate   string  `json:"date_create,omitempty"`
	PeriodCheck  string  `json:"periodCheck,omitempty"`
	ValueDate    string  `json:"value_date,omitempty"`
	LastValue    float64 `json:"last_value,omitempty"`
	ActualValue  float64 `json:"actual_value,omitempty"`
	AvgMonth     float64 `json:"avg_month,omitempty"`
	Tarif        float64 `json:"tarif,omitempty"`
	NormaSingle  float64 `json:"normaSingle,omitempty"`
	AvgItog      float64 `json:"avg_itog,omitempty"`
	KolNorma     float64 `json:"kol_norma,omitempty"`
}

//ToString Строковое представление ПУ
func (zap *DataCounter) ToString() string {
	return fmt.Sprintf("Лицевой: %d, Код: %6d, Услуга: %-6s № %-15s Значение: %-9g Дата установки: %s",
		zap.Occ, zap.CounterID, zap.ServName, zap.SerialNumber, zap.CountValue, zap.DateCreate)
	//zap.Occ, zap.CounterID, zap.ServName, zap.SerialNumber, zap.CountValue, util.DataFromSQLToFormat(zap.DateCreate, "02.01.2006"))
}

//ToStringMoreLines Строковое представление ПУ в несколько строк
func (zap *DataCounter) ToStringMoreLines() string {
	return fmt.Sprintf("Лицевой: %d\nКод: %d\nУслуга: %s\n№ %s\nЗначение: %g\nДата установки: %s",
		zap.Occ, zap.CounterID, zap.ServName, zap.SerialNumber, zap.CountValue, zap.DateCreate)
}
