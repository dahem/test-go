package main

// import "fmt"

func factorial(n int) int {
	if n == 1 {
		return 1
	}
	return n * factorial(n-1)
}

func factorialIterative(n int) int{
	fac := 1
	for i := 2; i <= n; i++ {
		fac = fac * i
	}
	return fac
}

// func main() {
// 	fac := factorialIterative(5)
// 	fmt.Println(fac)
// }