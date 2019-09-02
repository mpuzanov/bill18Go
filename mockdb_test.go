package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/mpuzanov/bill18go/models"
)

type mockDB struct{}

func (mdb *mockDB) GetAllStreets() ([]*models.Street, error) {
	tbl := make([]*models.Street, 0)
	tbl = append(tbl, &models.Street{"Молодёжная ул."})
	tbl = append(tbl, &models.Street{"Камбарская ул."})
	return tbl, nil
}
func (mdb *mockDB) GetBuilds(streetName string) (*models.Builds, error) {
	tbl := &models.Builds{"1-я Донская ул.",
		[]models.Build{models.Build{models.Street{"1-я Донская ул."}, "6", "       6"}}}
	return tbl, nil
}
func (mdb *mockDB) GetFlats(streetName, nomDom string) (*models.Flats, error) {
	tbl := &models.Flats{"Авангардная ул.",
		"3",
		[]models.Flat{
			{models.Build{models.Street{"Авангардная ул."}, "3", "       3"},
				"1",
				"       1"},
		},
	}
	return tbl, nil
}
func (mdb *mockDB) GetKvrLic(streetName, nomDom, nomKvr string) (*models.Lics, error) {
	tbl := &models.Lics{"Авангардная ул.",
		"3",
		"1",
		[]models.Lic{
			{345740},
			{345741},
		}}
	return tbl, nil
}
func (mdb *mockDB) GetCounterByOcc(occ int) ([]*models.DataCounter, error) {
	tbl := make([]*models.DataCounter, 0)
	tbl = append(tbl, &models.DataCounter{
		Occ: 45321, CounterID: 45802, ServName: "ГВС", SerialNumber: "1 г", CounterType: "СВ-15Г",
		MaxValue: 99999, UnitID: "кубм", DateCreate: "01.10.2011", PeriodCheck: "01.01.2050",
		ValueDate: "20.07.2019", LastValue: 239, ActualValue: 1, AvgMonth: 2.57,
		Tarif: 19.97, NormaSingle: 3.22, AvgItog: 2.57, KolNorma: 3.22})
	return tbl, nil
}
func (mdb *mockDB) GetCounterValueByOcc(occ int) ([]*models.CounterValue, error) {
	tbl := make([]*models.CounterValue, 0)
	tbl = append(tbl, &models.CounterValue{
		Occ: 45321, CounterID: 45802, InspectorDate: "20.07.2019", InspectorValue: 239, ActualValue: 1,
		FinStr: "июль 2019", ID: 4221868, SerialNumber: "1 г", ServName: "ГВС", FinID: 210, Sysuser: "",
	})
	return tbl, nil
}
func (mdb *mockDB) GetDataValueByOcc(occ int) ([]*models.DataValue, error) {
	tbl := make([]*models.DataValue, 0)
	tbl = append(tbl, &models.DataValue{FinStr: "июль 2019", Occ: 45321, Saldo: 1385.74, Value: 1333.06, Paid: 1333.06, Paymaccount: 1385.74, PaymaccountServ: 1385.74, Debt: 1333.06})
	return tbl, nil
}
func (mdb *mockDB) GetDataPaymByOcc(occ int) ([]*models.DataPaym, error) {
	tbl := make([]*models.DataPaym, 0)
	tbl = append(tbl, &models.DataPaym{FinStr: "июль 2019", Occ: 45321, Date: "17.07.2019", Summa: 1385.74})
	return tbl, nil
}
func (mdb *mockDB) GetDataOcc(occ int) (*models.DataOcc, error) {
	tbl := &models.DataOcc{45321, "komp", "Ижевск, ул. Баранова д.69 кв.1", "ТСЖ Исток", 31.6, 777045321, 210,
		"июль 2019", 1, 5, 24, "норм"}
	return tbl, nil
}

func TestBuildIndex(t *testing.T) {
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/buildIndex", nil)

	env := Env{db: &mockDB{}}
	http.HandlerFunc(env.buildIndex).ServeHTTP(rec, req)

	expected := `{"street_name":"1-я Донская ул.","dataBuilds":[{"nom_dom":"6","nom_dom_sort":"       6"}]}`

	if expected != rec.Body.String() {
		t.Errorf("\n...expected = %v\n...obtained = %v", expected, rec.Body.String())
	}
}

func TestFlatsIndex(t *testing.T) {
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/flatsIndex", nil)

	env := Env{db: &mockDB{}}
	http.HandlerFunc(env.flatsIndex).ServeHTTP(rec, req)

	expected := `{"street_name":"Авангардная ул.","nom_dom":"3","dataKvr":[{"nom_kvr":"1","nom_kvr_sort":"       1"}]}`

	if expected != rec.Body.String() {
		t.Errorf("\n...expected = %v\n...obtained = %v", expected, rec.Body.String())
	}
}

func TestLicsIndex(t *testing.T) {
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/licsIndex", nil)

	env := Env{db: &mockDB{}}
	http.HandlerFunc(env.licsIndex).ServeHTTP(rec, req)

	expected := `{"street_name":"Авангардная ул.","nom_dom":"3","nom_kvr":"1","dataKvrLic":[{"occ":345740},{"occ":345741}]}`

	if expected != rec.Body.String() {
		t.Errorf("\n...expected = %v\n...obtained = %v", expected, rec.Body.String())
	}
}

func TestInfoLicIndex(t *testing.T) {
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/infoLicIndex", nil)

	env := Env{db: &mockDB{}}
	http.HandlerFunc(env.infoLicIndex).ServeHTTP(rec, req)

	expected := `{"occ":45321,"basa_name":"komp","address":"Ижевск, ул. Баранова д.69 кв.1","tip_name":"ТСЖ Исток","total_sq":31.6,"occ_sup":777045321,"fin_current":210,"fin_current_str":"июль 2019","kol_people":1,"CV1":5,"CV2":24,"rejim":"норм"}`

	if expected != rec.Body.String() {
		t.Errorf("\n...expected = %v\n...obtained = %v", expected, rec.Body.String())
	}
}

func TestInfoDataCounter(t *testing.T) {
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/infoDataCounter", nil)

	env := Env{db: &mockDB{}}
	http.HandlerFunc(env.infoDataCounter).ServeHTTP(rec, req)

	expected := `[{"lic":45321,"counter_id":45802,"serv_name":"ГВС","serial_number":"1 г","counter_type":"СВ-15Г","max_value":99999,"unit_id":"кубм","date_create":"01.10.2011","periodCheck":"01.01.2050","value_date":"20.07.2019","last_value":239,"actual_value":1,"avg_month":2.57,"tarif":19.97,"normaSingle":3.22,"avg_itog":2.57,"kol_norma":3.22}]`
	if expected != rec.Body.String() {
		t.Errorf("\n...expected = %v\n...obtained = %v", expected, rec.Body.String())
	}
}

func TestInfoDataCounterValue(t *testing.T) {
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/infoDataCounterValue", nil)

	env := Env{db: &mockDB{}}
	http.HandlerFunc(env.infoDataCounterValue).ServeHTTP(rec, req)

	expected := `[{"occ":45321,"counter_id":45802,"inspector_date":"20.07.2019","inspector_value":239,"actual_value":1,"fin_str":"июль 2019","id":4221868,"serial_number":"1 г","serv_name":"ГВС","fin_id":210}]`
	if expected != rec.Body.String() {
		t.Errorf("\n...expected = %v\n...obtained = %v", expected, rec.Body.String())
	}
}

func TestInfoDataValue(t *testing.T) {
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/infoDataValue", nil)

	env := Env{db: &mockDB{}}
	http.HandlerFunc(env.infoDataValue).ServeHTTP(rec, req)

	expected := `[{"fin_str":"июль 2019","lic":45321,"saldo":1385.74,"value":1333.06,"paid":1333.06,"paymaccount":1385.74,"paymaccount_serv":1385.74,"debt":1333.06}]`
	if expected != rec.Body.String() {
		t.Errorf("\n...expected = %v\n...obtained = %v", expected, rec.Body.String())
	}
}

func TestInfoDataPaym(t *testing.T) {
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/infoDataPaym", nil)

	env := Env{db: &mockDB{}}
	http.HandlerFunc(env.infoDataPaym).ServeHTTP(rec, req)

	expected := `[{"fin_str":"июль 2019","lic":45321,"date":"17.07.2019","summa":1385.74}]`
	if expected != rec.Body.String() {
		t.Errorf("\n...expected = %v\n...obtained = %v", expected, rec.Body.String())
	}
}

func TestStreetIndex(t *testing.T) {
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/streets", nil)

	env := Env{db: &mockDB{}}
	http.HandlerFunc(env.streetIndex).ServeHTTP(rec, req)

	expected := `[{"name":"Молодёжная ул."},{"name":"Камбарская ул."}]`

	if expected != rec.Body.String() {
		t.Errorf("\n...expected = %v\n...obtained = %v", expected, rec.Body.String())
	}
}
