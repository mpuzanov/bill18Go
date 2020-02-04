package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/mpuzanov/bill18Go/internal/config"
	"github.com/mpuzanov/bill18Go/internal/controller"

	_ "github.com/denisenkom/go-mssqldb"

	"github.com/mpuzanov/bill18Go/internal/logger"
)

const (
	configFileName = "conf.yml"
)

func main() {

	//загружаем конфиг
	cfg, err := config.ReadConfig(configFileName)
	if err != nil {
		log.Fatalf("Не удалось загрузить %s: %s", configFileName, err)
	}

	// get port env var
	portEnv := os.Getenv("PORT")
	if len(portEnv) > 0 {
		cfg.Port = portEnv
		cfg.Listen = cfg.IP + ":" + cfg.Port
	}
	DatabaseURLEnv := os.Getenv("DatabaseURL")
	if len(DatabaseURLEnv) > 0 {
		cfg.DatabaseURL = DatabaseURLEnv
	}
	//инициализируем логгеры
	if err = logger.SetupLogger(cfg); err != nil {
		log.Fatal(err)
	}

	srv := controller.NewServer(cfg)
	go srv.Start()

	//ожидаем завершение программы по Ctrl-C
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
	<-sigChan
	logger.Info("CTRL-C: Завершаю работу.")
	srv.Shutdown()
	logger.Info("Shutdown done")
}
