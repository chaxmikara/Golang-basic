package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter the principal amount: ")
	principalInput, _ := reader.ReadString('\n')
	principalInput = strings.TrimSpace(principalInput)
	principal, err := strconv.ParseFloat(principalInput, 64)
	if err != nil || principal <= 0 {
		fmt.Println("Invalid principal amount.")
		return
	}

	fmt.Print("Enter the annual interest rate (in percent): ")
	rateInput, _ := reader.ReadString('\n')
	rateInput = strings.TrimSpace(rateInput)
	annualRate, err := strconv.ParseFloat(rateInput, 64)
	if err != nil || annualRate < 0 {
		fmt.Println("Invalid interest rate.")
		return
	}
	monthlyRate := annualRate / 12 / 100

	fmt.Print("Enter the loan term (in months): ")
	termInput, _ := reader.ReadString('\n')
	termInput = strings.TrimSpace(termInput)
	months, err := strconv.Atoi(termInput)
	if err != nil || months <= 0 {
		fmt.Println("Invalid loan term.")
		return
	}

	var payment float64
	if monthlyRate == 0 {
		payment = principal / float64(months)
	} else {
		payment = principal * (monthlyRate * math.Pow(1+monthlyRate, float64(months))) / (math.Pow(1+monthlyRate, float64(months)) - 1)
	}

	fmt.Printf("Your monthly payment is: %.2f\n", payment)
}
