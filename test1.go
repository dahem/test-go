package main

import "fmt"

func run(num int, items []int) int {
	for _, val := range items {
		fmt.Println(val)	
	}
	return num
}

func main() {
	var N int
	fmt.Scan(&N)
	var facts [8]int
	facts[0] = 1
	for i := 1; i < 8; i++ {
		facts[i] = facts[i-1] * (i + 1)
	}
	fmt.Println(facts)
	fmt.Println(N)
	run(N, facts[:])
	
}