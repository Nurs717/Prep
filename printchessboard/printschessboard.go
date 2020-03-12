package main

import (
	"os"
	"strconv"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]

	if len(args) > 2 || len(args) < 2 {
		displayStr("Error")
	} else {
		col, _ := strconv.Atoi(args[0])
		row, _ := strconv.Atoi(args[1])
		if row > 0 && col > 0 {
			k := 1
			for i := 1; i <= col; i++ {
				for j := 1; j <= row; j++ {
					if k == 1 {
						z01.PrintRune('#')
					} else {
						z01.PrintRune(' ')
					}
					k *= -1
				}
				if row%2 == 0 {
					k *= -1
				}
				z01.PrintRune('\n')
			}
		} else {
			displayStr("Error")
		}
	}
}
func displayStr(str string) {
	for _, s := range str {
		z01.PrintRune(s)
	}
	z01.PrintRune('\n')
}
