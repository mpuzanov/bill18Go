package models

//CounterValue Показание прибора учёта
type CounterValue struct {
	Occ            int     `json:"occ"`
	CounterID      int     `json:"counter_id"`
	InspectorDate  string  `json:"inspector_date"`
	InspectorValue float64 `json:"inspector_value"`
	ActualValue    float64 `json:"actual_value"`
	FinStr         string  `json:"fin_str"`
	ID             int     `json:"id"`
	SerialNumber   string  `json:"serial_number"`
	ServName       string  `json:"serv_name"`
	FinID          int     `json:"fin_id"`
	Sysuser        string  `json:"sysuser,omitempty"`
}
