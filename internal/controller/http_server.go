package controller

import (
	"context"
	"crypto/md5"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"

	"strconv"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	"github.com/mpuzanov/bill18Go/internal/config"
	log "github.com/mpuzanov/bill18Go/internal/logger"
)

var (
	usersAPI = map[string]string{
		"admin": "admin",
		"etton": "oxnpfAXv",
	}
)

//Bill18Server my http-server
type Bill18Server struct {
	server *http.Server
	config *config.Config
}

//NewServer ...
func NewServer(cfg *config.Config) *Bill18Server {
	log.Info("Запуск веб-сервера на ", cfg.Listen)

	//=====================================================================
	db, err := GetInitDB(cfg.DatabaseURL)
	if err != nil {
		panic(err)
	}
	env = &Env{Db: db, Cfg: cfg}
	log.Traceln("connString:", cfg.DatabaseURL)
	//=====================================================================

	defaultServer := &http.Server{
		Addr:           cfg.IP + ":" + cfg.Port,
		Handler:        nil,              // if nil use default http.DefaultServeMux
		ReadTimeout:    10 * time.Second, // max duration reading entire request
		WriteTimeout:   10 * time.Second, // max timing write response
		IdleTimeout:    15 * time.Second, // max time wait for the next request
		MaxHeaderBytes: 1 << 20,          // 2^20 or 128 kb
	}

	srv := &Bill18Server{
		server: defaultServer,
		config: cfg,
	}

	CreateTemplate()

	//router := http.NewServeMux()
	router := mux.NewRouter()

	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./public/static/"))))

	// для отдачи сервером статичных файлов из папки public/static
	//fs := http.FileServer(http.Dir("./public/static/"))
	//http.Handle("/static/", http.StripPrefix("/static/", fs))

	router.HandleFunc("/", env.HomePage)
	router.HandleFunc("/upload", env.upload)
	router.HandleFunc("/testapi", env.Testapi)
	router.HandleFunc("/api/streets", BasicAuth(env.StreetIndex))
	router.HandleFunc("/api/builds/{street_name}", BasicAuth(env.BuildIndex))
	router.HandleFunc("/api/flats/{street_name}/{nom_dom:[0-9]+}", BasicAuth(env.FlatsIndex))
	router.HandleFunc("/api/lics/{street_name}/{nom_dom:[0-9]+}/{nom_kvr:[0-9]+}", BasicAuth(env.LicsIndex))
	router.HandleFunc("/api/infoLic/{occ:[0-9]+}", BasicAuth(env.InfoLicIndex))
	router.HandleFunc("/api/infoDataCounter/{occ:[0-9]+}", BasicAuth(env.InfoDataCounter))
	router.HandleFunc("/api/infoDataCounterValue/{occ:[0-9]+}", BasicAuth(env.InfoDataCounterValue))
	router.HandleFunc("/api/infoDataValue/{occ:[0-9]+}", BasicAuth(env.InfoDataValue))
	router.HandleFunc("/api/infoDataPaym/{occ:[0-9]+}", BasicAuth(env.InfoDataPaym))
	router.HandleFunc("/api/puAddValue/{puId:[0-9]+}/{value:[0-9]+}", BasicAuth(env.PuAddValue))
	router.HandleFunc("/api/puDelValue/{puId:[0-9]+}/{id:[0-9]+}", BasicAuth(env.PuDelValue))
	router.HandleFunc("/api/infoDataCounterValueTip/{tip:[0-9]+}", BasicAuth(env.InfoDataCounterValueTip))
	//router=handlers.CompressHandler(router)
	//router=handlers.LoggingHandler(log.Logger.Out, router) //os.Stdout
	srv.server.Handler = handlers.LoggingHandler(log.Logger.Out, router)
	return srv
}

// Start launches the phishing server, listening on the configured address.
func (srv *Bill18Server) Start() error {
	log.Infof("Starting http-server at http://%s", srv.server.Addr)
	return srv.server.ListenAndServe()
}

// Shutdown attempts to gracefully shutdown the server.
func (srv *Bill18Server) Shutdown() error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	return srv.server.Shutdown(ctx)
}

// func logger(next http.HandlerFunc, cfg *config.Config) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Printf("server [net/http] method [%s]  connection from [%v]\n", r.Method, r.RemoteAddr)
// 		next.ServeHTTP(w, r)
// 	}
// }

//BasicAuth ...
func BasicAuth(handler http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		realm := "Please enter your username and password"
		user, pass, ok := r.BasicAuth()
		//fmt.Println(user, pass, ok)
		//if !ok || subtle.ConstantTimeCompare([]byte(user), []byte(userAPI)) != 1 || subtle.ConstantTimeCompare([]byte(pass), []byte(userAPIPsw)) != 1 {
		if !ok || !validate(user, pass) {
			w.Header().Set("WWW-Authenticate", `Basic realm="`+realm+`"`)
			w.WriteHeader(401)
			w.Write([]byte("You are Unauthorized to access the application.\n"))
			return
		}
		//fmt.Println("BasicAuth ok")
		handler(w, r)
	}
}

func validate(username, password string) bool {
	//Basic dGVzdDp0ZXN0
	return usersAPI[username] == password
}

// upload logic
func (env *Env) upload(w http.ResponseWriter, r *http.Request) {
	//fmt.Println("method:", r.Method)
	if r.Method == "GET" {
		crutime := time.Now().Unix()
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(crutime, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))
		//fmt.Println("token:", token)

		t, err := template.ParseFiles("templates/base.html", "templates/header.html", "templates/upload.html")
		if err != nil {
			log.Error("template.ParseFiles", err.Error())
			return
		}
		t.ExecuteTemplate(w, "base", &struct {
			Listen string
			Token  string
		}{env.Cfg.Listen, token})

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
