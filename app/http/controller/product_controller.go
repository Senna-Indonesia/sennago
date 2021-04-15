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
	route.Get("/product/one/:id", controller.One)
	route.Delete("/product/delete/:id", controller.Delete)
	route.Post("/product/update/:id", controller.Update)

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

func (controller *ProductController) One(c *fiber.Ctx) error {
	p := controller.p
	id, err := c.ParamsInt("id")
	handler.PanicIfNeeded(err)

	response, errRes := p.GetOne(id)
	handler.PanicIfNeeded(errRes)

	return c.JSON(handler.Response{
		Code:    200,
		Message: "OK",
		Data:    response,
	})

}

func (controller *ProductController) Delete(c *fiber.Ctx) error {
	p := controller.p
	id, err := c.ParamsInt("id")
	handler.PanicIfNeeded(err)

	response := p.Delete(id)

	return c.JSON(handler.Response{
		Code:    200,
		Message: "OK",
		Data:    response,
	})

}

func (controller *ProductController) Update(c *fiber.Ctx) error {
	p := controller.p
	id, err := c.ParamsInt("id")
	handler.PanicIfNeeded(err)

	if err := c.BodyParser(&p); err != nil {
		handler.PanicIfNeeded(err)
	}

	p.Validate(p)

	response := p.Update(id, p)

	return c.JSON(handler.Response{
		Code:    200,
		Message: "OK",
		Data:    response,
	})

}
