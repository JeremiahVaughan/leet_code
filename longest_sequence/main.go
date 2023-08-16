package main

import "container/heap"

func main() {

}

func longestConsecutive(nums []int) int {
	h := &MinHeap{}
	heap.Init(h)
	for _, num := range nums {
		heap.Push(h, num)
	}

	longestSequence := 0
	currentSequence := 0
	previousValue := 0
	for i := 0; i < len(nums); i++ {
		currentValue := heap.Pop(h).(int)
		if i != 0 {
			if previousValue == currentValue {
				//	do nothing because duplicates don't count for or against the current sequence
			} else if previousValue == currentValue-1 {
				currentSequence++
			} else {
				if longestSequence < currentSequence {
					longestSequence = currentSequence
				}
				currentSequence = 1
			}
		} else {
			currentSequence = 1
		}
		previousValue = currentValue
	}

	// If the longest sequenced happened to run to the bottom of the heap
	if longestSequence < currentSequence {
		longestSequence = currentSequence
	}

	return longestSequence
}

type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
