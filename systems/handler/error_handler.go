/*
 * SennaGo
 * Copyright (c) 2021.
 * author Muhammad Farid H
 */

package handler

import (
	"github.com/gofiber/fiber/v2"
)

func ErrorHandler(ctx *fiber.Ctx, err error) error {

	_, ok := err.(ValidationError)
	if ok {
		return ctx.JSON(Response{
			Code:    400,
			Message: "BAD_REQUEST",
			Data:    err.Error(),
		})
	}

	return ctx.JSON(Response{
		Code:    500,
		Message: "INTERNAL_SERVER_ERROR",
		Data:    err.Error(),
	})
}
