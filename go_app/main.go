// Made by Mohamad Tanbari
// Made on 2023 mar 24
//
// This script calculates salary post tax

package main

import (
	"fmt"
	"github.com/leekchan/accounting"
)

func main() {
	// Declare variables
	var wage float64
	var TAX_RATE float64 = 0.18 // 18% tax
	var hoursWorked float64

	// Get user input
	fmt.Print("Enter your hourly wage: ")
	fmt.Scan(&wage)
	fmt.Print("Enter hours worked: ")
	fmt.Scan(&hoursWorked)

	// Calculate salary
	var grossPay = wage * hoursWorked
	var netPay = grossPay - (grossPay * TAX_RATE)

	// format currency and print results
	accountingFormater := accounting.Accounting{Symbol: "$", Precision: 2}
	fmt.Println("Your net pay is: ", accountingFormater.FormatMoney(netPay))

	fmt.Println("\nDone.")
}
