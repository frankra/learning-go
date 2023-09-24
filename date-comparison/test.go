package main

import (
	"fmt"
	"time"
)

func main() {
	// Date in Singapore (UTC+8)
	nowInVendorTime := time.Now()
	loc, _ := time.LoadLocation("Europe/Lisbon") // UTC+1
	nowInVendorTime = nowInVendorTime.In(loc)

	dateInUTC := nowInVendorTime.
		In(time.UTC).
		Add(time.Minute * 10) // Add + 10 minutes so that the date is after

	// compare dates
	isAfter := nowInVendorTime.After(dateInUTC)
	fmt.Printf("Dates: %s is after %s: %v", nowInVendorTime, dateInUTC, isAfter)
}
