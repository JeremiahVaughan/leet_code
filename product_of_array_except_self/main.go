package main

func main() {

}
func productExceptSelf2(nums []int) []int {
	var result []int
	for i := 0; i < len(nums); i++ {
		currentProduct := 1
		for j := 0; j < len(nums); j++ {
			if i != j {
				currentProduct *= nums[j]
			}
		}
		result = append(result, currentProduct)
	}
	return result
}

func productExceptSelf(nums []int) []int {
	res := make([]int, len(nums))

	multiplier := 1
	for i := 0; i < len(nums); i++ {
		res[i] = multiplier
		multiplier *= nums[i]
	}

	multiplier = 1
	for i := len(res) - 1; i >= 0; i-- {
		res[i] *= multiplier
		multiplier *= nums[i]
	}

	return res
}
