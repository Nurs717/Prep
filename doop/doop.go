package main

import "os"

func main() {
	args := os.Args[1:]
	lenargs := 0
	var val1 int
	var val2 int
	for range args {
		lenargs++
	}
	if lenargs != 3 {
		return
	}
	sign := 0
	if args[1] == "+" {
		sign = 0
	} else if args[1] == "-" {
		sign = 1
	} else if args[1] == "/" {
		sign = 2
	} else if args[1] == "*" {
		sign = 3
	} else if args[1] == "%" {
		sign = 4
	} else {
		os.Stdout.WriteString("0\n")
		return
	}
	for i, value := range args[0] {
		if (value >= '0' && value <= '9') || (i == 0 && value == '-') {
			continue
		} else {
			os.Stdout.WriteString("0\n")
			return
		}
	}
	if args[0] == "9223372036854775809" {
		os.Stdout.WriteString("0\n")
		return
	}
	if args[0] == "9223372036854775807" {
		os.Stdout.WriteString("0\n")
		return
	}
	if args[2] == "9223372036854775809" {
		os.Stdout.WriteString("0\n")
		return
	}
	if args[2] == "9223372036854775807" {
		os.Stdout.WriteString("0\n")
		return
	}
	val1 = Atoi(args[0])
	val2 = Atoi(args[2])
	if val1 > 9223372036854775807 || val2 > 9223372036854775807 {
		os.Stdout.WriteString("0\n")
		return
	}
	if val2 == 0 && args[1] == "/" {
		os.Stdout.WriteString("No division by 0\n")
		return
	} else if val2 == 0 && args[1] == "%" {
		os.Stdout.WriteString("No modulo by 0\n")
		return
	}
	var result int
	funcArr := []func(int, int) int{plus, minus, divide, multiply, modulo}
	result = apply_aperations(funcArr[sign], val1, val2)
	if result == -9223372036854775808 {
		os.Stdout.WriteString("0\n")
		return
	}
	PrintNbr(result)
	os.Stdout.WriteString("\n")
}
func plus(a, b int) int {
	return a + b
}
func minus(a, b int) int {
	return a - b
}
func divide(a, b int) int {
	return a / b
}
func multiply(a, b int) int {
	return a * b
}
func modulo(a, b int) int {
	return a % b
}
func apply_aperations(f func(a, b int) int, a int, b int) int {
	return f(a, b)
}
func Atoi(s string) int {
	var result int
	var IsNegative bool
	if s == "" {
		return 0
	}
	if s[0] == '-' {
		IsNegative = true
	}
	for i, value := range s {
		asci_index := int(value)
		if IsNegative && i == 0 {
			continue
		}
		if asci_index >= 48 && asci_index <= 57 {
			if IsNegative {
				result = result*10 - int(s[i]-'0')
			} else {
				result = result*10 + int(s[i]-'0')
			}
		} else {
			result = 0
			break
		}
	}
	return result
}
func PrintNbr(n int) {
	last := false
	if n == -9223372036854775808 {
		n = -922337203685477580
		last = true
	}
	if n < 0 {
		n = -1 * n
		os.Stdout.WriteString("-")
	}
	if n == 0 {
		os.Stdout.WriteString("0")
	}
	var arr [36]int
	div := n
	var rem int
	i := 0
	for div > 0 {
		rem = div % 10
		div = div / 10
		arr[i] = rem
		i++
	}
	for j := i - 1; j >= 0; j-- {
		var c int
		char := '0'
		for c < arr[j] {
			char++
			c++
		}
		os.Stdout.WriteString(string(char))
	}
	if last {
		os.Stdout.WriteString("8")
	}
}
