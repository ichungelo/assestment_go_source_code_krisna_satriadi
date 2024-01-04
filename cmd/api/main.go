package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/ichungelo/assestment_go_source_code_krisna_satriadi/config"
	"github.com/ichungelo/assestment_go_source_code_krisna_satriadi/di"
	gormconnection "github.com/ichungelo/assestment_go_source_code_krisna_satriadi/sources/gorm/gorm_connection"
	fibermiddleware "github.com/ichungelo/assestment_go_source_code_krisna_satriadi/transports/fiber/fiber_middleware"
	"github.com/ichungelo/assestment_go_source_code_krisna_satriadi/utils"
)

func init() {
	config.LoadEnv()
	appStage := os.Getenv("APP_STAGE")

	stage, err := config.CheckStage(appStage)
	if err != nil {
		utils.Error(err, nil)
		log.Panic(err)
	}

	gormconnection.GetInstanceDB(stage)
}

func main() {
	var (
		appHost = os.Getenv("APP_HOST")
		appPort = os.Getenv("APP_PORT")
	)
	app := fiber.New()
	router := di.Initializer(gormconnection.DB)

	fibermiddleware.FiberMiddleware(app)
	router.Route(app, fibermiddleware.LoggerMiddleware())
	log.Fatal(app.Listen(fmt.Sprintf("%s:%s", appHost, appPort)))
}
