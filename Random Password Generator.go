package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

func getCharset(includeUpper, includeLower, includeDigits, includeSymbols bool) string {
	var charset strings.Builder
	if includeUpper {
		charset.WriteString("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	}
	if includeLower {
		charset.WriteString("abcdefghijklmnopqrstuvwxyz")
	}
	if includeDigits {
		charset.WriteString("0123456789")
	}
	if includeSymbols {
		charset.WriteString("!@#$%^&*()-_=+[]{}|;:,.<>?/~")
	}
	return charset.String()
}

func generatePassword(length int, charset string) string {
	rand.Seed(time.Now().UnixNano())
	password := make([]rune, length)
	for i := range password {
		password[i] = rune(charset[rand.Intn(len(charset))])
	}
	return string(password)
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter password length: ")
	var length int
	fmt.Scanf("%d\n", &length)
	if length <= 0 {
		fmt.Println("Invalid length.")
		return
	}

	fmt.Print("Include uppercase letters? (y/n): ")
	upper, _ := reader.ReadString('\n')
	fmt.Print("Include lowercase letters? (y/n): ")
	lower, _ := reader.ReadString('\n')
	fmt.Print("Include digits? (y/n): ")
	digits, _ := reader.ReadString('\n')
	fmt.Print("Include symbols? (y/n): ")
	symbols, _ := reader.ReadString('\n')

	includeUpper := strings.TrimSpace(strings.ToLower(upper)) == "y"
	includeLower := strings.TrimSpace(strings.ToLower(lower)) == "y"
	includeDigits := strings.TrimSpace(strings.ToLower(digits)) == "y"
	includeSymbols := strings.TrimSpace(strings.ToLower(symbols)) == "y"

	if !includeUpper && !includeLower && !includeDigits && !includeSymbols {
		fmt.Println("At least one character type must be selected.")
		return
	}

	charset := getCharset(includeUpper, includeLower, includeDigits, includeSymbols)
	password := generatePassword(length, charset)
	fmt.Printf("Generated password: %s\n", password)
}
