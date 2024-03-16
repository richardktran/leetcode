package main

import "fmt"

func romanToInt(s string) int {
	romanMapper := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}
	total := 0
	for i, _ := range s {
		total += romanMapper[s[i:i+1]]
		if i != 0 {
			if romanMapper[s[i-1:i]] < romanMapper[s[i:i+1]] {
				total -= 2 * romanMapper[s[i-1:i]]
			}
		}
	}

	return total
}

func RomanToIntTest() {
	fmt.Println(romanToInt("III"))
	fmt.Println(romanToInt("IV"))
	fmt.Println(romanToInt("IX"))
	fmt.Println(romanToInt("LVIII"))
	fmt.Println(romanToInt("MCMXCIV"))
}

// func main() {
// 	RomanToIntTest()
// }
