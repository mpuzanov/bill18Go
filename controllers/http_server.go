package controllers

import (
	"context"
	"crypto/md5"
	"crypto/subtle"
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
)

const (
	userAPI    = "admin"
	userAPIPsw = "admin"
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
	env = &Env{db: db, cfg: cfg}
	log.Traceln("connString:", cfg.DatabaseURL)
	//=====================================================================

	defaultServer := &http.Server{
		Addr:           cfg.Listen,
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

	//router := http.NewServeMux()
	router := mux.NewRouter()

	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./public/static/"))))

	// для отдачи сервером статичных файлов из папки public/static
	//http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./public/static"))))

	// static-files
	// fs := http.FileServer(http.Dir("./static/"))
	// http.Handle("/static/", http.StripPrefix("/static/", fs))

	//fileServer := http.FileServer(http.Dir("./public/"))
	//http.Handle("/static/", http.StripPrefix("/static/", fileServer))

	//static-files-gorilla-mux
	//router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))

	// http.HandleFunc("/", handler)
	// http.HandleFunc("/testapi", handlerTestapi)
	// http.ListenAndServe(":8090", nil)

	router.HandleFunc("/", env.homePage)
	router.HandleFunc("/upload", env.upload)
	router.HandleFunc("/testapi", env.testapi)
	router.HandleFunc("/api/streets", BasicAuth(logger(env.streetIndex)))
	router.HandleFunc("/api/builds", BasicAuth(logger(env.buildIndex)))
	router.HandleFunc("/api/flats", BasicAuth(logger(env.flatsIndex)))
	router.HandleFunc("/api/lics", BasicAuth(logger(env.licsIndex)))
	router.HandleFunc("/api/infoLic", BasicAuth(logger(env.infoLicIndex)))
	router.HandleFunc("/api/infoDataCounter", BasicAuth(logger(env.infoDataCounter)))
	router.HandleFunc("/api/infoDataCounterValue", BasicAuth(logger(env.infoDataCounterValue)))
	router.HandleFunc("/api/infoDataValue", BasicAuth(logger(env.infoDataValue)))
	router.HandleFunc("/api/infoDataPaym", BasicAuth(logger(env.infoDataPaym)))
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

func logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		fmt.Printf("server [net/http] method [%s]  connection from [%v]\n", r.Method, r.RemoteAddr)

		next.ServeHTTP(w, r)
	}
}

//BasicAuth ...
func BasicAuth(handler http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		realm := "Please enter your username and password"
		user, pass, ok := r.BasicAuth()

		if !ok || subtle.ConstantTimeCompare([]byte(user), []byte(userAPI)) != 1 || subtle.ConstantTimeCompare([]byte(pass), []byte(userAPIPsw)) != 1 {
			w.Header().Set("WWW-Authenticate", `Basic realm="`+realm+`"`)
			w.WriteHeader(401)
			w.Write([]byte("You are Unauthorized to access the application.\n"))
			return
		}

		handler(w, r)
	}
}

//BasicAuth ...
// func BasicAuth(next http.HandlerFunc) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {

// 		auth := strings.SplitN(r.Header.Get("Authorization"), " ", 2)

// 		if len(auth) != 2 || auth[0] != "Basic" {
// 			http.Error(w, "authorization failed", http.StatusUnauthorized)
// 			return
// 		}

// 		payload, _ := base64.StdEncoding.DecodeString(auth[1])
// 		pair := strings.SplitN(string(payload), ":", 2)

// 		if len(pair) != 2 || !validate(pair[0], pair[1]) {
// 			http.Error(w, "authorization failed", http.StatusUnauthorized)
// 			return
// 		}

// 		next.ServeHTTP(w, r)
// 	}
// }

func validate(username, password string) bool {
	if username == "test" && password == "test" { //Basic dGVzdDp0ZXN0
		return true
	}
	return false
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
		}{env.cfg.Listen, token})

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
