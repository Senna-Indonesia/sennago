/*
 * SennaGo
 * Copyright (c) 2021.
 * author Muhammad Farid H
 */

package handler

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
