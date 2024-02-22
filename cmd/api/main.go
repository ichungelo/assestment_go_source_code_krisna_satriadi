package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/ichungelo/assestment_go_source_code_krisna_satriadi/config"
	appconfig "github.com/ichungelo/assestment_go_source_code_krisna_satriadi/config/app_config"
	"github.com/ichungelo/assestment_go_source_code_krisna_satriadi/di"
	fibermiddleware "github.com/ichungelo/assestment_go_source_code_krisna_satriadi/transports/fiber/fiber_middleware"
	utillogger "github.com/ichungelo/assestment_go_source_code_krisna_satriadi/utils/util_logger"
)

func main() {
	cfg := config.GetEnv()

	_, err := appconfig.CheckEnumStage(cfg.AppConfig.Stage)
	if err != nil {
		utillogger.Error(err, nil)
		log.Panic(err)
	}

	app := fiber.New()
	router := di.Initializer(cfg)

	fibermiddleware.FiberMiddleware(app)
	router.Route(app)
	log.Fatal(app.Listen(fmt.Sprintf("%s:%s", cfg.AppConfig.Host, cfg.AppConfig.Port)))
}
