package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func celsiusToFahrenheit(c float64) float64 {
	return c*9/5 + 32
}

func fahrenheitToCelsius(f float64) float64 {
	return (f - 32) * 5 / 9
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Temperature Converter")
	fmt.Println("1. Celsius to Fahrenheit")
	fmt.Println("2. Fahrenheit to Celsius")
	fmt.Print("Choose an option (1 or 2): ")
	optionInput, _ := reader.ReadString('\n')
	option := strings.TrimSpace(optionInput)

	fmt.Print("Enter the temperature: ")
	tempInput, _ := reader.ReadString('\n')
	tempInput = strings.TrimSpace(tempInput)
	temp, err := strconv.ParseFloat(tempInput, 64)
	if err != nil {
		fmt.Println("Invalid temperature input.")
		return
	}

	switch option {
	case "1":
		result := celsiusToFahrenheit(temp)
		fmt.Printf("%.2f°C = %.2f°F\n", temp, result)
	case "2":
		result := fahrenheitToCelsius(temp)
		fmt.Printf("%.2f°F = %.2f°C\n", temp, result)
	default:
		fmt.Println("Invalid option.")
	}
}
