package main

func main() {

}

func twoSum(nums []int, target int) []int {
	for i, num := range nums {
		for j, compareNum := range nums {
			if num+compareNum == target && i != j {
				return []int{i, j}
			}
		}
	}
	return []int{-1}
}
