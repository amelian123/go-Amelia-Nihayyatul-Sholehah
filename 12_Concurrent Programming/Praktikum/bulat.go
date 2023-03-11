package main

import (
	"fmt"
	"time"
)

func printMultiples(x int) {
	for i := 1; i <= 10; i++ {
		fmt.Println(x * i)
	}
}

func main() {
	var x int
	fmt.Print("input nilai x : ")
	fmt.Scanln(&x)

	ticker := time.NewTicker(3 * time.Second)
	defer ticker.Stop()

	go func() {
		for {
			select {
			case <-ticker.C:
				printMultiples(x)
			}
		}
	}()

	fmt.Scanln()
}
