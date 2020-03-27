package web

import (
	"context"
	"crypto/md5"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	"go.uber.org/zap"

	"github.com/mpuzanov/bill18Go/internal/config"
	"github.com/mpuzanov/bill18Go/internal/store"
)

var (
	usersAPI = map[string]string{
		"admin": "admin",
		"etton": "oxnpfAXv",
	}
)

type myHandler struct {
	router *mux.Router
	logger *zap.Logger
	cfg    *config.Config
	Db     store.Datastore
}

// Start Запуск сервера
func Start(cfg *config.Config, logger *zap.Logger, db store.Datastore) error {

	logger.Info("Запуск веб-сервера", zap.String("http addres", cfg.HTTPAddr))

	handler := &myHandler{
		router: mux.NewRouter(),
		logger: logger,
		cfg:    cfg,
		Db:     db,
	}
	handler.configRouter()

	logger.Sugar().Infof("Database: %s, host: %s", cfg.DB.Database, cfg.DB.Host)
	//=====================================================================

	srv := &http.Server{
		Addr:           cfg.HTTPAddr,     // cfg.IP + ":" + cfg.Port,
		Handler:        handler,          // if nil use default http.DefaultServeMux
		ReadTimeout:    10 * time.Second, // max duration reading entire request
		WriteTimeout:   10 * time.Second, // max timing write response
		IdleTimeout:    15 * time.Second, // max time wait for the next request
		MaxHeaderBytes: 1 << 20,          // 2^20 or 128 kb
	}

	//запускаем веб-сервер
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
	log.Printf("Server http started: %s, file log: %s\n", srv.Addr, cfg.Log.File)
	logger.Info("Starting Http server", zap.String("address", srv.Addr))

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-done
	log.Print("Server stopped")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {
		cancel()
	}()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server shutdown failed:%+v", err)
	}
	log.Println("Shutdown done")
	os.Exit(0)

	return nil
}

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
			_, err := w.Write([]byte("You are Unauthorized to access the application.\n"))
			if err != nil {
				log.Println("Fail Write Unauthorized")
				return
			}
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
func (s *myHandler) upload(w http.ResponseWriter, r *http.Request) {
	//fmt.Println("method:", r.Method)
	if r.Method == "GET" {
		crutime := time.Now().Unix()
		h := md5.New()
		_, err := io.WriteString(h, strconv.FormatInt(crutime, 10))
		if err != nil {
			s.logger.Error("Fail WriteString", zap.Error(err))
			return
		}
		token := fmt.Sprintf("%x", h.Sum(nil))
		//fmt.Println("token:", token)

		err = t["upload.html"].ExecuteTemplate(w, "base", &struct {
			Listen string
			Token  string
		}{s.cfg.HTTPAddr, token})
		if err != nil {
			s.logger.Error("Fail ExecuteTemplate upload", zap.Error(err))
			return
		}

	} else {
		if err := r.ParseMultipartForm(32 << 20); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		file, handler, err := r.FormFile("uploadfile")
		if err != nil {
			s.logger.Error("Fail FormFile", zap.Error(err))
			return
		}
		defer file.Close()
		fmt.Fprintf(w, "%v", handler.Header)
		f, err := os.OpenFile("./loadfiles/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			s.logger.Error("Fail OpenFile", zap.Error(err))
			return
		}
		defer f.Close()
		_, err = io.Copy(f, file)
		if err != nil {
			s.logger.Error("Fail Copy", zap.Error(err))
			return
		}
	}
}
