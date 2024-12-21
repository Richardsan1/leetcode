package main

// minha versao
// func missingNumber(nums []int) int {
// 	max := 0
// 	aux_nums := make([]bool, len(nums)+1)
// 	for _, n := range nums {
// 		if n > max {
// 			max = n
// 		}
// 		aux_nums[n] = true
// 	}
// 	var val int = -1
// 	for n := 0; n < max; n++ {
// 		if !aux_nums[n] {
// 			val = n
// 			break
// 		}
// 	}
// 	if val == -1 {
// 		return max + 1
// 	}
// 	return val
// }

// versao do leet
func missingNumber(nums []int) int {
	sum := 0
	total := 0
	for n := 0; n < len(nums); n++ {
		sum += nums[n]
		total += n + 1
	}

	return total - sum
}

// func main() {
// 	fmt.Println(missingNumber([]int{45, 35, 38, 13, 12, 23, 48, 15, 44, 21, 43, 26, 6, 37, 1, 19, 22, 3, 11, 32, 4, 16, 28, 49, 29, 36, 33, 8, 9, 39, 46, 17, 41, 7, 2, 5, 27, 20, 40, 34, 30, 25, 47, 0, 31, 42, 24, 10, 14}))
// }
