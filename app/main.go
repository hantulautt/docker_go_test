package main

import (
	"docker_go_test/app/config"
	"docker_go_test/app/controller"
	"docker_go_test/app/entity"
	"docker_go_test/app/exception"
	"docker_go_test/app/helper"
	"docker_go_test/app/repository"
	"docker_go_test/app/service"
	"github.com/go-co-op/gocron"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"time"
)

func main() {
	db := config.NewMysqlDatabase()

	/**
	Register repository
	*/
	apnicRepository := repository.NewApnicRepository(db)

	/**
	Register service
	*/
	apnicService := service.NewApnicService(&apnicRepository)

	/**
	Register controller
	*/
	apnicController := controller.NewApnicController(&apnicService)

	/**
	Cron insert data to db every day at 00:00:00
	Please adjust the time as your need
	*/
	cron := gocron.NewScheduler(time.Local)
	_, err := cron.Every(1).Day().At("22:26:00").Do(apnicService.InsertData)
	if err != nil {
		helper.WriteLog("cron.log", "ERROR "+err.Error())
	}
	cron.StartAsync()

	/**
	Create simple table using migration
	*/
	if err := db.Set("gorm:table_options", "COLLATE=utf8_general_ci").Migrator().CreateTable(&entity.ApnicInetnum{}); err != nil {
		helper.WriteLog("migration.log", "ERROR "+err.Error())
	}

	/**
	Register Fiber
	*/
	app := fiber.New(config.NewFiberConfig())
	app.Use(recover.New())

	apnicController.Route(app)

	/**
	Start App
	*/
	err = app.Listen(":8081")
	exception.PanicIfNeeded(err)
}
