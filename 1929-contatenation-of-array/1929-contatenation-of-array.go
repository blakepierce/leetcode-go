package leetcode

func getConcatenation(nums []int) []int {
	return append(nums, nums...)
}

//go:noinline
func getConcatenation2(nums []int) []int {
	l := len(nums)
	retVal := make([]int, l*2)
	for i := 0; i < l; i++ {
		retVal[i] = nums[i]
		retVal[i+l] = nums[i]
	}
	return retVal
}
