package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"

	"./models"
	_ "github.com/denisenkom/go-mssqldb"
)

var (
	server       = "adm"
	port         = 1433
	user         = "sa"
	password     = "123"
	database     = "kv_all"
	isPrettyJSON = true
)

func main() {

	// Build connection string
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;",
		server, user, password, port, database)

	models.InitDB(connString)

	fmt.Println("Listening on port :8080")
	http.HandleFunc("/", homePage)
	http.HandleFunc("/streets", streetIndex)
	http.HandleFunc("/builds", buildIndex)
	http.HandleFunc("/flats", flatsIndex)
	http.HandleFunc("/lics", licsIndex)
	http.HandleFunc("/infoLic", infoLicIndex)
	http.ListenAndServe(":8080", nil)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/index.html", "templates/header.html", "templates/footer.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}

	t.ExecuteTemplate(w, "index", nil)
}

//prettyprint Делаем красивый json с отступами
func prettyprint(b []byte) ([]byte, error) {
	var out bytes.Buffer
	err := json.Indent(&out, b, "", "    ")
	return out.Bytes(), err
}

func getJSONResponse(w http.ResponseWriter, r *http.Request, data interface{}) {
	//jsData, err := json.MarshalIndent(data, "", "    ")
	jsData, err := json.Marshal(data)
	if err != nil {
		// handle error
	}
	if isPrettyJSON {
		jsData, err = prettyprint(jsData)
		if err != nil {
			// handle error
		}
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsData)
}

func streetIndex(w http.ResponseWriter, r *http.Request) {
	streets, err := models.AllStreets()
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}
	getJSONResponse(w, r, streets)
}

func buildIndex(w http.ResponseWriter, r *http.Request) {
	streetName := r.FormValue("street_name")
	if streetName == "" {
		streetName = "1-я Донская ул."
	}
	builds, err := models.GetBuilds(streetName)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}
	// for _, build := range builds.DataBuilds {
	// 	fmt.Printf("%s, %s, %s\n", builds.Street_name, build.Nom_dom, build.Nom_dom_sort)
	// }
	getJSONResponse(w, r, builds)
}

func flatsIndex(w http.ResponseWriter, r *http.Request) {
	streetName := r.FormValue("street_name")
	nomDom := r.FormValue("nom_dom")
	if streetName == "" {
		streetName = "1-я Донская ул."
	}
	if nomDom == "" {
		nomDom = "6"
	}
	flats, err := models.GetFlats(streetName, nomDom)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}
	// for _, flat := range flats {
	// 	fmt.Fprintf(w, "%s, %s, %s, %s\n", streetName, nomDom, flat.Nom_kvr, flat.Nom_kvr_sort)
	// }
	getJSONResponse(w, r, flats)
}

func licsIndex(w http.ResponseWriter, r *http.Request) {
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
	lics, err := models.GetKvrLic(streetName, nomDom, nomKvr)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}
	getJSONResponse(w, r, lics)
}

func infoLicIndex(w http.ResponseWriter, r *http.Request) {
	occ := r.FormValue("occ")
	if occ == "" {
		occ = "0"
	}
	fmt.Printf("FormValue: %s\n", occ)
	data, err := models.GetDataOcc(occ)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}
	getJSONResponse(w, r, data)
}
