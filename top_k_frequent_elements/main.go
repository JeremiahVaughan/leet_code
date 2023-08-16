package main

import "container/heap"

func main() {

}
func topKFrequent(nums []int, k int) []int {
	// key is the number, value is the occurrence
	occurrenceMap := make(map[int]int)
	for _, num := range nums {
		occurrenceMap[num] = occurrenceMap[num] + 1
	}

	h := &MaxHeap{}
	heap.Init(h)

	for num, score := range occurrenceMap {
		heap.Push(h, Treasure{
			num:   num,
			Score: score,
		})
	}

	var result []int
	for i := 0; i < k; i++ {
		t := heap.Pop(h).(Treasure)
		result = append(result, t.num)
	}

	return result

}

type Treasure struct {
	num   int
	Score int
}

type MaxHeap []Treasure

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i].Score > h[j].Score } // Compare by treasure value!
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(Treasure))
}

func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
