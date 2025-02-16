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
		fmt.Println("Temp C: " + strconv.FormatFloat(tempC, 'f', 2, 32))
		cpuPct := getCpuPct(isWindows)
		fmt.Println("CPU %:" + strconv.Itoa(cpuPct))

		currentTime := time.Now()
		fmt.Println("Running at " + currentTime.String())
		time.Sleep(3 * time.Second)
	}
}
