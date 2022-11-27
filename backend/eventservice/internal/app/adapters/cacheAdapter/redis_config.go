package cacheadapter

import "github.com/go-redis/redis/v8"

const Adress = "Localhost:6349"
const Password = ""
const DB = 0

func RedisOption() *redis.Options {
	return &redis.Options{
		Addr:     Adress,
		Password: "",
		DB:       DB,
	}
}
