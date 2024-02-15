package main

func romanToInt(s string) int {
	romans := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	var sum int
	var cache rune = 0

	for _, c := range s {
		if cache == 0 {
			cache = c
		}
		if romans[c] > romans[cache] {
			sum -= romans[cache] * 2
		}

		sum += romans[c]
		cache = c
	}

	return sum
}
