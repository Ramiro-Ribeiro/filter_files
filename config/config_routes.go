package config

import (
	"fmt"
	"read_files/util/constants"
	"strings"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func ConfigsRoutes() *fiber.App {
	c := cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: strings.Join([]string{
			fiber.MethodGet,
			fiber.MethodPost,
		}, ","),
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
	})

	app := fiber.New(fiber.Config{
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
		BodyLimit:   10 * 1024 * 1024,
	})

	app.Use(func(cxt *fiber.Ctx) error {
		if cxt.Path() == fmt.Sprintf("%s/%s/upload", constants.API, constants.V1) {
			cxt.Set("Content-Type", "application/zip")
			cxt.Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", constants.FileName))
		}

		return cxt.Next()
	})

	app.Use(c)

	return app
}
