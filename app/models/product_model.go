/*
 * SennaGo
 * Copyright (c) 2021.
 * author Muhammad Farid H
 */

package models

import (
	"sennago/config"
	"sennago/systems/handler"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type Product struct {
	Id    int    `json:"id" `
	Name  string `json:"name"`
	Stock int    `json:"stock"`
}

type ProductRequest struct {
	Name  string `json:"name" form:"name"`
	Stock int    `json:"stock" form:"stock"`
}

type ProdutModel struct {
}

func (p *ProdutModel) Validate(request ProductRequest) {
	err := validation.ValidateStruct(&request,
		validation.Field(&request.Name, validation.Required),
		validation.Field(&request.Stock, validation.Required, validation.Min(1)),
	)

	if err != nil {
		panic(handler.ValidationError{
			Message: err.Error(),
		})
	}
}

func (p *ProdutModel) Create(request ProductRequest) error {
	_, err := config.Db().Model(&Product{
		Name:  request.Name,
		Stock: request.Stock,
	}).Insert()
	if err != nil {
		handler.PanicIfNeeded(err)
	}
	return nil
}
func Update() {

}
func Delete() {

}

func GetOne() {

}

func List() {

}
