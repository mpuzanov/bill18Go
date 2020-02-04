package controller_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/mpuzanov/bill18Go/internal/config"
	contr "github.com/mpuzanov/bill18Go/internal/controller"
)

func TestBuildIndex(t *testing.T) {
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/buildIndex", nil)

	env := contr.Env{Db: &contr.MockDB{}, Cfg: &config.Config{}}
	http.HandlerFunc(env.BuildIndex).ServeHTTP(rec, req)

	expected := `{"street_name":"1-я Донская ул.","dataBuilds":[{"nom_dom":"6","nom_dom_sort":"       6"}]}`

	if expected != rec.Body.String() {
		t.Errorf("\n...expected = %v\n...obtained = %v", expected, rec.Body.String())
	}
}

func TestFlatsIndex(t *testing.T) {
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/flatsIndex", nil)

	env := contr.Env{Db: &contr.MockDB{}, Cfg: &config.Config{}}
	http.HandlerFunc(env.FlatsIndex).ServeHTTP(rec, req)

	expected := `{"street_name":"Авангардная ул.","nom_dom":"3","dataKvr":[{"nom_kvr":"1","nom_kvr_sort":"       1"}]}`

	if expected != rec.Body.String() {
		t.Errorf("\n...expected = %v\n...obtained = %v", expected, rec.Body.String())
	}
}

func TestLicsIndex(t *testing.T) {
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/licsIndex", nil)

	env := contr.Env{Db: &contr.MockDB{}, Cfg: &config.Config{}}
	http.HandlerFunc(env.LicsIndex).ServeHTTP(rec, req)

	expected := `{"street_name":"Авангардная ул.","nom_dom":"3","nom_kvr":"1","dataKvrLic":[{"occ":345740},{"occ":345741}]}`

	if expected != rec.Body.String() {
		t.Errorf("\n...expected = %v\n...obtained = %v", expected, rec.Body.String())
	}
}

func TestInfoLicIndex(t *testing.T) {
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/infoLicIndex", nil)

	env := contr.Env{Db: &contr.MockDB{}, Cfg: &config.Config{}}
	http.HandlerFunc(env.InfoLicIndex).ServeHTTP(rec, req)

	expected := `{"occ":45321,"basa_name":"komp","address":"Ижевск, ул. Баранова д.69 кв.1","tip_name":"ТСЖ Исток","total_sq":31.6,"occ_sup":777045321,"fin_current":210,"fin_current_str":"июль 2019","kol_people":1,"CV1":5,"CV2":24,"rejim":"норм"}`

	if expected != rec.Body.String() {
		t.Errorf("\n...expected = %v\n...obtained = %v", expected, rec.Body.String())
	}
}

func TestInfoDataCounter(t *testing.T) {
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/infoDataCounter", nil)

	env := contr.Env{Db: &contr.MockDB{}, Cfg: &config.Config{}}
	http.HandlerFunc(env.InfoDataCounter).ServeHTTP(rec, req)

	expected := `[{"lic":45321,"counter_id":45802,"serv_name":"ГВС","serial_number":"1 г","counter_type":"СВ-15Г","max_value":99999,"unit_id":"кубм","date_create":"01.10.2011","periodCheck":"01.01.2050","value_date":"20.07.2019","last_value":239,"actual_value":1,"avg_month":2.57,"tarif":19.97,"normaSingle":3.22,"avg_itog":2.57,"kol_norma":3.22}]`
	if expected != rec.Body.String() {
		t.Errorf("\n...expected = %v\n...obtained = %v", expected, rec.Body.String())
	}
}

func TestInfoDataCounterValue(t *testing.T) {
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/infoDataCounterValue", nil)

	env := contr.Env{Db: &contr.MockDB{}, Cfg: &config.Config{}}
	http.HandlerFunc(env.InfoDataCounterValue).ServeHTTP(rec, req)

	expected := `[{"occ":45321,"counter_id":45802,"inspector_date":"20.07.2019","inspector_value":239,"actual_value":1,"fin_str":"июль 2019","id":4221868,"serial_number":"1 г","serv_name":"ГВС","fin_id":210}]`
	if expected != rec.Body.String() {
		t.Errorf("\n...expected = %v\n...obtained = %v", expected, rec.Body.String())
	}
}

func TestInfoDataValue(t *testing.T) {
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/infoDataValue", nil)

	env := contr.Env{Db: &contr.MockDB{}, Cfg: &config.Config{}}
	http.HandlerFunc(env.InfoDataValue).ServeHTTP(rec, req)

	expected := `[{"fin_str":"июль 2019","lic":45321,"saldo":1385.74,"value":1333.06,"paid":1333.06,"paymaccount":1385.74,"paymaccount_serv":1385.74,"debt":1333.06}]`
	if expected != rec.Body.String() {
		t.Errorf("\n...expected = %v\n...obtained = %v", expected, rec.Body.String())
	}
}

func TestInfoDataPaym(t *testing.T) {
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/infoDataPaym", nil)

	env := contr.Env{Db: &contr.MockDB{}, Cfg: &config.Config{}}
	http.HandlerFunc(env.InfoDataPaym).ServeHTTP(rec, req)

	expected := `[{"fin_str":"июль 2019","lic":45321,"date":"17.07.2019","summa":1385.74}]`
	if expected != rec.Body.String() {
		t.Errorf("\n...expected = %v\n...obtained = %v", expected, rec.Body.String())
	}
}

func TestStreetIndex(t *testing.T) {
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/streets", nil)

	env := contr.Env{Db: &contr.MockDB{}, Cfg: &config.Config{}}
	http.HandlerFunc(env.StreetIndex).ServeHTTP(rec, req)

	expected := `[{"name":"Молодёжная ул."},{"name":"Камбарская ул."}]`

	if expected != rec.Body.String() {
		t.Errorf("\n...expected = %v\n...obtained = %v", expected, rec.Body.String())
	}
}
