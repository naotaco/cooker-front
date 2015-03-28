package main

import (
	"gopkg.in/redis.v2"
)

func getCookerRedisClient() *redis.Client {
	client := redis.NewTCPClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	return client
}

func getTargetTemp() float64 {
	client := getCookerRedisClient()
	value := client.Get("cooker_target_temperature")
	temp, err := value.Float64()
	if err != nil {
		return -1
	}
	return temp
}

func getCurrentTemp() float64 {
	client := getCookerRedisClient()
	value := client.Get("cooker_current_temperature")
	temp, err := value.Float64()
	if err != nil {
		return -1
	}
	return temp
}

func isHeating() string {
	client := getCookerRedisClient()
	return client.Get("cooker_is_heating").String()
}
