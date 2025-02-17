package main

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type EnvConfig struct {
	email     string
	pass      string
	sleepTime int
}

var config EnvConfig

func loadConfig() {
	if err := godotenv.Load("config.env"); err != nil {
		Logger.Println("Failed to read config.env file")
		panic(err)
	}

	Logger.Println("Config file read")

	sleepTimeStr := os.Getenv("SLEEP_TIME_SEC")
	sleepTimeSec, err := strconv.Atoi(sleepTimeStr)

	if err != nil {
		Logger.Println("Failed to parse SLEEP_TIME_SEC")
		panic(err)
	}

	config = EnvConfig{
		email:     os.Getenv("EMAIL"),
		pass:      os.Getenv("PASS"),
		sleepTime: sleepTimeSec,
	}
}

func getConfig() EnvConfig {
	return config
}
