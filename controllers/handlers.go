package controllers

import (
	"encoding/json"
	"html/template"
	"net/http"
	"strconv"

	"github.com/mpuzanov/bill18Go/config"
	log "github.com/mpuzanov/bill18Go/logger"
	"github.com/mpuzanov/bill18Go/util"
)

//Env Интерфейс для вызова функций sql
type Env struct {
	db  Datastore
	cfg *config.Config
}

var (
	env       *Env
	bootstrap *template.Template

	t map[string]*template.Template
)

func init() {

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
	if env.cfg.IsPrettyJSON {
		jsData, err = util.Prettyprint(jsData)
		if err != nil {
			log.Error(err)
		}
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsData)
}

func (env *Env) homePage(w http.ResponseWriter, r *http.Request) {
	t["index.html"].ExecuteTemplate(w, "base", &struct{ Listen string }{env.cfg.Listen})
}

func (env *Env) testapi(w http.ResponseWriter, r *http.Request) {
	t["testapi.html"].ExecuteTemplate(w, "base", &struct{ Listen string }{env.cfg.Listen})
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
	data, err := env.db.GetBuilds(streetName)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}
	log.Traceln("buildIndex", streetName)
	env.getJSONResponse(w, r, data)
}

func (env *Env) flatsIndex(w http.ResponseWriter, r *http.Request) {
	//r.ParseForm()
	//log.Tracef("%v\n", r.Form)

	streetName := r.FormValue("street_name")
	nomDom := r.FormValue("nom_dom")
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
