package fiber

import (
	"deepfit/configs"
	"deepfit/constants"
	"deepfit/tools/fiber/handler"
	"deepfit/tools/fiber/middleware"
	"log"

	"github.com/gofiber/fiber/v2"
)

func Router() {
	app := fiber.New()

	api := app.Group(constants.API_PREFIX)

	api.Post(constants.REGISTER, middleware.TokenCantGo, handler.RegisterHandler)

	log.Fatal(app.Listen(":" + configs.PORT))
}
