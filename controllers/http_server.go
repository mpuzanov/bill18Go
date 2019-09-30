package controllers

import (
	"bytes"
	"context"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"

	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/mpuzanov/bill18Go/config"
	log "github.com/mpuzanov/bill18Go/logger"
	"github.com/mpuzanov/bill18Go/models"
)

//Env Интерфейс для вызова функций sql
type Env struct {
	db  models.Datastore
	cfg *config.Config
}

//Bill18Server my http-server
type Bill18Server struct {
	server *http.Server
	config *config.Config
}

var (
	env *Env
)

//NewServer ...
func NewServer(cfg *config.Config) *Bill18Server {
	log.Info("Запуск веб-сервера на ", cfg.Listen)

	//=====================================================================
	db, err := models.GetInitDB(cfg.DatabaseURL)
	if err != nil {
		panic(err)
	}
	env = &Env{db: db, cfg: cfg}
	log.Traceln("connString:", cfg.DatabaseURL)
	//=====================================================================

	defaultServer := &http.Server{
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		Addr:         cfg.Listen,
	}
	srv := &Bill18Server{
		server: defaultServer,
		config: cfg,
	}
	router := mux.NewRouter()

	// fileServer := http.FileServer(http.Dir("public"))
	// http.Handle("/public/", http.StripPrefix("/public/", fileServer))

	fileServer := http.FileServer(http.Dir("public"))
	router.PathPrefix("/public/").Handler(http.StripPrefix("/public/", fileServer))

	router.HandleFunc("/", env.homePage)
	router.HandleFunc("/upload", env.upload)
	router.HandleFunc("/streets", env.streetIndex)
	router.HandleFunc("/builds", env.buildIndex)
	router.HandleFunc("/flats", env.flatsIndex)
	router.HandleFunc("/lics", env.licsIndex)
	router.HandleFunc("/infoLic", env.infoLicIndex)
	router.HandleFunc("/infoDataCounter", env.infoDataCounter)

	router.HandleFunc("/infoDataCounterValue", env.infoDataCounterValue)
	router.HandleFunc("/infoDataValue", env.infoDataValue)
	router.HandleFunc("/infoDataPaym", env.infoDataPaym)
	srv.server.Handler = router
	return srv
}

// Start launches the phishing server, listening on the configured address.
func (srv *Bill18Server) Start() error {
	log.Infof("Starting http-server at http://%s", srv.config.Listen)
	return srv.server.ListenAndServe()
}

// Shutdown attempts to gracefully shutdown the server.
func (srv *Bill18Server) Shutdown() error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	return srv.server.Shutdown(ctx)
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
	if env.cfg.IsPrettyJSON {
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

// upload logic
func (env *Env) upload(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method)
	if r.Method == "GET" {
		crutime := time.Now().Unix()
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(crutime, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))
		fmt.Println("token:", token)
		t, _ := template.ParseFiles("upload.html")
		//t.ExecuteTemplate(w, "Upload", &struct{ token string })
		//t.Execute(w, token)

		t, err := template.ParseFiles("templates/upload.html", "templates/header.html", "templates/footer.html")
		if err != nil {
			log.Error("template.ParseFiles", err.Error())
			return
		}
		t.ExecuteTemplate(w, "Upload", &struct {
			Listen string
			Token  string
		}{env.cfg.Listen, token}) //&struct{ token string }{token})

	} else {
		r.ParseMultipartForm(32 << 20)
		file, handler, err := r.FormFile("uploadfile")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()
		fmt.Fprintf(w, "%v", handler.Header)
		f, err := os.OpenFile("./loadfiles/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer f.Close()
		io.Copy(f, file)
	}
}
