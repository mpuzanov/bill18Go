package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"os"

	"net/http"
	"strconv"

	"github.com/mpuzanov/bill18Go/models"
	"github.com/sirupsen/logrus"

	_ "github.com/denisenkom/go-mssqldb"
)

const (
	configFileName = "files/conf.yaml"
	logFileName    = "files/logfile.txt"
)

var (
	cfg          *Config
	isPrettyJSON bool
	log          = logrus.New()
)

//Env Интерфейс для вызова функций sql
type Env struct {
	db models.Datastore
}

func main() {

	//=====================================================================
	cfg, err := reloadConfig(configFileName) //загружаем конфиг
	if err != nil {
		if err != errNotModified {
			log.Fatalf("Не удалось загрузить %s: %s", configFileName, err)
		}
	}
	log.Formatter = new(logrus.TextFormatter)
	log.Formatter.(*logrus.TextFormatter).TimestampFormat = "02-01-2006 15:04:05"
	log.Formatter.(*logrus.TextFormatter).FullTimestamp = true

	level, err := logrus.ParseLevel(cfg.LogLevel)
	if err != nil {
		log.Error(err)
	}
	log.SetLevel(level)

	if cfg.LogToFile {
		file, err := os.OpenFile(logFileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
		if err == nil {
			log.Infoln("logToFile: ", logFileName)
			log.Out = file
		} else {
			log.Info("Failed to log to file, using default stderr")
		}
	}
	log.Infoln("log.Level:", log.Level)
	isPrettyJSON = cfg.IsPrettyJSON
	//=====================================================================
	//connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;", cfg.Server, cfg.User, cfg.Password, cfg.Port, cfg.Database)
	connString := fmt.Sprintf("sqlserver://%s:%s@%s:%d?database=%s",
		cfg.SQLConnect.User, cfg.SQLConnect.Password, cfg.SQLConnect.Server, cfg.SQLConnect.Port, cfg.SQLConnect.Database)

	db, err := models.GetInitDB(connString)
	if err != nil {
		panic(err)
	}
	env := &Env{db}
	//=====================================================================

	log.Info("Listening on port :8080")
	http.HandleFunc("/", homePage)
	http.HandleFunc("/streets", env.streetIndex)
	http.HandleFunc("/builds", env.buildIndex)
	http.HandleFunc("/flats", env.flatsIndex)
	http.HandleFunc("/lics", env.licsIndex)
	http.HandleFunc("/infoLic", env.infoLicIndex)
	http.HandleFunc("/infoDataCounter", env.infoDataCounter)

	http.HandleFunc("/infoDataCounterValue", env.infoDataCounterValue)
	http.HandleFunc("/infoDataValue", env.infoDataValue)
	http.HandleFunc("/infoDataPaym", env.infoDataPaym)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func homePage(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/index.html", "templates/header.html", "templates/footer.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}

	t.ExecuteTemplate(w, "index", nil)
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
func getJSONResponse(w http.ResponseWriter, r *http.Request, data interface{}) {
	//jsData, err := json.MarshalIndent(data, "", "    ")
	jsData, err := json.Marshal(data)
	if err != nil {
		http.Error(w, "Error creating JSON response", http.StatusInternalServerError)
		log.Println(err)
		return
	}
	if isPrettyJSON == true {
		jsData, err = prettyprint(jsData)
		if err != nil {
			// handle error
		}
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsData)
}

func (env *Env) streetIndex(w http.ResponseWriter, r *http.Request) {
	data, err := env.db.GetAllStreets()
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}
	log.Traceln("streetIndex")
	getJSONResponse(w, r, data)
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
	getJSONResponse(w, r, data)
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
	// for _, flat := range flats {
	// 	fmt.Fprintf(w, "%s, %s, %s, %s\n", streetName, nomDom, flat.Nom_kvr, flat.Nom_kvr_sort)
	// }
	log.Traceln("flatsIndex", streetName, nomDom)
	getJSONResponse(w, r, data)
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
	getJSONResponse(w, r, data)
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
	getJSONResponse(w, r, data)
}

func (env *Env) infoDataCounter(w http.ResponseWriter, r *http.Request) {
	parOcc := r.FormValue("occ")
	if parOcc == "" {
		parOcc = "0"
	}
	occ, _ := strconv.Atoi(parOcc)

	data, err := env.db.GetCounterByOcc(occ)
	if err != nil {
		fmt.Printf("infoDataCounter: %s\n", err)
		http.Error(w, http.StatusText(500), 500)
		return
	}
	log.Traceln("infoDataCounter", parOcc)
	getJSONResponse(w, r, data)
}

func (env *Env) infoDataCounterValue(w http.ResponseWriter, r *http.Request) {
	parOcc := r.FormValue("occ")
	if parOcc == "" {
		parOcc = "0"
	}
	occ, _ := strconv.Atoi(parOcc)

	data, err := env.db.GetCounterValueByOcc(occ)
	if err != nil {
		fmt.Printf("infoDataCounterValue: %s\n", err)
		http.Error(w, http.StatusText(500), 500)
		return
	}
	log.Traceln("infoDataCounterValue", parOcc)
	getJSONResponse(w, r, data)
}

func (env *Env) infoDataValue(w http.ResponseWriter, r *http.Request) {
	parOcc := r.FormValue("occ")
	if parOcc == "" {
		parOcc = "0"
	}
	occ, _ := strconv.Atoi(parOcc)

	data, err := env.db.GetDataValueByOcc(occ)
	if err != nil {
		fmt.Printf("infoDataValue: %s\n", err)
		http.Error(w, http.StatusText(500), 500)
		return
	}
	log.Traceln("infoDataValue", parOcc)
	getJSONResponse(w, r, data)
}

func (env *Env) infoDataPaym(w http.ResponseWriter, r *http.Request) {
	parOcc := r.FormValue("occ")
	if parOcc == "" {
		parOcc = "0"
	}
	occ, _ := strconv.Atoi(parOcc)
	data, err := env.db.GetDataPaymByOcc(occ)
	if err != nil {
		fmt.Printf("infoDataPaym: %s\n", err)
		http.Error(w, http.StatusText(500), 500)
		return
	}
	log.Traceln("infoDataPaym", parOcc)
	getJSONResponse(w, r, data)
}
