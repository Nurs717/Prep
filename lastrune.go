package main

import "github.com/01-edu/z01"

func main() {
	z01.PrintRune(LastRune("Hello Alem"))
	z01.PrintRune('\n')
}

func LastRune(s string) rune {
	len := 0
	var rrune rune
	for range s {
		len++
	}
	for i, r := range s {
		if len-1 != i {

		} else {
			rrune = r
		}
	}
	return rrune
}
