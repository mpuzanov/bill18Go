package main

import (
	"log"

	"github.com/mpuzanov/bill18Go/internal/config"
	"github.com/mpuzanov/bill18Go/internal/store"
	"github.com/mpuzanov/bill18Go/internal/web"

	_ "github.com/denisenkom/go-mssqldb"

	"github.com/mpuzanov/bill18Go/pkg/logger"
	flag "github.com/spf13/pflag"
)

func main() {
	var cfgPath string
	flag.StringVarP(&cfgPath, "config", "c", "", "path to the configuration file")
	flag.Parse()

	cfg, err := config.LoadConfig(cfgPath)
	if err != nil {
		log.Fatalf("Не удалось загрузить %s: %s", cfgPath, err)
	}

	logger := logger.NewLogger(cfg.Log)

	db, err := store.NewStorageDB(cfg.DB)
	if err != nil {
		log.Fatalf("NewStorage failed: %s", err)
	}

	if err := web.Start(cfg, logger, db); err != nil {
		log.Fatal(err)
	} // //ожидаем завершение программы по Ctrl-C
}
