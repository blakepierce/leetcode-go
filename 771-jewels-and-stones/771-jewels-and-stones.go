package leetcode

// 669.7 ns/op           224 B/op         1 allocs/op
func numJewelsInStones(jewels string, stones string) int {
	retval := 0
	store := make(map[rune]int, len(jewels))
	for _, jewel := range jewels {
		store[jewel] = 0
	}
	for _, stone := range stones {
		if _, ok := store[stone]; ok {
			retval++
		}
	}

	return retval
}

// 137.9 ns/op             0 B/op         0 allocs/op
func numJewelsInStones2(jewels string, stones string) int {
	retval := 0
	for j := 0; j < len(jewels); j++ {
		for s := 0; s < len(stones); s++ {
			if jewels[j] == stones[s] {
				retval++
			}
		}
	}

	return retval
}
