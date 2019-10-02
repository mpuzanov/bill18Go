package util

import (
	"bytes"
	"encoding/json"
)

//CheckErr функция обработки ошибок
func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}

//Prettyprint Делаем красивый json с отступами
func Prettyprint(b []byte) ([]byte, error) {
	var out bytes.Buffer
	err := json.Indent(&out, b, "", "    ")
	return out.Bytes(), err
}
