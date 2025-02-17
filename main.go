package main

import (
	"log"
	"os"
	"time"
)

var logFlags int = log.LstdFlags | log.Lmicroseconds
var Logger *log.Logger = log.New(os.Stdout, "", logFlags)

func main() {
	Logger.Println("Script started")

	initialize()

	for {
		tempC := getTempC()
		//Logger.Println("Temp: " + strconv.FormatFloat(tempC, 'f', 2, 64) + "ºC")

		if tempC > float64(getConfig().threshholdTemp) {
			sendWarningEmail("Temperature", "ºC", getConfig().threshholdTemp, tempC)
		}

		cpuPct := getCpuPct()
		//Logger.Println("CPU: " + strconv.FormatFloat(cpuPct, 'f', 2, 64) + "%")

		var disk1Usage float64
		var disk2Usage float64

		for index, diskName := range getDiskNames() {
			diskUsage := getDiskUsage(diskName)

			if index == 0 {
				disk1Usage = diskUsage
			} else if index == 1 {
				disk2Usage = diskUsage
			}

			if disk1Usage > float64(getConfig().threshholdDisk1) {
				sendWarningEmail("Disk 1 Usage", "%", getConfig().threshholdDisk1, disk1Usage)
			} else if disk2Usage > float64(getConfig().threshholdDisk2) {
				sendWarningEmail("Disk 2 Usage", "%", getConfig().threshholdDisk2, disk2Usage)
			}

			//Logger.Println(diskName + "/ usage: " + strconv.FormatFloat(diskUsage, 'f', 2, 64) + "%")
		}

		memoryUsage := getRamPct()
		//Logger.Println("Ram: " + strconv.FormatFloat(memoryUsage, 'f', 2, 64) + "%")

		uptimeSec := getUptimeSec()
		//uptimeHours := float64(uptimeSec) / 60 / 60
		//Logger.Println("Uptime: " + strconv.FormatFloat(uptimeHours, 'f', 2, 64) + "hrs")

		currentTime := time.Now()

		// Save to SQLite file
		insertMetrics(currentTime, tempC, cpuPct, disk1Usage, disk2Usage, memoryUsage, uptimeSec)

		Logger.Println("Running at " + currentTime.String())

		time.Sleep(time.Duration(getConfig().sleepTime) * time.Second)
	}
}

func initialize() {
	loadConfig()
	createDatabaseIdempotent()
}
