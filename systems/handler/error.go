/*
 * SennaGo
 * Copyright (c) 2021.
 * author Muhammad Farid H
 */

package handler

func PanicIfNeeded(err interface{}) {
	if err != nil {
		panic(err)
	}
}
