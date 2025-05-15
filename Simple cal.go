package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the first number:")
	num1Input, _ := reader.ReadString('\n')
	num1Input = strings.TrimSpace(num1Input)
	num1, _ := strconv.ParseFloat(num1Input, 64)

	fmt.Println("Enter the second number:")
	num2Input, _ := reader.ReadString('\n')
	num2Input = strings.TrimSpace(num2Input)
	num2, _ := strconv.ParseFloat(num2Input, 64)

	fmt.Print("Enter operation (+, -, *, /): ")
	opInput, _ := reader.ReadString('\n')
	op := strings.TrimSpace(opInput)

	var result float64
	switch op {
	case "+":
		result = num1 + num2
	case "-":
		result = num1 - num2
	case "*":
		result = num1 * num2
	case "/":
		if num2 != 0 {
			result = num1 / num2
		} else {
			fmt.Println("Error: Division by zero is not allowed.")
			return
		}
	default:
		fmt.Println("Error: Invalid operator.")
		return
	}

	fmt.Printf("Result: %.2f\n", result)

}
