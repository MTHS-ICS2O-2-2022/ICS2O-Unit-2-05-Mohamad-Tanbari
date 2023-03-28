// Made by Mohamad Tanbari
// Made on 2023 mar 24
//
// This script calculates salary post tax

package main

import (
	"fmt"
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

 // Round the results to 2 decimal points
	roundedNetPay := fmt.Sprintf("%.2f", netPay)
	roundedTAX := fmt.Sprintf("%.2f", grossPay-netPay)

// Output the results
	fmt.Println("\nYour net pay is:", ("$" + roundedNetPay))
	fmt.Println("The government took:", ("$" + roundedTAX), "from your pay.")

	fmt.Println("\nDone.")
}
