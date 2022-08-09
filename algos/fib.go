package main

import "fmt"

func main() {
	fmt.Println(FibDp(50))
}

func FibRecur(n int) int {
	if n <= 1 {
		return n
	}

	return FibRecur(n-1) + FibRecur(n-2)
}

func FibDp(n int) int {
	arr := []int{0, 1}
	for i := 2; i <= n; i++ {
		arr = append(arr, (arr[i-1] + arr[i-2]))
	}

	return arr[n]
}
