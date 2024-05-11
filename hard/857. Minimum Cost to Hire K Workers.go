package main

import (
	"container/heap"
	"fmt"
	"sort"
)

// Max heap for quality
type maxHeap []int

func (h maxHeap) Len() int {
	return len(h)
}
func (h maxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h maxHeap) Less(i, j int) bool {
	return h[i] > h[j]
}
func (h *maxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *maxHeap) Pop() interface{} {
	x := (*h)[h.Len()-1]
	*h = (*h)[0 : h.Len()-1]
	return x
}

func (h *maxHeap) Top() int {
	return (*h)[0]
}

type worker struct {
	ratio   float64
	quality int
}

// TODO: Need to understand
func mincostToHireWorkers(quality []int, wage []int, k int) float64 {
	n := len(quality)
	workers := make([]worker, n)

	for i := 0; i < n; i++ {
		workers[i] = worker{float64(wage[i]) / float64(quality[i]), quality[i]}
	}

	sort.Slice(workers, func(i, j int) bool {
		return workers[i].ratio < workers[j].ratio
	})

	pq := &maxHeap{}

	qualitySum := 0

	for i := 0; i < k; i++ {
		qualitySum += workers[i].quality
		heap.Push(pq, workers[i].quality)
	}

	minCost := float64(qualitySum) * workers[k-1].ratio

	for i := k; i < n; i++ {
		heap.Push(pq, workers[i].quality)
		qualitySum += workers[i].quality - heap.Pop(pq).(int)
		minCost = min(minCost, float64(qualitySum)*workers[i].ratio)
	}

	return minCost
}

func mincostToHireWorkersTest() {
	quality := []int{10, 20, 5}
	wage := []int{70, 50, 30}
	k := 2
	fmt.Println(mincostToHireWorkers(quality, wage, k)) // 105.00000
	quality = []int{3, 1, 10, 10, 1}
	wage = []int{4, 8, 2, 2, 7}
	k = 3
	fmt.Println(mincostToHireWorkers(quality, wage, k)) // 30.66667
}

// func main() {
// 	mincostToHireWorkersTest()
// }
