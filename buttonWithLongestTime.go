package main

func buttonWithLongestTime(events [][]int) int {
	var maiorTempo int = events[0][1]
	var maiorInd int = events[0][0]
	for i := 1; i < len(events); i++ {
		if maiorTempo < (events[i][1] - events[i-1][1]) {
			maiorTempo = events[i][1] - events[i-1][1]
			maiorInd = events[i][0]
		} else if maiorTempo == (events[i][1] - events[i-1][1]) {
			if maiorInd > events[i][0] {
				maiorInd = events[i][0]
			}
		}
	}

	return maiorInd
}

// func main() {
// 	fmt.Println(buttonWithLongestTime([][]int{{9, 4}, {19, 5}, {2, 8}, {3, 11}, {2, 15}}))
// }
