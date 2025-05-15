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

	fmt.Print("Enter the length of the rectangle: ")
	lengthInput, _ := reader.ReadString('\n')
	lengthInput = strings.TrimSpace(lengthInput)
	length, _ := strconv.ParseFloat(lengthInput, 64)

	fmt.Print("Enter the width of the rectangle: ")
	widthInput, _ := reader.ReadString('\n')
	widthInput = strings.TrimSpace(widthInput)
	width, _ := strconv.ParseFloat(widthInput, 64)

	area := length * width
	fmt.Printf("The area of the rectangle is: %.2f\n", area)
}
