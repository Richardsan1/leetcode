package main

import "fmt"

func maxScoreSightseeingPair(values []int) int {
	var vali int
	var valj int = 1
	var backi int
	for i := 1; i < len(values); i++ {
		if backi < valj {
			vali = backi
		}
		if (values[i]-i > values[valj]-valj) || (values[i]-i+values[backi]+backi > values[valj]-valj+values[backi]+backi) {
			valj = i
		}

		if values[i]+i > values[backi]+backi {
			backi = i
		}
	}

	if backi < valj {
		return values[backi] + backi + values[valj] - valj
	}
	return values[vali] + vali + values[valj] - valj
}

func main() {
	fmt.Println(maxScoreSightseeingPair([]int{2, 10, 9, 3, 2}))
}
