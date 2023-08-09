package main

func containsDuplicate(nums []int) bool {
	dupeMap := make(map[int]bool)
	for _, num := range nums {
		_, ok := dupeMap[num]
		if ok {
			return true
		}
		dupeMap[num] = true
	}
	return false
}
