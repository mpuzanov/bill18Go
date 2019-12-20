package util

import (
	"bytes"
	"encoding/json"
)

//PanicIfErr функция обработки ошибок
func PanicIfErr(err error) {
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
