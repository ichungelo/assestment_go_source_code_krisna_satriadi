package appconfig

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

const (
	DEV     = "dev"
	STAGING = "staging"
	PROD    = "prod"
)

func CheckEnumStage(stage string) (string, error) {
	mapStage := map[string]string{
		"dev":     DEV,
		"staging": STAGING,
		"prod":    PROD,
	}

	if v, ok := mapStage[stage]; ok {
		return v, nil
	}

	return "", fmt.Errorf("stage \"%s\" is not on list. select \"%s\", \"%s\", or \"%s\"", stage, DEV, STAGING, PROD)
}

type AppConfig struct {
	Name  string `envconfig:"APP_NAME" default:"my_app"`
	Host  string `envconfig:"APP_HOST" default:"8080"`
	Port  string `envconfig:"APP_PORT" default:"8080"`
	Stage string `envconfig:"APP_STAGE" default:"dev"`
}

func NewAppConfig() *AppConfig {
	var appCfg AppConfig
	envconfig.MustProcess("", &appCfg)
	return &appCfg
}
