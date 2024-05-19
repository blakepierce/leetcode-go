package leetcode

func numIdenticalPairsWorse(nums []int) int {
	pairs := 0
	store := make(map[int]int)
	for _, num := range nums {
		if _, ok := store[num]; ok {
			store[num]++
		} else {
			store[num] = 0
		}
	}
	for _, count := range store {
		for count > 0 {
			pairs += count
			count--
		}
	}
	return pairs
}

func numIdenticalPairsBest(nums []int) int {
	pairs := 0
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] == nums[j] {
				pairs++
			}
		}
	}
	return pairs
}
