package main

func largestDivisibleSubset(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}

	memo := [][]int{}
	subset := []int{}

	for i := 1; i <= len(nums); {
		if i == len(nums) && len(nums)%2 == 1 {
			subset = append(subset, nums[i-1])
			i++
			continue
		}
		if nums[i-1]%nums[i] == 0 || nums[i]%nums[i-1] == 0 {
			subset = append(subset, nums[i-1], nums[i])
			i += 2
		} else if i+1 < len(nums) {
			if nums[i+1]%nums[i] == 0 || nums[i]%nums[i+1] == 0 {
				subset = append(subset, nums[i], nums[i+1])
				i += 3
			}
		} else {
			memo = append(memo, subset)
			subset = []int{}
			i += 2
		}
	}
	memo = append(memo, subset)
	subset = []int{}

	bigSub := []int{}
	for _, sub := range memo {
		if len(sub) > len(bigSub) {
			bigSub = sub
		}
	}
	if len(bigSub) == 0 {
		return []int{nums[0]}
	}
	return bigSub
}
