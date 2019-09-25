package apiserver

import (
	"bytes"
	"encoding/json"
	"html/template"
	"net/http"
	"strconv"

	"github.com/mpuzanov/bill18Go/config"
	log "github.com/mpuzanov/bill18Go/logger"
	"github.com/mpuzanov/bill18Go/models"
)

//Env Интерфейс для вызова функций sql
type Env struct {
	db  models.Datastore
	cfg *config.Config
}

var (
	env *Env
)

//RunHTTP ...
func RunHTTP(cfg *config.Config) *http.Server {
	log.Info("Запуск веб-сервера на", cfg.Listen)

	//=====================================================================
	db, err := models.GetInitDB(cfg.DatabaseURL)
	if err != nil {
		panic(err)
	}
	env = &Env{db: db, cfg: cfg}
	log.Traceln("connString:", cfg.DatabaseURL)
	//=====================================================================

	srv := &http.Server{Addr: cfg.Listen}

	http.HandleFunc("/", env.homePage)
	http.HandleFunc("/streets", env.streetIndex)
	http.HandleFunc("/builds", env.buildIndex)
	http.HandleFunc("/flats", env.flatsIndex)
	http.HandleFunc("/lics", env.licsIndex)
	http.HandleFunc("/infoLic", env.infoLicIndex)
	http.HandleFunc("/infoDataCounter", env.infoDataCounter)

	http.HandleFunc("/infoDataCounterValue", env.infoDataCounterValue)
	http.HandleFunc("/infoDataValue", env.infoDataValue)
	http.HandleFunc("/infoDataPaym", env.infoDataPaym)

	go func() {
		log.Fatal(srv.ListenAndServe())
	}()

	return srv
}

//checkErr функция обработки ошибок
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

//prettyprint Делаем красивый json с отступами
func prettyprint(b []byte) ([]byte, error) {
	var out bytes.Buffer
	err := json.Indent(&out, b, "", "    ")
	return out.Bytes(), err
}

//getJSONResponse Возвращаем информацию в JSON формате
func (env *Env) getJSONResponse(w http.ResponseWriter, r *http.Request, data interface{}) {
	//jsData, err := json.MarshalIndent(data, "", "    ")
	jsData, err := json.Marshal(data)
	if err != nil {
		http.Error(w, "Error creating JSON response", http.StatusInternalServerError)
		log.Error(err)
		return
	}
	if env.cfg.IsPrettyJSON == true {
		jsData, err = prettyprint(jsData)
		if err != nil {
			log.Error(err)
		}
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsData)
}

func (env *Env) homePage(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/index.html", "templates/header.html", "templates/footer.html")
	if err != nil {
		log.Error("template.ParseFiles", err.Error())
		return
	}
	t.ExecuteTemplate(w, "index", &struct{ Listen string }{env.cfg.Listen})
}

func (env *Env) streetIndex(w http.ResponseWriter, r *http.Request) {
	data, err := env.db.GetAllStreets()
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}
	log.Traceln("streetIndex")
	env.getJSONResponse(w, r, data)
}

func (env *Env) buildIndex(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	streetName := r.FormValue("street_name")
	if streetName == "" {
		streetName = "1-я Донская ул."
	}
	data, err := env.db.GetBuilds(streetName)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}
	log.Traceln("buildIndex", streetName)
	env.getJSONResponse(w, r, data)
}

func (env *Env) flatsIndex(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	//log.Tracef("%v\n", r.Form)

	streetName := r.FormValue("street_name")
	nomDom := r.FormValue("nom_dom")
	if streetName == "" {
		streetName = "1-я Донская ул."
	}
	if nomDom == "" {
		nomDom = "6"
	}
	data, err := env.db.GetFlats(streetName, nomDom)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	log.Traceln("flatsIndex", streetName, nomDom)
	env.getJSONResponse(w, r, data)
}

func (env *Env) licsIndex(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	streetName := r.FormValue("street_name")
	nomDom := r.FormValue("nom_dom")
	nomKvr := r.FormValue("nom_kvr")
	if streetName == "" {
		streetName = "1-я Донская ул."
	}
	if nomDom == "" {
		nomDom = "6"
	}
	if nomKvr == "" {
		nomKvr = "2"
	}
	data, err := env.db.GetKvrLic(streetName, nomDom, nomKvr)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}
	log.Traceln("licsIndex", streetName, nomDom, nomKvr)
	env.getJSONResponse(w, r, data)
}

func (env *Env) infoLicIndex(w http.ResponseWriter, r *http.Request) {
	parOcc := r.FormValue("occ")
	if parOcc == "" {
		parOcc = "0"
	}
	occ, _ := strconv.Atoi(parOcc)

	data, err := env.db.GetDataOcc(occ)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}
	log.Traceln("infoLicIndex", parOcc)
	env.getJSONResponse(w, r, data)
}

func (env *Env) infoDataCounter(w http.ResponseWriter, r *http.Request) {
	parOcc := r.FormValue("occ")
	if parOcc == "" {
		parOcc = "0"
	}
	occ, _ := strconv.Atoi(parOcc)

	data, err := env.db.GetCounterByOcc(occ)
	if err != nil {
		log.Errorf("infoDataCounter: %s\n", err)
		http.Error(w, http.StatusText(500), 500)
		return
	}
	log.Traceln("infoDataCounter", parOcc)
	env.getJSONResponse(w, r, data)
}

func (env *Env) infoDataCounterValue(w http.ResponseWriter, r *http.Request) {
	parOcc := r.FormValue("occ")
	if parOcc == "" {
		parOcc = "0"
	}
	occ, _ := strconv.Atoi(parOcc)

	data, err := env.db.GetCounterValueByOcc(occ)
	if err != nil {
		log.Errorf("infoDataCounterValue: %s\n", err)
		http.Error(w, http.StatusText(500), 500)
		return
	}
	log.Traceln("infoDataCounterValue", parOcc)
	env.getJSONResponse(w, r, data)
}

func (env *Env) infoDataValue(w http.ResponseWriter, r *http.Request) {
	parOcc := r.FormValue("occ")
	if parOcc == "" {
		parOcc = "0"
	}
	occ, _ := strconv.Atoi(parOcc)

	data, err := env.db.GetDataValueByOcc(occ)
	if err != nil {
		log.Errorf("infoDataValue: %s\n", err)
		http.Error(w, http.StatusText(500), 500)
		return
	}
	log.Traceln("infoDataValue", parOcc)
	env.getJSONResponse(w, r, data)
}

func (env *Env) infoDataPaym(w http.ResponseWriter, r *http.Request) {
	parOcc := r.FormValue("occ")
	if parOcc == "" {
		parOcc = "0"
	}
	occ, _ := strconv.Atoi(parOcc)
	data, err := env.db.GetDataPaymByOcc(occ)
	if err != nil {
		log.Errorf("infoDataPaym: %s\n", err)
		http.Error(w, http.StatusText(500), 500)
		return
	}
	log.Traceln("infoDataPaym", parOcc)
	env.getJSONResponse(w, r, data)
}
