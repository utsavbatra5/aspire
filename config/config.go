package config

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v3"
)

var Conf = config{}

type config struct {
	Server
	Db     DbConfig
	JwtKey string
	Logger LoggerConfig
}

type Server struct {
	Port int32
}

type LoggerConfig struct {
	Level            string
	Logfilepath      string
	Isstderrenabled  bool
	Isconsoleenabled bool
	Isfileenabled    bool
	Maxbackups       int
	Maxsize          int
	Maxage           int
}

type DbConfig struct {
	Name     string
	Database string
	Host     string
	Password string
	Port     int32
	Username string
	Driver   string
}

func InitConfig() {

	yamlFile, err := ioutil.ReadFile("./config/conf.yaml")
	if err != nil {
		log.Printf("yamlFile.Get err #%v ", err)
	}

	err = yaml.Unmarshal(yamlFile, &Conf)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

}

func InitConfigForTest() {
	yamlFile, err := ioutil.ReadFile("../config/conf.yaml")
	if err != nil {
		log.Printf("yamlFile.Get err #%v ", err)
	}

	err = yaml.Unmarshal(yamlFile, &Conf)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
}
