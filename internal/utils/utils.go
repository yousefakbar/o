package utils

import (
	"fmt"
)

// Contains: Checks if a `slice` contains a specific `item`, regardless of type
func Contains[T comparable](slice []T, item T) bool {
	for _, val := range slice {
		if val == item {
			return true
		}
	}
	return false
}

func Prompt(prompt string) (string, error) {
	fmt.Printf("%s", prompt)

	var input string
	_, err := fmt.Scanln(&input)
	if err != nil {
		return "", fmt.Errorf("Failed to read note name: %w", err)
	}

	return input, nil
}
