package models

// Начисления
type DataValue struct {
	Fin_str          string
	Lic              int
	Saldo            float64
	Value            float64
	Added            float64
	Paid             float64
	Paymaccount      float64
	Paymaccount_peny float64
	Paymaccount_serv float64
	Debt             float64
	sup_name         string
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