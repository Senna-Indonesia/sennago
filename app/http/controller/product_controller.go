/*
 * SennaGo
 * Copyright (c) 2021.
 * author Muhammad Farid H
 */
package controller

import (
	"sennago/app/models"
	"sennago/systems/handler"

	"github.com/gofiber/fiber/v2"
)

type ProductController struct {
}

func (controller *ProductController) Route(route fiber.Router) {
	route.Post("/products/create", controller.Insert)
}

func (controller *ProductController) Insert(c *fiber.Ctx) error {
	var p models.Product
	err := c.BodyParser(&p)

	handler.PanicIfNeeded(err)

	p.Validate(p)
	response := p.Create(p)

	return c.JSON(handler.Response{
		Code:    200,
		Message: "OK",
		Data:    response,
	})
}
