package models

//Street Улица
type Street struct {
	StreetName string `json:"name"`
}

//Streets Улицы
type Streets struct {
	DataStreets []Street `json:"dataStreets"`
}

// AllStreets Возвращаем список улиц
func AllStreets() (*Streets, error) {
	rows, err := db.Query("k_show_streets")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	streets := new(Streets)
	for rows.Next() {
		var street Street
		err := rows.Scan(&street.StreetName)
		if err != nil {
			return nil, err
		}
		streets.DataStreets = append(streets.DataStreets, street)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return streets, nil
}
