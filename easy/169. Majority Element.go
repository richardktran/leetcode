package main

func majorityElement(nums []int) int {
	majority := len(nums) / 2
	hashSet := map[int]int{}

	for _, v := range nums {
		if _, ok := hashSet[v]; !ok {
			hashSet[v] = 1
		} else {
			hashSet[v]++
		}
	}

	for key, element := range hashSet {
		if element > majority {
			return key
		}
	}

	return -1
}

func MajorityElementTest() {
	println(majorityElement([]int{3, 2, 3}))             // 3
	println(majorityElement([]int{2, 2, 1, 1, 1, 2, 2})) // 2
}

// func main() {
// 	MajorityElementTest()
// }
