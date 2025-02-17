package main

func sendWarningEmail(metric string, units string, threshhold int, currentValue float64) {
	Logger.Println("Sending email: ")
	Logger.Printf("Warning: Metric '%s' has exceeded threshold of %v%s. Current value: %f%s.", metric, threshhold, units, currentValue, units)
}
