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
	taxRate := 0.18 // 18% tax
	var wage float64
	var hoursWorked float64

	// Get user input
	fmt.Print("Enter your hourly wage: ")
	fmt.Scan(&wage)
	fmt.Print("Enter hours worked: ")
	fmt.Scan(&hoursWorked)

	// Calculate salary
	var grossPay = wage * hoursWorked
	var netPay = grossPay - (grossPay * taxRate)
	var tax = grossPay - netPay

	// Round the results
	roundedNetPay := fmt.Sprintf("%.2f", netPay)
	roundedTax := fmt.Sprintf("%.2f", tax)

	// Output the results

	fmt.Println("\nYour net pay is:", ("$" + roundedNetPay))
	fmt.Println("The government took:", ("$" + roundedTax), "from your pay.")

	fmt.Println("\nDone.")
}
