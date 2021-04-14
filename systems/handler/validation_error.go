/*
 * SennaGo
 * Copyright (c) 2021.
 * author Muhammad Farid H
 */

package handler

type ValidationError struct {
	Message string
}

func (validationError ValidationError) Error() string {
	return validationError.Message
}
