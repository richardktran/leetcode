package main

func timeRequiredToBuy(tickets []int, k int) int {
	iTrack := k
	n := len(tickets)
	time := 0

	for true {
		if iTrack == 0 && tickets[0]-1 == 0 {
			return time + 1
		}
		if tickets[0] == 0 {
			tickets = tickets[1:]
			n--
		} else {
			time++
			tickets = append(tickets, tickets[0]-1)
			tickets = tickets[1:]
		}
		iTrack--
		if iTrack < 0 {
			iTrack = n - 1
		}
	}
	return time
}

func timeRequiredToBuyTest() {
	println(timeRequiredToBuy([]int{2, 3, 2}, 2))                      // 6
	println(timeRequiredToBuy([]int{5, 1, 1, 1}, 0))                   // 8
	println(timeRequiredToBuy([]int{84, 49, 5, 24, 70, 77, 87, 8}, 3)) // 154
}

// func main() {
// 	timeRequiredToBuyTest()
// }
