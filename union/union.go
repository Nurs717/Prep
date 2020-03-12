package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	if len(args) != 2 {
		return
	}
	first := []rune(args[0])
	second := []rune(args[1])
	str1 := ""
	str2 := ""
	for _, val1 := range first {
		if isUnique(val1, []rune(str1)) {
			str1 += string(val1)
		}
	}
	for _, val2 := range second {
		if isUnique(val2, []rune(str2)) {
			str2 += string(val2)
		}
	}
	str := str1 + str2
	res := ""
	for _, v := range str {
		if isUnique(v, []rune(res)) {
			res += string(v)
		}
	}
	for _, v := range res {
		z01.PrintRune(v)
	}
	z01.PrintRune('\n')
}
func isUnique(a rune, arr []rune) bool {
	for _, v := range arr {
		if v == a {
			return false
		}
	}
	return true
}
