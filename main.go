/*
 * SennaGo
 * Copyright (c) 2021.
 * author Muhammad Farid H
 */

package main

import (
	"sennago/config"
	"sennago/routes"
	"sennago/systems/handler"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/etag"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {

	//Initial Config
	configuration := config.New()

	//Initial App
	app := fiber.New(config.App())
	app.Use(recover.New())
	app.Use(cors.New())
	app.Use(etag.New())

	//Initial Route
	routes.Routes(app)

	// Start App
	err := app.Listen(configuration.Get("APP_PORT"))
	handler.PanicIfNeeded(err)
}
