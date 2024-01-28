package config

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
)

// type to devine stage
type Stage string

const (
	DEV     Stage = "dev"
	STAGING Stage = "staging"
	PROD    Stage = "prod"
)

func CheckStage(stage string) (Stage, error) {
	mapStage := map[string]Stage{
		"dev":     DEV,
		"staging": STAGING,
		"prod":    PROD,
	}

	if v, ok := mapStage[stage]; ok {
		return v, nil
	}

	return "", fmt.Errorf("stage \"%s\" is not on list. select \"%s\", \"%s\", or \"%s\"", stage, DEV, STAGING, PROD)
}

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}
}
