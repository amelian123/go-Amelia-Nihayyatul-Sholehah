package main

import "fmt"

func primeNumber(n int) bool {
	if n < 2 {
		return false
	}

	primes := make([]bool, n+1)
	for i := 2; i <= n; i++ {
		primes[i] = true
	}

	for i := 2; i*i <= n; i++ {
		if primes[i] {
			for j := i * i; j <= n; j += i {
				primes[j] = false
			}
		}
	}

	return primes[n]
}

func main() {
	fmt.Println(primeNumber(1000000007)) // true
	fmt.Println(primeNumber(13))         // true
	fmt.Println(primeNumber(17))         // true
	fmt.Println(primeNumber(20))         // false
	fmt.Println(primeNumber(35))         // false
}
