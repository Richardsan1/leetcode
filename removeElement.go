package main

func removeElement(nums []int, val int) int {
	posWrite := 0
	size := 0
	for posRead := 0; posRead < len(nums); posRead++ {
		if nums[posRead] != val {
			nums[posWrite] = nums[posRead]
			posWrite++
			size++
		}
	}
	return size
}

// func main() {
// 	val := removeElement([]int{0, 1, 2, 2, 3, 0, 4, 2}, 2)

// 	fmt.Println(val)
// }
