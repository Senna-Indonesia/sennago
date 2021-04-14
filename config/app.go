/*
 * SennaGo
 * Copyright (c) 2021.
 * author Muhammad Farid H
 */

package config

import (
	"sennago/systems/handler"

	"github.com/gofiber/fiber/v2"
)

func App() fiber.Config {
	return fiber.Config{
		ErrorHandler: handler.ErrorHandler,
		//using etag from cache
	}
}
