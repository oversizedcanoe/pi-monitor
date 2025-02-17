package main

import (
	"errors"
	"os"
	"strconv"
	"time"
)

func createDatabaseIdempotent() {
	// Ensure file exists
	if _, err := os.Stat("./database.db"); errors.Is(err, os.ErrNotExist) {
		Logger.Println("Creating database file")
		if _, err := os.Create("./database.db"); err != nil {
			panic(err)
		}
	} else {
		Logger.Println("Database file exists")
	}

	query := "SELECT name FROM sqlite_master WHERE type='table' AND name='Metrics';"

	result := queryRow(query)
	var name string
	result.Scan(&name)

	if name == "" {
		Logger.Println("Creating table")

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
	} else {
		Logger.Println("'Metrics' table exists")
	}
}

func insertMetrics(time time.Time, tempC float64, cpuPct float64,
	disk1Usage float64, disk2Usage float64, memoryUsage float64, uptimeSec uint64) {
	command := "INSERT INTO Metrics (QueryTime, TempC, CpuPct, Disk1Usage, Disk2Usage, MemoryUsage, UptimeSec) VALUES ($1, $2, $3, $4, $5, $6, $7)"

	// Convert time to seconds since Epoch so it can be stored as int in DB.
	// To convert back to time.Time, do time.Unix(value, 0)
	affectedRows := executeCommand(command, time.Unix(), tempC, cpuPct, disk1Usage, disk2Usage, memoryUsage, uptimeSec)

	Logger.Println(strconv.Itoa(affectedRows) + " rows affected")
}
