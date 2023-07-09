package controller

import (
	"docker_go_test/app/model"
	"docker_go_test/app/service"
	"fmt"
	"github.com/gofiber/fiber/v2"
)

type ApnicController struct {
	ApnicService service.ApnicService
}

func NewApnicController(apnicService *service.ApnicService) ApnicController {
	return ApnicController{
		ApnicService: *apnicService,
	}
}

func (controller ApnicController) Route(app *fiber.App) {
	app.Get("insert-data", controller.InsertData)
	app.Get("api/whois/:inetnum/:range", controller.Index)
}

func (controller ApnicController) InsertData(ctx *fiber.Ctx) error {
	go controller.ApnicService.InsertData()
	return ctx.Status(fiber.StatusOK).JSON(model.WebResponse{
		Code:    fiber.StatusOK,
		Status:  true,
		Message: "This process is run in background.",
	})
}

func (controller ApnicController) Index(ctx *fiber.Ctx) error {
	inetNum := ctx.Params("inetnum")
	inetRange := ctx.Params("range")
	param := fmt.Sprintf("%s/%s", inetNum, inetRange)
	response := controller.ApnicService.WhoisIp(param)
	return ctx.Status(fiber.StatusOK).JSON(model.WebResponse{
		Code:    fiber.StatusOK,
		Status:  true,
		Message: "OK",
		Data:    response,
	})
}
