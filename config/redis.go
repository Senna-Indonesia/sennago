/*
 * SennaGo
 * Copyright (c) 2021.
 * author Muhammad Farid H
 */

package config

import (
	"sennago/systems/handler"
	"sync"

	"github.com/go-redis/redis"
)

type Redis struct {
	*redis.Client
}

var onceRedis sync.Once
var rdDb *Redis

func RedisDb() *Redis {
	onceRedis.Do(func() {
		configuration := New()
		client := redis.NewClient(&redis.Options{
			Addr:     configuration.Get("REDIS_ADDR"),
			Password: configuration.Get("REDIS_PWD"),
			DB:       0,
		})
		rdDb = &Redis{client}

		_, err := rdDb.Ping().Result()

		handler.PanicIfNeeded(err)

	})

	return rdDb
}
