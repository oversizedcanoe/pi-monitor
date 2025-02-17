package main

import (
	"fmt"
	"net/smtp"
	"strconv"
)

func sendWarningEmail(metric string, units string, threshhold int, currentValue float64) {
	Logger.Println("Sending email: ")
	Logger.Printf("Warning: Metric '%s' has exceeded threshold of %v%s. Current value: %f%s.", metric, threshhold, units, currentValue, units)

	auth := smtp.PlainAuth("", getConfig().email, getConfig().pass, getConfig().emailHost)
	to := []string{getConfig().email}

	msg := []byte("To: " + getConfig().email + "\r\n" +
		fmt.Sprintf("Subject: %s Warning", getConfig().deviceName) + "\r\n" +
		"\r\n" +
		fmt.Sprintf("Warning: Metric '%s' has exceeded threshold of %v%s. Current value: %f%s.", metric, threshhold, units, currentValue, units) + "\r\n")

	err := smtp.SendMail(getConfig().emailHost+":"+strconv.Itoa(getConfig().emailPort), auth, getConfig().email, to, msg)

	if err != nil {
		Logger.Println("Failed to send email")
		Logger.Println(err)
		panic(err)
	}
}
