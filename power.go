package main

import "fmt"

func main() {
	var nb int64 = 2
	var power int64 = 64
	var result int64 = 1
	for i := int64(1); i <= power; i++ {
		if i == 63 {
			result = (result * nb) - 1
			fmt.Println(result)
			return
		}
		result = result * nb
	}
	fmt.Println(result)
}
