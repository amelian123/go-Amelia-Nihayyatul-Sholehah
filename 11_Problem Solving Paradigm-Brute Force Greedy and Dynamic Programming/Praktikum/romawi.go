package main

import "fmt"

func toRoman(num int) string {
	if num <= 0 || num > 3999 {
		return ""
	}

	var result string

	// Tambahkan huruf-huruf Romawi ke dalam hasil
	for num >= 1000 {
		result += "M"
		num -= 1000
	}

	if num >= 900 {
		result += "CM"
		num -= 900
	}

	if num >= 500 {
		result += "D"
		num -= 500
	}

	if num >= 400 {
		result += "CD"
		num -= 400
	}

	for num >= 100 {
		result += "C"
		num -= 100
	}

	if num >= 90 {
		result += "XC"
		num -= 90
	}

	if num >= 50 {
		result += "L"
		num -= 50
	}

	if num >= 40 {
		result += "XL"
		num -= 40
	}

	for num >= 10 {
		result += "X"
		num -= 10
	}

	if num >= 9 {
		result += "IX"
		num -= 9
	}

	if num >= 5 {
		result += "V"
		num -= 5
	}

	if num >= 4 {
		result += "IV"
		num -= 4
	}

	for num > 0 {
		result += "I"
		num -= 1
	}

	return result
}

func main() {
	var num int
	fmt.Print("Input : ")
	fmt.Scan(&num)
	fmt.Printf("Output : %s\n", toRoman(num))
}
