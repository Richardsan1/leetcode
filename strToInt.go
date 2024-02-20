package main

import "fmt"

func myAtoi(s string) int {
	INT_MAX := 2147483647
	var num int
	op := 0
	getnum := false

	for _, r := range s {
		if r == '-' {
			if op != 0 || getnum {
				break
			}
			op = -1
			getnum = true
			continue
		}
		if r == '+' {
			if op != 0 || getnum {
				break
			}
			op = 1
			getnum = true
			continue
		}
		if r >= '0' && r <= '9' {
			num = num*10 + int(r-48)
			getnum = true
			if num > INT_MAX {
				if op < 0 {
					return (INT_MAX * -1) - 1
				}
				return INT_MAX
			}

			continue
		}
		if r == ' ' && !getnum {
			continue
		}
		break
	}
	if op == -1 {
		num *= op
	}

	return num
}

func main() {
	fmt.Println(myAtoi("  +  413"))
}
