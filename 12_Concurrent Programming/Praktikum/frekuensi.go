package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	freq := make(map[rune]int)

	// input teks
	fmt.Println("Input teks:")
	var text string
	fmt.Scan(&text)

	// tambahkan goroutine untuk menghitung frekuensi huruf dalam teks
	wg.Add(1)
	go func(t string) {
		defer wg.Done()
		for _, r := range t {
			freq[r]++
		}
	}(text)

	// tunggu semua goroutine selesai
	wg.Wait()

	// tampilkan output
	for r, f := range freq {
		fmt.Printf("%c: %d\n", r, f)
	}
}
