package models

//CounterValue Показание прибора учёта
type CounterValue struct {
	Occ            int     `json:"occ" db:"occ"`
	CounterID      int     `json:"counter_id" db:"counter_id"`
	InspectorDate  string  `json:"inspector_date" db:"inspector_date"`
	InspectorValue float64 `json:"inspector_value" db:"inspector_value"`
	ActualValue    float64 `json:"actual_value" db:"actual_value"`
	FinStr         string  `json:"fin_str" db:"fin_str"`
	ID             int     `json:"id" db:"id"`
	SerialNumber   string  `json:"serial_number" db:"serial_number"`
	ServName       string  `json:"serv_name" db:"serv_name"`
	FinID          int     `json:"fin_id" db:"fin_id"`
	Sysuser        string  `json:"sysuser,omitempty" db:"sysuser"`
}
