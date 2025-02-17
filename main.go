package main

import (
	"log"
	"os"
	"strconv"
	"time"
)

var logFlags int = log.LstdFlags | log.Lmicroseconds
var Logger *log.Logger = log.New(os.Stdout, "", logFlags)

func main() {
	Logger.Println("Script started")
	createDatabaseIdempotent()

	for {
		tempC := getTempC()
		Logger.Println("Temp: " + strconv.FormatFloat(tempC, 'f', 2, 64) + "ºC")

		cpuPct := getCpuPct()
		Logger.Println("CPU: " + strconv.FormatFloat(cpuPct, 'f', 2, 64) + "%")

		var disk1Usage float64
		var disk2Usage float64

		for index, diskName := range getDiskNames() {
			diskUsage := getDiskUsage(diskName)

			if index == 0 {
				disk1Usage = diskUsage
			} else if index == 1 {
				disk2Usage = diskUsage
			}

			Logger.Println(diskName + "/ full: " + strconv.FormatFloat(diskUsage, 'f', 2, 64) + "%")
		}

		memoryUsage := getRamPct()
		Logger.Println("Ram: " + strconv.FormatFloat(memoryUsage, 'f', 2, 64) + "%")

		uptimeSec := getUptimeSec()
		uptimeHours := float64(uptimeSec) / 60 / 60
		Logger.Println("Uptime: " + strconv.FormatFloat(uptimeHours, 'f', 2, 64) + "hrs")

		currentTime := time.Now()

		// Save to SQLite file
		insertMetrics(currentTime, tempC, cpuPct, disk1Usage, disk2Usage, memoryUsage, uptimeSec)

		Logger.Println("Running at " + currentTime.String())
		time.Sleep(60 * time.Second)
	}
}
