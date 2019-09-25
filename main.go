package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/mpuzanov/bill18Go/apiserver"
	"github.com/mpuzanov/bill18Go/config"
	"github.com/mpuzanov/bill18Go/logger"

	_ "github.com/denisenkom/go-mssqldb"
	log "github.com/mpuzanov/bill18Go/logger"
)

//github.com/mpuzanov/bill18Go/
const (
	configFileName = "files/conf.yaml"
)

func main() {

	//загружаем конфиг
	cfg, err := config.ReadConfig(configFileName)
	if err != nil {
		log.Fatalf("Не удалось загрузить %s: %s", configFileName, err)
	}

	//инициализируем логгеры
	if err = logger.SetupLogger(cfg); err != nil {
		log.Fatal(err)
	}

	srv := apiserver.RunHTTP(cfg)

	//ожидаем завершение программы по Ctrl-C
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, syscall.SIGTERM)

	for {
		select {
		case <-sigChan:
			log.Info("CTRL-C: Завершаю работу.")
			srv.Shutdown(context.TODO())
			log.Info("Shutdown done")
			return
		}
	}
}
