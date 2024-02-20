package main

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	tmp := x
	var reverse int
	for tmp > 0 {
		reverse = reverse*10 + tmp%10
		tmp = tmp / 10
	}

	return x == reverse
}
