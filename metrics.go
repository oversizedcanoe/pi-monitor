package main

import (
	"strconv"
	"strings"
)

func getTempC() float32 {
	processResult := runProcess("cmd", "/C", "wmic", "/namespace:\\\\root\\wmi", "PATH", "MSAcpi_ThermalZoneTemperature", "get", "CurrentTemperature")

	tempText := strings.Split(processResult, "\n")[1]
	tempText = strings.TrimSpace(tempText)

	var tempTenthKelvin int
	var err error

	tempTenthKelvin, err = strconv.Atoi(tempText)

	if err != nil {
		panic(err)
	}

	tempC := (float32(tempTenthKelvin) / 10) - 273

	// Windows -- Must be run as admin :(
	// wmic /namespace:\\root\wmi PATH MSAcpi_ThermalZoneTemperature get CurrentTemperature
	// Output:
	// CurrentTemperature
	// 3072
	//
	//

	// Two lines (maybe four), in tenth degrees Kelvin
	// C = (value / 10) - 273

	// Linux
	// vcgencmd measure_temp
	// Output:
	// temp=40.4'C

	return tempC
}

func getCpuPct() int {
	return 1
}

func getGpuPct() int {
	return 1
}
