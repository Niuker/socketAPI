package config

import (
	"github.com/jinzhu/configor"
)

type Config struct {
	DB struct {
		User     string `default:"root"`
		Password string `required:"true" env:"DBPassword"`
		Port     string `default:"3306"`
	}
	NET struct {
		Http   string `default:":8091"`
		Socket string `default:"8000"`
	}
	ENCR struct {
		Desckey string `default:"14725891"`
	}

	//Contacts []struct {
	//	Name  string
	//	Email string `required:"true"`
	//}
}

func get() *Config {
	var conf = Config{}
	err := configor.Load(&conf, "config.yml")
	if err != nil {
		panic(err)
	}
	return &conf
}

var MyConfig *Config = get()
