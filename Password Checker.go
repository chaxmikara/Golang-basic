package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func checkPasswordStrength(password string) []string {
	var issues []string
	if len([]rune(password)) < 8 {
		issues = append(issues, "at least 8 characters")
	}
	hasUpper, hasLower, hasDigit := false, false, false
	for _, ch := range password {
		switch {
		case unicode.IsUpper(ch):
			hasUpper = true
		case unicode.IsLower(ch):
			hasLower = true
		case unicode.IsDigit(ch):
			hasDigit = true
		}
	}
	if !hasUpper {
		issues = append(issues, "an uppercase letter")
	}
	if !hasLower {
		issues = append(issues, "a lowercase letter")
	}
	if !hasDigit {
		issues = append(issues, "a number")
	}
	return issues
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter a password: ")
	password, _ := reader.ReadString('\n')
	password = strings.TrimSpace(password)

	issues := checkPasswordStrength(password)
	if len(issues) == 0 {
		fmt.Println("Password is strong.")
	} else {
		fmt.Println("Password is weak. It should contain:")
		for _, issue := range issues {
			fmt.Printf("- %s\n", issue)
		}
	}
}
