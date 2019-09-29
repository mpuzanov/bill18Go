package util

//checkErr функция обработки ошибок
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
