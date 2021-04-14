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
	p models.Product
}

func (controller *ProductController) Route(route fiber.Router) {
	route.Post("/product/create", controller.Insert)
	route.Get("/product/all", controller.List)

}

func (controller *ProductController) Insert(c *fiber.Ctx) error {
	// var p models.Product
	p := controller.p
	if err := c.BodyParser(&p); err != nil {
		handler.PanicIfNeeded(err)
	}

	p.Validate(p)

	response := p.Create(p)

	return c.JSON(handler.Response{
		Code:    200,
		Message: "OK",
		Data:    response,
	})
}

func (controller *ProductController) List(c *fiber.Ctx) error {
	p := controller.p

	response, err := p.List()
	handler.PanicIfNeeded(err)

	return c.JSON(handler.Response{
		Code:    200,
		Message: "OK",
		Data:    response,
	})

}
