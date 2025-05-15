package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func countWords(text string) int {
	words := strings.Fields(text)
	return len(words)
}

func countCharacters(text string) int {
	return len([]rune(text))
}

func countSpecificLetter(text, letter string) int {
	count := 0
	for _, ch := range text {
		if strings.EqualFold(string(ch), letter) {
			count++
		}
	}
	return count
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	fmt.Println("Choose analysis:")
	fmt.Println("1. Count words")
	fmt.Println("2. Count characters")
	fmt.Println("3. Count specific letter")
	fmt.Print("Enter option (1, 2, or 3): ")
	option, _ := reader.ReadString('\n')
	option = strings.TrimSpace(option)

	switch option {
	case "1":
		fmt.Printf("Word count: %d\n", countWords(input))
	case "2":
		fmt.Printf("Character count: %d\n", countCharacters(input))
	case "3":
		fmt.Print("Enter the letter to count: ")
		letter, _ := reader.ReadString('\n')
		letter = strings.TrimSpace(letter)
		if len([]rune(letter)) != 1 {
			fmt.Println("Please enter a single letter.")
			return
		}
		fmt.Printf("Occurrences of '%s': %d\n", letter, countSpecificLetter(input, letter))
	default:
		fmt.Println("Invalid option.")
	}
}
