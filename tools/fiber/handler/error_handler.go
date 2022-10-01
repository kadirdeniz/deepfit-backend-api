package handler

import (
	"deepfit/constants"
	"deepfit/pkg"
	"fmt"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/gofiber/fiber/v2"
)

func ErrorHandler(ctx *fiber.Ctx, err error) error {

	errObject, ok := err.(*pkg.Error)
	if ok {

		if errObject.ErrorInfo != nil {
			fmt.Println(errObject.ErrorInfo)
		}

		return ctx.
			Status(errObject.StatusCode).
			JSON(
				pkg.NewResponse(false, errObject.ErrorMessage, nil),
			)
	}

	validationErr, ok := err.(validation.InternalError)
	if ok {
		fmt.Println(validationErr)
		return ctx.
			Status(fiber.StatusInternalServerError).
			JSON(
				pkg.NewResponse(false, constants.VALIDATION_ERROR, validationErr.InternalError()),
			)
	}

	fmt.Println(err)
	return ctx.Status(fiber.StatusInternalServerError).JSON(
		pkg.NewResponse(false, "Internal Server Error", nil),
	)
}
