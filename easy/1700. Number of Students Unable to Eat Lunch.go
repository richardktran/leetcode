package main

import "fmt"

func countStudents(students []int, sandwiches []int) int {
	nums0, nums1 := 0, 0

	for _, student := range students {
		if student == 0 {
			nums0++
		} else {
			nums1++
		}
	}

	for true {
		if nums0 == 0 && nums1 == 0 {
			return 0
		}

		if nums0 == 0 && sandwiches[0] == 0 {
			return len(students)
		}

		if nums1 == 0 && sandwiches[0] == 1 {
			return len(students)
		}

		if students[0] == sandwiches[0] {
			if students[0] == 0 {
				nums0--
			} else {
				nums1--
			}

			students = students[1:]
			sandwiches = sandwiches[1:]
		} else {
			students = append(students, students[0])
			students = students[1:]
		}

	}

	return 0
}

func countStudentsTest() {
	fmt.Println(countStudents([]int{1, 1, 0, 0}, []int{0, 1, 0, 1}))                                                                   // 0
	fmt.Println(countStudents([]int{1, 1, 1, 0, 0, 1}, []int{1, 0, 0, 0, 1, 1}))                                                       // 3
	fmt.Println(countStudents([]int{0, 1, 0, 1, 0, 1, 1, 1, 0, 0, 0, 0, 1, 1, 1}, []int{0, 1, 0, 1, 1, 0, 0, 1, 1, 1, 0, 1, 0, 1, 0})) // 0
}

// func main() {
// 	countStudentsTest()
// }
