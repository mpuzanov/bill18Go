package main

import (
	"errors"
	"io/ioutil"
	"os"

	"github.com/go-yaml/yaml"
)

var (
	configModtime  int64
	errNotModified = errors.New("Not modified")
)

// SQLConnect - структура для считывания строки подключения к БД
type SQLConnect struct {
	Server   string `yaml:"server"`
	Port     int    `yaml:"port"`
	Database string `yaml:"database"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}

// Config - структура для считывания конфигурационного файла
type Config struct {
	LogLevel     string     `yaml:"loglevel"`
	IsPrettyJSON bool       `yaml:"isPrettyJSON"`
	SQLConnect   SQLConnect `yaml:"sqlconnect"`
}

func readConfig(ConfigName string) (x *Config, err error) {
	var file []byte
	if file, err = ioutil.ReadFile(ConfigName); err != nil {
		return nil, err
	}
	x = new(Config)
	if err = yaml.Unmarshal(file, x); err != nil {
		return nil, err
	}
	if x.LogLevel == "" {
		x.LogLevel = "Debug"
	}
	return x, nil
}

//Проверяет время изменения конфигурационного файла
//и перезагружает его если он изменился
//Возвращает errNotModified если изменений нет
func reloadConfig(configName string) (cfg *Config, err error) {
	info, err := os.Stat(configName)
	if err != nil {
		return nil, err
	}
	if configModtime != info.ModTime().UnixNano() {
		configModtime = info.ModTime().UnixNano()
		cfg, err = readConfig(configName)
		if err != nil {
			return nil, err
		}
		return cfg, nil
	}
	return nil, errNotModified
}
