package models

import (
	"fmt"
)

//DataCounter Прибор учёта
type DataCounter struct {
	Occ          int     `json:"lic,omitempty" db:"lic"`
	CounterID    int     `json:"counter_id,omitempty" db:"counter_id"`
	ServName     string  `json:"serv_name,omitempty" db:"serv_name"`
	SerialNumber string  `json:"serial_number,omitempty" db:"serial_number"`
	CounterType  string  `json:"counter_type,omitempty" db:"type"`
	MaxValue     int     `json:"max_value,omitempty" db:"max_value"`
	UnitID       string  `json:"unit_id,omitempty" db:"unit_id"`
	CountValue   float64 `json:"count_value,omitempty" db:"count_value"`
	DateCreate   string  `json:"date_create,omitempty" db:"date_create"`
	PeriodCheck  string  `json:"periodCheck,omitempty" db:"PeriodCheck"`
	ValueDate    string  `json:"value_date,omitempty" db:"value_date"`
	LastValue    float64 `json:"last_value,omitempty" db:"last_value"`
	ActualValue  float64 `json:"actual_value,omitempty" db:"actual_value"`
	AvgMonth     float64 `json:"avg_month,omitempty" db:"avg_month"`
	Tarif        float64 `json:"tarif,omitempty" db:"tarif"`
	NormaSingle  float64 `json:"normaSingle,omitempty" db:"NormaSingle"`
	AvgItog      float64 `json:"avg_itog,omitempty" db:"avg_itog"`
	KolNorma     float64 `json:"kol_norma,omitempty" db:"kol_norma"`
}

// DataCounters список приборов учёта
type DataCounters struct {
	DataSet []DataCounter
}

//String Строковое представление ПУ
func (zap *DataCounter) String() string {
	return fmt.Sprintf("Лицевой: %d, Код: %6d, Услуга: %-6s № %-15s Значение: %-9g Дата установки: %s",
		zap.Occ, zap.CounterID, zap.ServName, zap.SerialNumber, zap.CountValue, zap.DateCreate)
	//zap.Occ, zap.CounterID, zap.ServName, zap.SerialNumber, zap.CountValue, util.DataFromSQLToFormat(zap.DateCreate, "02.01.2006"))
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
