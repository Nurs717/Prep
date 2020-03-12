package main

import "fmt"

func main() {
	// args := os.Args[1:]
	// str := args[0]
	// runestr := []rune(str)
	// for i := range runestr {
	// 	if (runestr[i] >= 97 && runestr[i] <= 122) || (runestr[i] >= 65 && runestr[i] <= 90) {
	// 		if runestr[i] >= 108 && runestr[i] <= 122 {
	// 			runestr[i] = runestr[i] - 13
	// 		} else if runestr[i] >= 76 && runestr[i] <= 90 {
	// 			runestr[i] = runestr[i] - 13
	// 		} else {
	// 			runestr[i] = runestr[i] + 13
	// 		}
	// 	}
	// }
	// for i := range runestr {
	// 	z01.PrintRune(runestr[i])
	// }
	// z01.PrintRune('\n')
	fmt.Println(split("HelloHAhowHAareHAyou?HAds", "HA"))
}

func index(str, charset string) int {
	for i := 0; i < len(str); i++ {
		for j := i + 1; j <= len(str); j++ {
			if str[i:j] == charset {
				return i
			}
		}
	}
	return -1
}

func split(str, charset string) []string {
	arr := []string{}
	for index(str, charset) != -1 {
		arr = append(arr, str[:index(str, charset)])
		str = str[index(str, charset)+len(charset):]
	}
	if len(str) != 0 {
		arr = append(arr, str)
	}
	return arr
}
