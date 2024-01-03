package fibermiddleware

import (
	"encoding/json"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/ichungelo/assestment_go_source_code_krisna_satriadi/utils"
)

// ! Fiber Middleware
func FiberMiddleware(a *fiber.App) {
	mapformat := map[string]string{
		"time":       fmt.Sprintf("${%s}", logger.TagTime),
		"httpStatus": fmt.Sprintf("${%s}", logger.TagStatus),
		"httpMethod": fmt.Sprintf("${%s}", logger.TagMethod),
		"path":       fmt.Sprintf("${%s}", logger.TagPath),
	}

	jsonByte, _ := json.MarshalIndent(mapformat, "", "	")
	jsonString := fmt.Sprintf("%s\n", string(jsonByte))
	a.Use(
		cors.New(),
		LoggerMiddleware(),
		logger.New(
			logger.Config{
				Format:     jsonString,
				TimeFormat: "2006-01-02T15:04:05.999999999Z07:00",
			},
		),
	)
}

func LoggerMiddleware() func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		utils.Request("start request", c)

		return c.Next()
	}
}
