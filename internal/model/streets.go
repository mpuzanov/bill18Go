package model

import (
	"fmt"
)

//Street Улица
type Street struct {
	StreetName string `json:"name" db:"name"`
}

//String Строковое представление
func (zap *Street) String() string {
	return fmt.Sprintf("%s", zap.StreetName)
}
