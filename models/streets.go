package models

import (
	"fmt"
)

//Street Улица
type Street struct {
	StreetName string `json:"name"`
}

//ToString Строковое представление Платежа
func (zap *Street) ToString() string {
	return fmt.Sprintf("%s", zap.StreetName)
}
