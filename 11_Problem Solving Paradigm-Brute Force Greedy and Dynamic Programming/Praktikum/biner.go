package main

import "fmt"

func main() {
	var n int
	fmt.Print("Masukkan nilai n: ")
	fmt.Scan(&n)

	ans := make([]int, n+1)
	for i := 0; i <= n; i++ {
		ans[i] = countSetBits(i)
	}

	fmt.Println("Array ans: ", ans)
}

func countSetBits(num int) int {
	count := 0
	for num > 0 {
		count += num & 1
		num >>= 1
	}
	return count
}
