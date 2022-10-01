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
		panic(pkg.NewError(fiber.StatusUnauthorized, constants.UNAUTHORIZED_USER, nil))
	}

	splittedAuthorization := strings.Split(authorization, " ")
	if len(splittedAuthorization) != 2 {
		panic(pkg.NewError(fiber.StatusUnauthorized, constants.UNAUTHORIZED_USER, nil))
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
		panic(pkg.NewError(fiber.StatusUnauthorized, constants.AUTHORIZED_USER, nil))
	}

	return c.Next()
}

func VerifiedCanGo(c *fiber.Ctx) error {

	if c.Locals("is_verified") == false {
		panic(pkg.NewError(fiber.StatusUnauthorized, constants.UNVERIFIED_USER, nil))
	}

	return c.Next()
}

func VerifiedCantGo(c *fiber.Ctx) error {

	if c.Locals("is_verified") == true {
		panic(pkg.NewError(fiber.StatusUnauthorized, constants.VERIFIED_USER, nil))
	}

	return c.Next()
}
