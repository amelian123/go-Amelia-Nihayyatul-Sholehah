package main

import (
	"fmt"
	"math"
)

func pow(x, n int) int {
	return int(math.Pow(float64(x), float64(n)))
}

func main() {
	fmt.Println(pow(2, 3))  // 8
	fmt.Println(pow(5, 3))  // 125
	fmt.Println(pow(10, 2)) // 100
	fmt.Println(pow(2, 5))  // 32
	fmt.Println(pow(7, 3))  // 343}
}
