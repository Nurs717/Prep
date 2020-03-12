package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) == 2 {
		str := os.Args[1]
		start := 0
		end := 0
		for i := len(str) - 1; i != 0; i-- {
			if str[i] != ' ' && end == 0 {
				end = i + 1
			}
			if str[i] == ' ' && end != 0 {
				start = i + 1
				break
			}
		}
		for _, i := range str[start:end] {
			z01.PrintRune(i)
		}
	}
	z01.PrintRune(10)
}
