package router

import (
	"fmt"
	"read_files/config"
	"read_files/router/handler"
	"read_files/util/constants"

	"github.com/gofiber/fiber/v2"
)

func InitializeRoutes() *fiber.App {
	app := config.ConfigsRoutes()

	api := app.Group(constants.API)

	v1 := api.Group(fmt.Sprint("/", constants.V1), func(c *fiber.Ctx) error {
		c.Set(constants.VERSION, constants.V1)
		return c.Next()
	})

	v1.Get("/health", handler.HealthCheck)

	v1.Post("/upload", handler.Upload)

	app.Get("/openapi.yaml", func(c *fiber.Ctx) error {
		c.Type("yaml")
		return c.SendFile("docs/openapi.yaml")
	})

	// Rota para Swagger UI
	app.Get("/docs", func(c *fiber.Ctx) error {
		c.Type("html")
		return c.SendString(`<!DOCTYPE html>
<html lang="pt-BR">
<head>
  <meta charset="UTF-8">
  <title>Filter Files - Swagger UI</title>
  <link rel="stylesheet" href="https://unpkg.com/swagger-ui-dist@5/swagger-ui.css" />
  <style>
    html, body { margin: 0; padding: 0; height: 100%; }
    #swagger-ui { height: 100%; }
  </style>
</head>
<body>
  <div id="swagger-ui"></div>
  <script src="https://unpkg.com/swagger-ui-dist@5/swagger-ui-bundle.js"></script>
  <script>
    window.onload = () => {
      window.ui = SwaggerUIBundle({
        url: '/openapi.yaml',
        dom_id: '#swagger-ui',
        presets: [SwaggerUIBundle.presets.apis],
        layout: 'BaseLayout'
      });
    };
  </script>
</body>
</html>`)
	})

	return app
}
