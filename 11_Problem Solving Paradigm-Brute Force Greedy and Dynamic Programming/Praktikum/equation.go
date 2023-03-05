package main

import "fmt"

func SimpleEquations(a, b, c int) {
	//kodeku bree
	found := false
	for x := -1000; x <= 1000 && !found; x++ {
		for y := -1000; y <= 1000 && !found; y++ {
			z := a - x - y
			if x != y && x != z && y != z && x+y+z == a && x*y*z == b && x*x+y*y+z*z == c {
				fmt.Println(x, y, z)
				found = true
			}
		}
	}
	if !found {
		fmt.Println("no solutions")
	}
}

func main() {
	SimpleEquations(1, 2, 3)  // no solutions
	SimpleEquations(6, 6, 14) //1 2 3
}
