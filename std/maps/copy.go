package main

import (
	"fmt"
	"maps"
)

func main() {
	m1 := map[string]int{
		"one": 1,
		"two": 2,
	}
	m2 := map[string]int{
		"one": 10,
	}

	maps.Copy(m2, m1)
	fmt.Println("m2 is:", m2) // one:1, two:2

	m2["one"] = 100
	fmt.Println("m1 is:", m1) //one:1, two:2
	fmt.Println("m2 is:", m2) //one: 100, two:2

	m3 := map[string][]int{
		"one": {1, 2, 3},
		"two": {4, 5, 6},
	}
	m4 := map[string][]int{
		"one": {7, 8, 9},
	}

	maps.Copy(m4, m3)
	fmt.Println("m4 is:", m4) //one:[1,2,3], two:[4,5,6]

	m4["one"][0] = 100
	fmt.Println("m3 is:", m3) //m3 and m4 is same
	fmt.Println("m4 is:", m4)

}
