package config

import (
	"log"
	"net"
	"strings"

	"github.com/mpuzanov/bill18Go/pkg/logger"
	"github.com/spf13/viper"
)

//Config Структура файла с конфигурацией
type Config struct {
	Log          logger.LogConf `yaml:"log" mapstructure:"log"`
	HTTPAddr     string         `yaml:"http_listen" mapstructure:"http_listen"`
	Host         string         `yaml:"host" mapstructure:"http_host"`
	Port         string         `yaml:"port" mapstructure:"http_port"`
	IsPrettyJSON bool           `yaml:"isPrettyJSON" mapstructure:"isPrettyJSON"`
	DB           DBConf         `yaml:"db" mapstructure:"db"`
}

// DBConf .
type DBConf struct {
	Name     string `yaml:"name" mapstructure:"name"`
	Host     string `yaml:"host" mapstructure:"host"`
	Port     string `yaml:"port" mapstructure:"port"`
	User     string `yaml:"user" mapstructure:"user"`
	Password string `yaml:"password" mapstructure:"password"`
	Database string `yaml:"database" mapstructure:"database"`
	SSL      string `yaml:"ssl" mapstructure:"ssl"`
}

// LoadConfig Загрузка конфигурации из файла
func LoadConfig(filePath string) (*Config, error) {
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.SetDefault("log.level", "info")
	viper.SetDefault("log.file", "")
	//viper.SetDefault("http_host", getOutboundIP().String())
	viper.SetDefault("http_port", "8090")
	viper.SetDefault("db.port", "1433")
	//viper.SetDefault("database_url", "sqlserver://sa:123@localhost:1433?database=kv_all")

	if filePath != "" {
		log.Printf("Parsing config: %s\n", filePath)
		viper.SetConfigFile(filePath)
		viper.SetConfigType("yaml")
		//log.Println(viper.ConfigFileUsed())
		err := viper.ReadInConfig()
		if err != nil {
			return nil, err
		}
	} else {
		log.Println("Config file is not specified.")
	}
	//log.Println(viper.AllSettings())
	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}
	// if config.Host == "" {
	// 	config.Host = getOutboundIP().String()
	// }
	config.HTTPAddr = net.JoinHostPort(config.Host, config.Port)

	//log.Println(config)
	return &config, nil
}

//GetOutboundIP Get preferred outbound ip of this machine
// func getOutboundIP() net.IP {
// 	conn, err := net.Dial("udp", "8.8.8.8:80")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer conn.Close()
// 	localAddr := conn.LocalAddr().(*net.UDPAddr)
// 	return localAddr.IP
// }
