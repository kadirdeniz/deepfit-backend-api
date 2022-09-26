package handler

import (
	"deepfit/pkg"
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func ErrorHandler(ctx *fiber.Ctx, err error) error {

	errObject, ok := err.(*pkg.Error)
	if !ok {
		fmt.Println("Error is not of type *pkg.Error")
		return ctx.Status(500).JSON(
			pkg.NewResponse(false, "Internal Server Error", nil),
		)
	}

	fmt.Println(errObject.ErrorLog)

	return ctx.Status(errObject.StatusCode).JSON(pkg.Response{
		Status:  false,
		Message: errObject.ErrorMessage,
		Data:    nil,
	})
}
