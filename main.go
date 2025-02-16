package main

import (
	"fmt"
	"runtime"
	"strconv"
	"time"
)

func main() {
	for {
		isWindows := runtime.GOOS == "windows"
		tempC := getTempC(isWindows)
		fmt.Println("Temp: " + strconv.FormatFloat(tempC, 'f', 2, 64) + "C")
		cpuPct := getCpuPct(isWindows)
		fmt.Println("CPU: " + strconv.Itoa(cpuPct) + "%")

		percentUsedInC := getFullStoragePctInDrive(isWindows, "c")
		fmt.Println("C:/ full: " + strconv.FormatFloat(percentUsedInC, 'f', 2, 64) + "%")
		percentUsedInD := getFullStoragePctInDrive(isWindows, "d")
		fmt.Println("D:/ full: " + strconv.FormatFloat(percentUsedInD, 'f', 2, 64) + "%")

		currentTime := time.Now()
		fmt.Println("Running at " + currentTime.String())
		time.Sleep(3 * time.Second)
	}
}
