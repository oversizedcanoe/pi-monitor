package main

import (
	"strconv"
	"strings"
)

func getTempC(isWindows bool) float64 {
	var tempC float64

	if isWindows {
		// Must be run as admin?
		processResult := runProcess("cmd", "/C", "wmic", "/namespace:\\\\root\\wmi", "PATH", "MSAcpi_ThermalZoneTemperature", "get", "CurrentTemperature")

		// Output is four lines, 2nd line has temp
		tempText := strings.Split(processResult, "\n")[1]
		tempText = strings.TrimSpace(tempText)

		// Result is in tenth degrees Kelvin
		var tempTenthKelvin int
		var err error

		tempTenthKelvin, err = strconv.Atoi(tempText)

		if err != nil {
			panic(err)
		}

		tempC = (float64(tempTenthKelvin) / 10) - 273
	} else {
		// Linux
		// vcgencmd measure_temp
		// Output:
		// temp=40.4'C

	}

	return tempC
}

func getCpuPct(isWindows bool) int {
	var cpuPct int

	if isWindows {
		processResult := runProcess("cmd", "/C", "wmic", "cpu", "get", "loadpercentage")

		// Output is three lines, 2nd line has percent
		tempText := strings.Split(processResult, "\n")[1]
		tempText = strings.TrimSpace(tempText)

		var err error

		cpuPct, err = strconv.Atoi(tempText)

		if err != nil {
			panic(err)
		}
	} else {

	}

	return cpuPct
}

func getGpuPct() int {
	return 1
}
