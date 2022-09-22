package fiber

import (
	"deepfit/configs"
	"log"

	"github.com/gofiber/fiber/v2"
)

func Router() {
	app := fiber.New()

	api := app.Group(configs.API_VERSION_BASE_PATH)

	api.Post(configs.REGISTER, RegisterHandler)

	log.Fatal(app.Listen(":" + configs.PORT))
}
