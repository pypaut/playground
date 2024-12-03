package main

import (
	d02 "d02/internal"
	"fmt"
)

func main() {
	reports, err := d02.ParseReports("input")
	if err != nil {
		panic(err)
	}

	number := d02.NumberOfSafeReports(reports)
	fmt.Printf("Number of safe reports: %d\n", number)

	return
}
