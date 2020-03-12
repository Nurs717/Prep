package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args[1:]) == 1 {
		arg := os.Args[1]
		argrune := []rune(arg)
		for _, l := range argrune {
			if l >= 'A' && l <= 'Z' {
				z01.PrintRune('Z' - (l - 'A'))
			} else if l >= 'a' && l <= 'z' {
				z01.PrintRune('z' - (l - 'a'))
			} else {
				z01.PrintRune(l)
			}
		}

	}
	z01.PrintRune('\n')
}
