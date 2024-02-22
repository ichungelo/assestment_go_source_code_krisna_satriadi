package config

import (
	"sync"

	appconfig "github.com/ichungelo/assestment_go_source_code_krisna_satriadi/config/app_config"
	dbconfig "github.com/ichungelo/assestment_go_source_code_krisna_satriadi/config/db_config"
)

type Config struct {
	DBConfig  *dbconfig.DbConfig
	AppConfig *appconfig.AppConfig
}

var (
	cfg  Config
	once sync.Once
)

func GetEnv() *Config {
	once.Do(func() {
		cfg = Config{
			DBConfig:  dbconfig.NewDbConfig(),
			AppConfig: appconfig.NewAppConfig(),
		}
	})
	return &cfg
}
