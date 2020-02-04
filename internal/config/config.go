package config

import (
	"io/ioutil"

	"github.com/go-yaml/yaml"
)

//Config - структура для считывания конфигурационного файла
type Config struct {
	Listen       string `yaml:"listen"`
	IP           string `yaml:"ip"`
	Port         string `yaml:"port"`
	LogLevel     string `yaml:"loglevel"`
	LogToFile    bool   `yaml:"logToFile"`
	LogFileName  string `yaml:"logFileName"`
	IsPrettyJSON bool   `yaml:"isPrettyJSON"`
	DatabaseURL  string `yaml:"database_url"` // строка подключения к БД
}

//ReadConfig ...
func ReadConfig(ConfigName string) (x *Config, err error) {
	var file []byte
	if file, err = ioutil.ReadFile(ConfigName); err != nil {
		return nil, err
	}
	x = new(Config)
	if err = yaml.Unmarshal(file, x); err != nil {
		return nil, err
	}
	if x.LogLevel == "" {
		x.LogLevel = "Trace"
	}
	x.Listen = x.IP + ":" + x.Port
	return x, nil
}
