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

func (p *Product) Validate(request Product) {
	if err := validation.ValidateStruct(&request,
		validation.Field(&request.Name, validation.Required),
		validation.Field(&request.Stock, validation.Required, validation.Min(1)),
	); err != nil {
		panic(handler.ValidationError{
			Message: err.Error(),
		})
	}
}

func (p *Product) Create(request Product) error {
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

func (p *Product) List() ([]Product, error) {
	var data []Product
	err := config.Db().Model(&data).Select()
	if err != nil {
		handler.PanicIfNeeded(err)
	}
	return data, nil

}
