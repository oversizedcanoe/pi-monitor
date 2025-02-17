package main

import (
	"fmt"
	"strconv"
	"time"
)

func createDatabaseIdempotent() {
	query := "SELECT name FROM sqlite_master WHERE type='table' AND name='Metrics';"

	result := queryRow(query)
	var name string
	result.Scan(&name)

	if name == "" {
		fmt.Println("Creating table")

		command := `CREATE TABLE "Metrics" (
		"QueryTime"	INTEGER UNIQUE,
		"TempC"	REAL,
		"CpuPct"	REAL,
		"Disk1Usage"	REAL,
		"Disk2Usage"	REAL,
		"MemoryUsage"	REAL,
		"UptimeSec"	INTEGER
		);`

		executeCommand(command)
	}
}

func insertMetrics(time time.Time, tempC float64, cpuPct float64,
	disk1Usage float64, disk2Usage float64, memoryUsage float64, uptimeSec uint64) {
	command := "INSERT INTO Metrics (QueryTime, TempC, CpuPct, Disk1Usage, Disk2Usage, MemoryUsage, UptimeSec) VALUES ($1, $2, $3, $4, $5, $6, $7)"

	// Convert time to seconds since Epoch so it can be stored as int in DB.
	// To convert back to time.Time, do time.Unix(value, 0)
	affectedRows := executeCommand(command, time.Unix(), tempC, cpuPct, disk1Usage, disk2Usage, memoryUsage, uptimeSec)

	fmt.Println(strconv.Itoa(affectedRows) + " rows affected")
}
