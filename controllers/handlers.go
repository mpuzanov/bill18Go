package controllers

import (
	"encoding/json"
	"html/template"
	"net/http"
	"strconv"

	"github.com/mpuzanov/bill18Go/config"
	log "github.com/mpuzanov/bill18Go/logger"
	"github.com/mpuzanov/bill18Go/util"

	"github.com/gorilla/mux"
)

//Env Интерфейс для вызова функций sql
type Env struct {
	Db  Datastore
	Cfg *config.Config
}

var (
	env *Env
	t   map[string]*template.Template
)

// CreateTemplate формирование шаблонов
func CreateTemplate() {
	t = make(map[string]*template.Template)
	temp := template.Must(template.ParseFiles("templates/base.html", "templates/header.html", "templates/index.html"))
	t["index.html"] = temp
	temp = template.Must(template.ParseFiles("templates/base.html", "templates/header.html", "templates/testapi.html"))
	t["testapi.html"] = temp
	temp = template.Must(template.ParseFiles("templates/base.html", "templates/header.html", "templates/upload.html"))
	t["upload.html"] = temp
}

//getJSONResponse Возвращаем информацию в JSON формате
func (env *Env) getJSONResponse(w http.ResponseWriter, r *http.Request, data interface{}) {
	jsData, err := json.Marshal(data)
	if err != nil {
		http.Error(w, "Error creating JSON response", http.StatusInternalServerError)
		log.Error(err)
		return
	}
	if env.Cfg.IsPrettyJSON {
		jsData, err = util.Prettyprint(jsData)
		if err != nil {
			log.Error(err)
		}
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsData)
}

func (env *Env) HomePage(w http.ResponseWriter, r *http.Request) {
	t["index.html"].ExecuteTemplate(w, "base", &struct{ Listen string }{env.Cfg.Listen})
}

func (env *Env) Testapi(w http.ResponseWriter, r *http.Request) {
	t["testapi.html"].ExecuteTemplate(w, "base", &struct{ Listen string }{env.Cfg.Listen})
}

func (env *Env) StreetIndex(w http.ResponseWriter, r *http.Request) {
	data, err := env.Db.GetAllStreets()
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}
	log.Traceln("streetIndex")
	env.getJSONResponse(w, r, data)
}

func (env *Env) BuildIndex(w http.ResponseWriter, r *http.Request) {
	//r.ParseForm()
	//streetName := r.FormValue("street_name")

	vars := mux.Vars(r)
	streetName := vars["street_name"]

	data, err := env.Db.GetBuilds(streetName)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}
	log.Traceln("buildIndex", streetName)
	env.getJSONResponse(w, r, data)
}

func (env *Env) FlatsIndex(w http.ResponseWriter, r *http.Request) {
	//r.ParseForm()
	//log.Tracef("%v\n", r.Form)
	//streetName := r.FormValue("street_name")
	//nomDom := r.FormValue("nom_dom")

	vars := mux.Vars(r)
	streetName := vars["street_name"]
	nomDom := vars["nom_dom"]

	data, err := env.Db.GetFlats(streetName, nomDom)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	log.Traceln("flatsIndex", streetName, nomDom)
	env.getJSONResponse(w, r, data)
}

func (env *Env) LicsIndex(w http.ResponseWriter, r *http.Request) {
	// r.ParseForm()
	// streetName := r.FormValue("street_name")
	// nomDom := r.FormValue("nom_dom")
	// nomKvr := r.FormValue("nom_kvr")

	vars := mux.Vars(r)
	streetName := vars["street_name"]
	nomDom := vars["nom_dom"]
	nomKvr := vars["nom_kvr"]

	data, err := env.Db.GetKvrLic(streetName, nomDom, nomKvr)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}
	log.Traceln("licsIndex", streetName, nomDom, nomKvr)
	env.getJSONResponse(w, r, data)
}

func (env *Env) InfoLicIndex(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	parOcc := vars["occ"]
	log.Traceln("infoLicIndex", parOcc)
	var occ int
	occ, _ = strconv.Atoi(parOcc) // если неудача пусть будет 0

	data, err := env.Db.GetDataOcc(occ)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	env.getJSONResponse(w, r, data)
}

func (env *Env) InfoDataCounter(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	parOcc := vars["occ"]
	var occ int
	occ, _ = strconv.Atoi(parOcc) // если неудача пусть будет 0

	data, err := env.Db.GetCounterByOcc(occ)
	if err != nil {
		log.Errorf("infoDataCounter: %s\n", err)
		http.Error(w, http.StatusText(500), 500)
		return
	}
	log.Traceln("infoDataCounter", parOcc)
	env.getJSONResponse(w, r, data)
}

func (env *Env) InfoDataCounterValue(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	parOcc := vars["occ"]
	var occ int
	occ, _ = strconv.Atoi(parOcc) // если неудача пусть будет 0

	data, err := env.Db.GetCounterValueByOcc(occ)
	if err != nil {
		log.Errorf("infoDataCounterValue: %s\n", err)
		http.Error(w, http.StatusText(500), 500)
		return
	}
	log.Traceln("infoDataCounterValue", parOcc)
	env.getJSONResponse(w, r, data)
}

func (env *Env) InfoDataValue(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	parOcc := vars["occ"]
	var occ int
	occ, _ = strconv.Atoi(parOcc) // если неудача пусть будет 0

	data, err := env.Db.GetDataValueByOcc(occ)
	if err != nil {
		log.Errorf("infoDataValue: %s\n", err)
		http.Error(w, http.StatusText(500), 500)
		return
	}
	log.Traceln("infoDataValue", parOcc)
	env.getJSONResponse(w, r, data)
}

func (env *Env) InfoDataPaym(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	parOcc := vars["occ"]
	var occ int
	occ, _ = strconv.Atoi(parOcc) // если неудача пусть будет 0

	data, err := env.Db.GetDataPaymByOcc(occ)
	if err != nil {
		log.Errorf("infoDataPaym: %s\n", err)
		http.Error(w, http.StatusText(500), 500)
		return
	}
	log.Traceln("infoDataPaym", parOcc)
	env.getJSONResponse(w, r, data)
}
