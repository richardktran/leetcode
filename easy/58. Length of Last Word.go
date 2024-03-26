package main

func lengthOfLastWord(s string) int {
	result := 0
	for i, c := range s {
		if i == 0 || string(c) == " " {
			continue
		}

		if string(c) != " " && string(s[i-1]) == " " {
			result = 0
		} else {
			result++
		}
	}

	return result + 1
}

func LengthOfLastWordTest() {
	println(lengthOfLastWord("Hello World")) // 5
	println(lengthOfLastWord("a "))          // 1
	println(lengthOfLastWord("a"))           // 1
	println(lengthOfLastWord(" "))           // 0
}

// func main() {
// 	LengthOfLastWordTest()
// }
