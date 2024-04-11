package main

func removeKdigits(num string, k int) string {
	stack := []int{}
	for _, v := range num {
		if k == 0 {
			stack = append(stack, int(v))
			continue
		}
		for k > 0 && len(stack) > 0 && stack[len(stack)-1] > int(v) {
			stack = stack[:len(stack)-1]
			k--
		}
		stack = append(stack, int(v))
	}

	stack = stack[:len(stack)-k]

	var result []byte

	for _, v := range stack {
		if len(result) == 0 && v == int(byte('0')) {
			continue
		}
		result = append(result, byte(v))
	}

	if len(result) == 0 {
		return "0"
	}

	return string(result)
}

func removeKdigitsTest() {
	num := "1432219"
	k := 3
	println(removeKdigits(num, k)) // 1219

	num = "10200"
	k = 1
	println(removeKdigits(num, k)) // 200

	num = "10"
	k = 2
	println(removeKdigits(num, k)) // 0
}

// func main() {
// 	removeKdigitsTest()
// }
