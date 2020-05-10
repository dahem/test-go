package main

import (
	"fmt"
)

// func Min(x, y int) int {
//     if x > y {
//         return y
//     }
//     return x
// }

// var dp[100000][8] int
// var facts [8]int


// func Run(num int, items []int) int {
// 	if num == 0 {
// 		return 0
// 	}
// 	if num == 1 {
// 		return 1
// 	}
// 	var res int = 1000000
// 	for _, val := range items {
// 		if val <= num {
// 			res = Min(res, 1 + Run(num - val, items))
// 		}
// 	}
// 	// dp[num][]
// 	return res
// }

func main() {
	var N int
	fmt.Scan(&N)
	var facts [8]int
	facts[0] = 1
	for i := 1; i < 8; i++ {
		facts[i] = facts[i-1] * (i + 1)
	}

	var ans int = 0
	for i := 7; i >= 0; i-- {
		if facts[i] <= N && N > 0 {
			ans = ans + (N / facts[i])
			N = N % facts[i]
		} 
	}

	fmt.Println(ans)
}