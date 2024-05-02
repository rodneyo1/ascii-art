package asciiart

import (
	"fmt"
)

func Printer(s string, b [][]string) {
	for i := 1; i < 9; i++ {
		for _, char := range s {
			toPrint := char - 32
			fmt.Print(b[toPrint][i])
		}
		fmt.Println()
	}
}
