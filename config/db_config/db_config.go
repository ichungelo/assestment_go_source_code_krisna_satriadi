package dbconfig

import (
	"github.com/kelseyhightower/envconfig"
)

type DbConfig struct {
	DbHost     string `envconfig:"DB_HOST" default:"localhost"`
	DbPort     string    `envconfig:"DB_PORT" default:"5432"`
	DbUsername string `envconfig:"DB_USER" default:"root2"`
	DbPassword string `envconfig:"DB_PASSWORD" default:"root2"`
}

func NewDbConfig() *DbConfig {
	var dbCfg DbConfig
	envconfig.MustProcess("", &dbCfg)
	return &dbCfg
}
