package main

import (
	"fmt"
)

func Compare(a, b string) string {
	//kodemu bree
	var longestSubstring string
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(b); j++ {
			k := 0
			for i+k < len(a) && j+k < len(b) && a[i+k] == b[j+k] {
				k++
			}
			if k > len(longestSubstring) {
				longestSubstring = a[i : i+k]
			}
		}
	}
	if len(longestSubstring) == 0 {
		return "Tidak ada substring yang sama di antara kedua string"
	} else {
		return fmt.Sprintf("%s dan %s adalah %s\n", a, b, longestSubstring)
	}
}
func main() {
	fmt.Println(Compare("AKA", "AKASHI"))
	fmt.Println(Compare("KANGOROO", "KANG"))
	fmt.Println(Compare("KI", "KIJANG"))
	fmt.Println(Compare("KUPU-KUPU", "KUPU"))
	fmt.Println(Compare("ILALANG", "ILA"))
}
