package main

import (
	"time"

	"github.com/shirou/gopsutil/v4/cpu"
	"github.com/shirou/gopsutil/v4/disk"
	"github.com/shirou/gopsutil/v4/host"
	"github.com/shirou/gopsutil/v4/mem"
	"github.com/shirou/gopsutil/v4/sensors"
)

func getTempC() float64 {
	sensorsTemperatures, err := sensors.SensorsTemperatures()

	if err != nil {
		panic(err)
	}

	return sensorsTemperatures[0].Temperature
}

func getCpuPct() float64 {
	cpu, err := cpu.Percent(time.Duration(0), false)

	if err != nil {
		panic(err)
	}

	return cpu[0]
}

func getDiskNames() []string {
	disks, err := disk.Partitions(false)

	if err != nil {
		panic(err)
	}

	var diskNames []string

	for _, partition := range disks {
		diskNames = append(diskNames, partition.Mountpoint)
	}

	return diskNames
}

func getDiskUsage(drivePath string) float64 {
	diskUsage, err := disk.Usage(drivePath)

	if err != nil {
		panic(err)
	}

	return diskUsage.UsedPercent
}

func getRamPct() float64 {
	virtualMemory, err := mem.VirtualMemory()

	if err != nil {
		panic(err)
	}

	return virtualMemory.UsedPercent
}

func getUptimeSec() uint64 {
	uptimeSec, err := host.Uptime()

	if err != nil {
		panic(err)
	}

	return uptimeSec
}
