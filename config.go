package main

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type EnvConfig struct {
	email           string
	pass            string
	sleepTime       int
	threshholdTemp  int
	threshholdDisk1 int
	threshholdDisk2 int
}

var config EnvConfig

func getConfig() EnvConfig {
	return config
}

func loadConfig() {
	if err := godotenv.Load("config.env"); err != nil {
		Logger.Println("Failed to read config.env file")
		panic(err)
	}

	Logger.Println("Config file read")

	config = EnvConfig{
		email:           os.Getenv("EMAIL"),
		pass:            os.Getenv("PASS"),
		sleepTime:       parseConfigToInt("SLEEP_TIME_SEC"),
		threshholdTemp:  parseConfigToInt("THRESHHOLD_TEMP"),
		threshholdDisk1: parseConfigToInt("THRESHHOLD_DISK1_USAGE"),
		threshholdDisk2: parseConfigToInt("THRESHHOLD_DISK2_USAGE"),
	}
}

func parseConfigToInt(key string) int {
	envValueStr := os.Getenv(key)
	envValue, err := strconv.Atoi(envValueStr)

	if err != nil {
		Logger.Println("Failed to parse " + key)
		panic(err)
	}

	return envValue
}
