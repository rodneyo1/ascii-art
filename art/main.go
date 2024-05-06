package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"asciiart"
)

func main() {
	// Open the file
	file, err := os.Open("standard.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Read lines from the file
	scanner := bufio.NewScanner(file)
	lines := []string{}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	// Split lines into characters
	characters := [][]string{}
	for i := 0; i < len(lines); i += 9 {
		end := i + 9
		if end > len(lines) {
			end = len(lines)
		}
		characters = append(characters, lines[i:end])
	}

	// Check command-line arguments
	if len(os.Args) != 2 {
		fmt.Println("Expected usage: [program_name] [argument]")
		return
	}

	// Check for non-printable characters in the input string
	input := os.Args[1]
	for i := 0; i < len(input); i++ {
		if (input[i] < 32 || input[i] > 126) && input[i] != 10 {
			fmt.Println("Cannot print one or more characters")
			return
		}
	}

	// Format ("\n") in input string
	input = strings.Replace(input, "\\n", "\n", -1)
	input = strings.Replace(input, "\\t", "    ", -1)

	// Generate ASCII art
	asciiart.HandleLn(input, characters)
}
