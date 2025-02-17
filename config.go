package main

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type EnvConfig struct {
	deviceName      string
	email           string
	pass            string
	emailPort       int
	emailHost       string
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
		deviceName:      os.Getenv("DEVICE"),
		email:           os.Getenv("EMAIL_ADDRESS"),
		pass:            os.Getenv("EMAIL_PASSWORD"),
		emailPort:       parseConfigToInt("EMAIL_PORT"),
		emailHost:       os.Getenv("EMAIL_HOST"),
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
