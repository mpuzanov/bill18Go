package model

//Result структура возврата
type Result struct {
	Res      bool   `json:"res" db:"res"`
	Strerror string `json:"strerror" db:"strerror"`
}
