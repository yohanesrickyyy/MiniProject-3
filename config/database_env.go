package config

import (
	"github.com/kelseyhightower/envconfig"
)

type DBEnv struct {
	Hostname string `envconfig:"HOSTNAME"`
	Username string `envconfig:"USERNAME"`
	Password string `envconfig:"PASSWORD"`
	Port     string `envconfig:"PORT"`
	DBName   string `envconfig:"DBNAME"`
}

func InitDB() DBEnv {
	var dbEnv DBEnv
	err := envconfig.Process("DATABASE", &dbEnv)
	if err != nil {
		panic(err)
	}
	return dbEnv
}
