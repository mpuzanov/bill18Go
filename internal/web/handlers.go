package web

import (
	"encoding/json"
	"html/template"
	"net/http"
	"strconv"

	"github.com/mpuzanov/bill18Go/internal/util"
	"go.uber.org/zap"

	"github.com/gorilla/mux"
)

// //Env Интерфейс для вызова функций sql
// type Env struct {
// 	Db     Datastore
// 	Cfg    *config.Config
// 	logger *zap.Logger
// }

var (
	//env *Env
	t map[string]*template.Template
)

// CreateTemplate формирование шаблонов
func CreateTemplate() {
	t = make(map[string]*template.Template)
	temp := template.Must(template.ParseFiles("public/templates/base.html", "public/templates/header.html", "public/templates/index.html"))
	t["index.html"] = temp
	temp = template.Must(template.ParseFiles("public/templates/base.html", "public/templates/header.html", "public/templates/testapi.html"))
	t["testapi.html"] = temp
	temp = template.Must(template.ParseFiles("public/templates/base.html", "public/templates/header.html", "public/templates/upload.html"))
	t["upload.html"] = temp
}

func (s *myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *myHandler) logRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//logger:=s.logger.With(zap.String("remote_addr", r.RemoteAddr))
		s.logger.Info("Request",
			zap.String("Method", r.Method),
			zap.String("URI", r.RequestURI),
			zap.String("remote_addr", r.RemoteAddr),
		)
		next.ServeHTTP(w, r)
	})
}

func (s *myHandler) configRouter() {

	CreateTemplate()

	s.router.Use(s.logRequest)
	// s.router.HandleFunc("/", s.homePage)
	// s.router.HandleFunc("/hello", s.helloPage)
	// s.router.HandleFunc("/hello/{name}", s.helloPage)

	s.router.PathPrefix("/bill18/static/").Handler(http.StripPrefix("/bill18/static/", http.FileServer(http.Dir("./public/static/"))))

	// для отдачи сервером статичных файлов из папки public/static
	//fs := http.FileServer(http.Dir("./public/static/"))
	//http.Handle("/static/", http.StripPrefix("/static/", fs))

	//s.router.HandleFunc("/", s.HomePage)
	router2 := s.router.PathPrefix("/bill18").Subrouter()
	router2.HandleFunc("/", s.HomePage)
	router2.HandleFunc("/upload", s.upload)
	router2.HandleFunc("/testapi", s.Testapi)

	// router.HandleFunc("/api/streets", BasicAuth(env.StreetIndex))
	// router.HandleFunc("/api/builds/{street_name}", BasicAuth(env.BuildIndex))
	apirouter := router2.PathPrefix("/api/v1").Subrouter()
	//apirouter.Use(BasicAuth)
	apirouter.HandleFunc("/streets", BasicAuth(s.StreetIndex))
	apirouter.HandleFunc("/builds/{street_name}", BasicAuth(s.BuildIndex))
	apirouter.HandleFunc("/flats/{street_name}/{nom_dom:[0-9]+}", BasicAuth(s.FlatsIndex))
	apirouter.HandleFunc("/lics/{street_name}/{nom_dom:[0-9]+}/{nom_kvr:[0-9]+}", BasicAuth(s.LicsIndex))
	apirouter.HandleFunc("/infoLic/{occ:[0-9]+}", BasicAuth(s.InfoLicIndex))
	apirouter.HandleFunc("/infoDataCounter/{occ:[0-9]+}", BasicAuth(s.InfoDataCounter))
	apirouter.HandleFunc("/infoDataCounterValue/{occ:[0-9]+}", BasicAuth(s.InfoDataCounterValue))
	apirouter.HandleFunc("/infoDataValue/{occ:[0-9]+}", BasicAuth(s.InfoDataValue))
	apirouter.HandleFunc("/infoDataPaym/{occ:[0-9]+}", BasicAuth(s.InfoDataPaym))
	apirouter.HandleFunc("/puAddValue/{puId:[0-9]+}/{value:[0-9]+}", BasicAuth(s.PuAddValue))
	apirouter.HandleFunc("/puDelValue/{puId:[0-9]+}/{id:[0-9]+}", BasicAuth(s.PuDelValue))
	apirouter.HandleFunc("/infoDataCounterValueTip/{tip:[0-9]+}", BasicAuth(s.InfoDataCounterValueTip))

}

//getJSONResponse Возвращаем информацию в JSON формате
func (s *myHandler) getJSONResponse(w http.ResponseWriter, r *http.Request, data interface{}) {
	jsData, err := json.Marshal(data)
	if err != nil {
		http.Error(w, "Error creating JSON response", http.StatusInternalServerError)
		s.logger.Error(err.Error())
		return
	}
	if s.cfg.IsPrettyJSON {
		jsData, err = util.Prettyprint(jsData)
		if err != nil {
			s.logger.Error(err.Error())
		}
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(jsData)
	if err != nil {
		s.logger.Error("Fail Write jsData", zap.Error(err))
		return
	}
}

//HomePage ...
func (s *myHandler) HomePage(w http.ResponseWriter, r *http.Request) {
	err := t["index.html"].ExecuteTemplate(w, "base", &struct{ Listen string }{s.cfg.HTTPAddr})
	if err != nil {
		s.logger.Error("Fail ExecuteTemplate index", zap.Error(err))
		return
	}
}

//Testapi ...
func (s *myHandler) Testapi(w http.ResponseWriter, r *http.Request) {
	err := t["testapi.html"].ExecuteTemplate(w, "base", &struct{ Listen string }{s.cfg.HTTPAddr})
	if err != nil {
		s.logger.Error("Fail ExecuteTemplate testapi", zap.Error(err))
		return
	}
}

//StreetIndex ...
func (s *myHandler) StreetIndex(w http.ResponseWriter, r *http.Request) {
	data, err := s.Db.GetAllStreets()
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}
	s.logger.Debug("streetIndex")
	s.getJSONResponse(w, r, data)
}

//BuildIndex ...
func (s *myHandler) BuildIndex(w http.ResponseWriter, r *http.Request) {
	//r.ParseForm()
	//streetName := r.FormValue("street_name")

	vars := mux.Vars(r)
	streetName := vars["street_name"]

	data, err := s.Db.GetBuilds(streetName)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}
	s.logger.Debug("buildIndex", zap.String("streetName", streetName))
	s.getJSONResponse(w, r, data)
}

// FlatsIndex ...
func (s *myHandler) FlatsIndex(w http.ResponseWriter, r *http.Request) {
	//r.ParseForm()
	//log.Tracef("%v\n", r.Form)
	//streetName := r.FormValue("street_name")
	//nomDom := r.FormValue("nom_dom")

	vars := mux.Vars(r)
	streetName := vars["street_name"]
	nomDom := vars["nom_dom"]

	data, err := s.Db.GetFlats(streetName, nomDom)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	s.logger.Debug("flatsIndex", zap.String("streetName", streetName), zap.String("nomDom", nomDom))
	s.getJSONResponse(w, r, data)
}

//LicsIndex ...
func (s *myHandler) LicsIndex(w http.ResponseWriter, r *http.Request) {
	// r.ParseForm()
	// streetName := r.FormValue("street_name")
	// nomDom := r.FormValue("nom_dom")
	// nomKvr := r.FormValue("nom_kvr")

	vars := mux.Vars(r)
	streetName := vars["street_name"]
	nomDom := vars["nom_dom"]
	nomKvr := vars["nom_kvr"]

	data, err := s.Db.GetKvrLic(streetName, nomDom, nomKvr)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}
	s.logger.Debug("licsIndex", zap.String("streetName", streetName), zap.String("nomDom", nomDom), zap.String("nomKvr", nomKvr))
	s.getJSONResponse(w, r, data)
}

//InfoLicIndex ...
func (s *myHandler) InfoLicIndex(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	parOcc := vars["occ"]
	s.logger.Debug("infoLicIndex", zap.String("parOcc", parOcc))
	var occ int
	occ, _ = strconv.Atoi(parOcc) // если неудача пусть будет 0

	data, err := s.Db.GetDataOcc(occ)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}
	s.getJSONResponse(w, r, data)
}

// InfoDataCounter ...
func (s *myHandler) InfoDataCounter(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	parOcc := vars["occ"]
	var occ int
	occ, _ = strconv.Atoi(parOcc) // если неудача пусть будет 0

	data, err := s.Db.GetCounterByOcc(occ)
	if err != nil {
		s.logger.Error(err.Error())
		http.Error(w, http.StatusText(500), 500)
		return
	}
	s.logger.Debug("infoDataCounter", zap.String("parOcc", parOcc))
	s.getJSONResponse(w, r, data)
}

//InfoDataCounterValue ...
func (s *myHandler) InfoDataCounterValue(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	parOcc := vars["occ"]
	var occ int
	occ, _ = strconv.Atoi(parOcc) // если неудача пусть будет 0

	data, err := s.Db.GetCounterValueByOcc(occ)
	if err != nil {
		s.logger.Error(err.Error())
		http.Error(w, http.StatusText(500), 500)
		return
	}
	s.logger.Debug("infoDataCounterValue", zap.String("parOcc", parOcc))
	s.getJSONResponse(w, r, data)
}

//InfoDataValue ...
func (s *myHandler) InfoDataValue(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	parOcc := vars["occ"]
	var occ int
	occ, _ = strconv.Atoi(parOcc) // если неудача пусть будет 0

	data, err := s.Db.GetDataValueByOcc(occ)
	if err != nil {
		s.logger.Error(err.Error())
		http.Error(w, http.StatusText(500), 500)
		return
	}
	s.logger.Debug("infoDataValue", zap.String("parOcc", parOcc))
	s.getJSONResponse(w, r, data)
}

//InfoDataPaym ...
func (s *myHandler) InfoDataPaym(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	parOcc := vars["occ"]
	var occ int
	occ, _ = strconv.Atoi(parOcc) // если неудача пусть будет 0

	data, err := s.Db.GetDataPaymByOcc(occ)
	if err != nil {
		s.logger.Error(err.Error())
		http.Error(w, http.StatusText(500), 500)
		return
	}
	s.logger.Debug("infoDataPaym", zap.String("parOcc", parOcc))
	s.getJSONResponse(w, r, data)
}

//PuAddValue ...
func (s *myHandler) PuAddValue(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	puIDStr := vars["puId"]
	s.logger.Debug("PuDelValue", zap.String("puId", puIDStr))
	puID, err := strconv.Atoi(puIDStr)
	if err != nil {
		s.logger.Error(err.Error())
	}
	valueStr := vars["value"]
	value, err := strconv.Atoi(valueStr)
	if err != nil {
		s.logger.Error(err.Error())
	}
	//valueDateStr := vars["valueDate"]

	data, err := s.Db.PuAddValue(puID, value)
	if err != nil {
		s.logger.Error(err.Error())
		http.Error(w, http.StatusText(500), 500)
		return
	}
	s.getJSONResponse(w, r, data)
}

//PuDelValue ...
func (s *myHandler) PuDelValue(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	puIDStr := vars["puId"]
	s.logger.Debug("PuDelValue", zap.String("puIDStr", puIDStr))
	puID, err := strconv.Atoi(puIDStr)
	if err != nil {
		s.logger.Error("PuDelValue", zap.String("error", err.Error()))
	}
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		s.logger.Error(err.Error())
	}
	//valueDateStr := vars["valueDate"]

	data, err := s.Db.PuDelValue(puID, id)
	if err != nil {
		s.logger.Error(err.Error())
		http.Error(w, http.StatusText(500), 500)
		return
	}
	s.getJSONResponse(w, r, data)
}

//InfoDataCounterValueTip ...
func (s *myHandler) InfoDataCounterValueTip(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	parTip := vars["tip"]
	s.logger.Debug("InfoDataCounterValueTip", zap.String("parTip", parTip))
	var tip int
	tip, _ = strconv.Atoi(parTip) // если неудача пусть будет 0

	data, err := s.Db.GetCounterValueByTip(tip)
	if err != nil {
		s.logger.Error(err.Error())
		http.Error(w, http.StatusText(500), 500)
		return
	}

	s.getJSONResponse(w, r, data)
}
