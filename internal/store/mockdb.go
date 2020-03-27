package store

import (
	"github.com/mpuzanov/bill18Go/internal/model"
)

//MockDB ...
type MockDB struct{}

//GetAllStreets ...
func (mdb *MockDB) GetAllStreets() (*[]model.Street, error) {
	tbl := make([]model.Street, 0)
	tbl = append(tbl, model.Street{StreetName: "Молодёжная ул."})
	tbl = append(tbl, model.Street{StreetName: "Камбарская ул."})
	return &tbl, nil
}

//GetBuilds ...
func (mdb *MockDB) GetBuilds(streetName string) (*model.Builds, error) {
	tbl := &model.Builds{StreetName: "1-я Донская ул.",
		DataBuilds: []model.Build{model.Build{Street: model.Street{StreetName: "1-я Донская ул."}, NomDom: "6", NomDomSort: "       6"}}}
	return tbl, nil
}

//GetFlats ...
func (mdb *MockDB) GetFlats(streetName, nomDom string) (*model.Flats, error) {
	tbl := &model.Flats{
		StreetName: "Авангардная ул.",
		NomDom:     "3",
		DataKvr: []model.Flat{
			{Build: model.Build{Street: model.Street{StreetName: "Авангардная ул."}, NomDom: "3", NomDomSort: "       3"},
				NomKvr:     "1",
				NomKvrSort: "       1"},
		}}
	return tbl, nil
}

//GetKvrLic ...
func (mdb *MockDB) GetKvrLic(streetName, nomDom, nomKvr string) (*model.Lics, error) {
	tbl := &model.Lics{StreetName: "Авангардная ул.",
		NomDom: "3",
		NomKvr: "1",
		DataKvrLic: []model.Lic{
			{Occ: 345740},
			{Occ: 345741},
		}}
	return tbl, nil
}

//GetCounterByOcc ...
func (mdb *MockDB) GetCounterByOcc(occ int) (*[]model.DataCounter, error) {
	tbl := make([]model.DataCounter, 0)
	tbl = append(tbl, model.DataCounter{
		Occ: 45321, CounterID: 45802, ServName: "ГВС", SerialNumber: "1 г", CounterType: "СВ-15Г",
		MaxValue: 99999, UnitID: "кубм", DateCreate: "01.10.2011", PeriodCheck: "01.01.2050",
		ValueDate: "20.07.2019", LastValue: 239, ActualValue: 1, AvgMonth: 2.57,
		Tarif: 19.97, NormaSingle: 3.22, AvgItog: 2.57, KolNorma: 3.22})
	return &tbl, nil
}

//GetCounterValueByOcc ...
func (mdb *MockDB) GetCounterValueByOcc(occ int) (*[]model.CounterValue, error) {
	tbl := make([]model.CounterValue, 0)
	tbl = append(tbl, model.CounterValue{
		Occ: 45321, CounterID: 45802, InspectorDate: "20.07.2019", InspectorValue: 239, ActualValue: 1,
		FinStr: "июль 2019", ID: 4221868, SerialNumber: "1 г", ServName: "ГВС", FinID: 210, Sysuser: "",
	})
	return &tbl, nil
}

//GetDataValueByOcc ...
func (mdb *MockDB) GetDataValueByOcc(occ int) (*[]model.DataValue, error) {
	tbl := make([]model.DataValue, 0)
	tbl = append(tbl, model.DataValue{FinStr: "июль 2019", Occ: 45321, Saldo: 1385.74, Value: 1333.06, Paid: 1333.06, Paymaccount: 1385.74, PaymaccountServ: 1385.74, Debt: 1333.06})
	return &tbl, nil
}

//GetDataPaymByOcc ...
func (mdb *MockDB) GetDataPaymByOcc(occ int) (*[]model.DataPaym, error) {
	tbl := make([]model.DataPaym, 0)
	tbl = append(tbl, model.DataPaym{FinStr: "июль 2019", Occ: 45321, Date: "17.07.2019", Summa: 1385.74})
	return &tbl, nil
}

//GetDataOcc ...
func (mdb *MockDB) GetDataOcc(occ int) (*model.DataOcc, error) {
	tbl := &model.DataOcc{Occ: 45321, BasaName: "komp", Address: "Ижевск, ул. Баранова д.69 кв.1", TipName: "ТСЖ Исток", TotalSq: 31.6, OccSup: 777045321, FinCurrent: 210,
		FinCurrentStr: "июль 2019", KolPeople: 1, CV1: 5, CV2: 24, Rejim: "норм"}
	return tbl, nil
}

//PuAddValue ...
func (mdb *MockDB) PuAddValue(puID int, value int) (*model.Result, error) {
	tbl := &model.Result{Res: true, Strerror: ""}
	return tbl, nil
}

//PuDelValue ...
func (mdb *MockDB) PuDelValue(puID int, id int) (*model.Result, error) {
	tbl := &model.Result{Res: true, Strerror: ""}
	return tbl, nil
}

//GetCounterValueByTip ...
func (mdb *MockDB) GetCounterValueByTip(tipID int) (*[]model.CounterValueTip, error) {
	tbl := make([]model.CounterValueTip, 0)
	tbl = append(tbl, model.CounterValueTip{Occ: 45321})
	return &tbl, nil
}
