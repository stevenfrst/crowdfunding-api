package config

import (
	"github.com/tkanos/gonfig"
	"log"
)

type Configuration struct {
	DB_USERNAME string
	DB_PASSWORD string
	DB_PORT string
	DB_HOST string
	DB_NAME string
	JWT_SECRET string
	JWT_EXPIRED int
	CONTEXT_TIMEOUT int
	SERVER_ADDR string
}

func GetConfig() Configuration {
	conf := Configuration{}
	err := gonfig.GetConf("app/config/config.json",&conf)
	if err != nil {
		log.Fatalln(err)
	}
	return conf
}