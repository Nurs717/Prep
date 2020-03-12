package main

import "fmt"

func main() {
	str := "Hello Alem"
	fmt.Println(StrLen(str))
}

func StrLen(str string) int {
	nb := 0
	for range str {
		nb++
	}
	return nb
}
