package main

import (
	"os"
	"asciiart"
	"bufio"
	"fmt"
	"strings"
)

func main() {
	file, _ := os.Open("standard.txt")

	scanner := bufio.NewScanner(file)
	lines := []string{}

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	character := make([]string, 9)
	result := [][]string{}
	for i := 0; i < len(lines); i += 9 {
		end := i + 9
		if end < len(lines) {
			end = len(lines)
		}
		character = lines[i:end]
		result = append(result, character)
	}

	if len(os.Args) < 2 || len(os.Args) > 2  {
		fmt.Println("Expected usage: [program_name] [argument]")
		return
	}
	
	r := os.Args[1]

	//Error handling for non-printable characters
	for i := 0; i <= len(r)-1; i++ {
		if (r[i] < 32 || r[i] > 126 ) && r[i] != 10 {//Skip \n (10) if present
			fmt.Println("Cant print one or more characters")
			return
		}
	}

	// Format ("\n") in input string
	s := strings.Replace(r, "\\n", "\n", -1)
	s = strings.Replace(s, "\\t", "    ", -1)

	asciiart.HandleLn(s, result)
}
