package models

// Показания приборов учёта
type DataCounterValue struct {
	Occ             int
	Counter_id      int
	Inspector_date  string
	Inspector_value float64
	Actual_value    float64
	Fin_str         string
	Id              int
	Serial_number   string
	Serv_name       string
	Fin_id          int
}

// // GetDataOcc Выдаём информацию по заданному лицевому счёту
// func GetDataOcc(occ int) ([]*DataOcc, error) {
// 	rows, err := db.Query("k_show_occ @occ",
// 		sql.Named("occ", occ))
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer rows.Close()

// 	SliseDataOcc := make([]*DataOcc, 0)
// 	for rows.Next() {
// 		DataOcc1 := new(DataOcc)
// 		err := rows.Scan(&DataOcc1.Occ, &DataOcc1.BasaName, &DataOcc1.Address, &DataOcc1.TipName, &DataOcc1.TotalSq,
// 			&DataOcc1.OccSup, &DataOcc1.FinCurrent, &DataOcc1.FinCurrentStr, &DataOcc1.KolPeople,
// 			&DataOcc1.CV1, &DataOcc1.CV2, &DataOcc1.Rejim)
// 		if err != nil {
// 			return nil, err
// 		}
// 		SliseDataOcc = append(SliseDataOcc, DataOcc1)
// 	}
// 	if err = rows.Err(); err != nil {
// 		return nil, err
// 	}
// 	return SliseDataOcc, nil
// }