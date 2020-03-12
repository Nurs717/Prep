package main

import "github.com/01-edu/z01"

func main() {
	z01.PrintRune(FirstRune("Hello Alem"))
	z01.PrintRune('\n')
}

func FirstRune(s string) rune {
	var rrune rune
	for _, r := range s {
		rrune = r
		break
	}
	return rrune
}
