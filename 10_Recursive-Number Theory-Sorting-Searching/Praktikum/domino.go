package main

import (
	"fmt"
)

func playingDomino(cards [][]int, deck []int) interface{} {
	//kodeku bree
	var match bool
	var card []int
	for _, c := range cards {
		if c[0] == deck[0] || c[0] == deck[1] || c[1] == deck[0] || c[1] == deck[1] {
			match = true
			card = c
			break
		}
	}

	// output hasil
	if match {
		return card
	} else {
		return "tutup kartu"
	}
}

func main() {
	fmt.Println(playingDomino([][]int{[]int{6, 5}, []int{3, 4}, []int{2, 1}, []int{3, 3}}, []int{4, 3})) //[3, 4]
	fmt.Println(playingDomino([][]int{[]int{6, 5}, []int{3, 3}, []int{3, 4}, []int{2, 1}}, []int{3, 6})) //[6, 5]
	fmt.Println(playingDomino([][]int{[]int{6, 6}, []int{2, 4}, []int{3, 6}}, []int{5, 1}))              //"tutup kartu"
}
