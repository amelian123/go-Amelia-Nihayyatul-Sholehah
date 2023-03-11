package main

import (
	"fmt"
	"sync"
)

func factorial(n int, mu *sync.Mutex, wg *sync.WaitGroup) {
	mu.Lock() // locking mutex
	defer mu.Unlock()

	fact := 1
	for i := 1; i <= n; i++ {
		fact *= i
	}
	fmt.Printf("Faktorial dari %d adalah %d\n", n, fact)
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	var mu sync.Mutex

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go factorial(i, &mu, &wg)
	}

	wg.Wait()
	fmt.Println("Semua goroutine berhasil di eksekusi.")
}
