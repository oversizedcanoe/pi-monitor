package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	for {
		tempC := getTempC()
		fmt.Println("Temp: " + strconv.FormatFloat(tempC, 'f', 2, 64) + "ºC")

		cpuPct := getCpuPct()
		fmt.Println("CPU: " + strconv.FormatFloat(cpuPct, 'f', 2, 64) + "%")

		for _, diskName := range getDiskNames() {
			diskUsage := getDiskUsage(diskName)
			fmt.Println(diskName + "/ full: " + strconv.FormatFloat(diskUsage, 'f', 2, 64) + "%")
		}

		memoryUsage := getRamPct()
		fmt.Println("Ram: " + strconv.FormatFloat(memoryUsage, 'f', 2, 64) + "%")

		uptimeSec := getUptimeSec()
		uptimeHours := float64(uptimeSec) / 60 / 60
		fmt.Println("Uptime: " + strconv.FormatFloat(uptimeHours, 'f', 2, 64) + "hrs")

		currentTime := time.Now()
		fmt.Println("Running at " + currentTime.String())
		time.Sleep(3 * time.Second)
	}
}
