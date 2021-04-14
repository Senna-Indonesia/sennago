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
	var request models.ProductRequest
	var model models.ProdutModel
	err := c.BodyParser(&request)

	handler.PanicIfNeeded(err)

	model.Validate(request)
	response := model.Create(request)

	return c.JSON(handler.Response{
		Code:    200,
		Message: "OK",
		Data:    response,
	})
}
