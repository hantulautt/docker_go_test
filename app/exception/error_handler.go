package exception

import (
	"docker_go_test/app/model"
	"github.com/gofiber/fiber/v2"
)

func ErrorHandler(ctx *fiber.Ctx, err error) error {
	_, ok := err.(ValidationError)
	if ok {
		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(model.WebResponse{
			Code:    fiber.StatusUnprocessableEntity,
			Status:  false,
			Message: err.Error(),
		})
	}

	return ctx.Status(fiber.StatusInternalServerError).JSON(model.WebResponse{
		Code:    500,
		Status:  false,
		Message: "INTERNAL_SERVER_ERROR",
		Data:    err.Error(),
	})
}
