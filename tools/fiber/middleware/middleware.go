package middleware

import (
	"deepfit/constants"
	"deepfit/pkg"
	"deepfit/tools/jwt"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func TokenCanGo(c *fiber.Ctx) error {

	authorization := c.Get("Authorization")
	if !strings.Contains(authorization, "Bearer ") {
		return c.Status(fiber.StatusUnauthorized).JSON(pkg.Response{
			Status:  false,
			Message: constants.UNAUTHORIZED_USER,
		})
	}

	splittedAuthorization := strings.Split(authorization, " ")
	if len(splittedAuthorization) != 2 {
		return c.Status(fiber.StatusUnauthorized).JSON(pkg.Response{
			Status:  false,
			Message: constants.UNAUTHORIZED_USER,
		})
	}

	token := splittedAuthorization[1]

	_jwt := jwt.New().SetToken(token).DecodeToken()

	c.Locals("userId", _jwt.GetUserId())
	c.Locals("is_verified", _jwt.GetIsVerified())

	return c.Next()
}

func TokenCantGo(c *fiber.Ctx) error {

	authorization := c.Get("Authorization")

	if len(authorization) != 0 || authorization != "" {
		return c.Status(fiber.StatusUnauthorized).JSON(pkg.Response{
			Status:  false,
			Message: constants.AUTHORIZED_USER,
		})
	}

	return c.Next()
}

func VerifiedCanGo(c *fiber.Ctx) error {

	if c.Locals("is_verified") == false {
		return c.Status(fiber.StatusUnauthorized).JSON(pkg.Response{
			Status:  false,
			Message: constants.UNVERIFIED_USER,
		})
	}

	return c.Next()
}

func VerifiedCantGo(c *fiber.Ctx) error {

	if c.Locals("is_verified") == true {
		return c.Status(fiber.StatusUnauthorized).JSON(pkg.Response{
			Status:  false,
			Message: constants.VERIFIED_USER,
		})
	}

	return c.Next()
}
