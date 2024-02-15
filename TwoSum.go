package main

// O(n)
func twoSum(nums []int, target int) []int {
	memo := make(map[int]int, len(nums))

	for i, num := range nums {
		index, err := memo[num]
		if err {
			return []int{i, index}
		} else {
			memo[target-num] = i
		}
	}

	return []int{}
}
