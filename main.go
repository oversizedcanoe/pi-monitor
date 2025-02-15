package main

import (
	"fmt"
	"time"
)

func main() {
	for {
		tempC := getTempC()
		fmt.Println(tempC)
		currentTime := time.Now()
		fmt.Println("Running at " + currentTime.String())
		time.Sleep(3 * time.Second)
	}
}
