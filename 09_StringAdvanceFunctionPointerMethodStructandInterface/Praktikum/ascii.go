package main

import "fmt"

func caesar(offset int, input string) string {
	// kodemu bree
	caesar := ""

	// Iterasi setiap karakter pada string input
	for _, char := range input {
		if char == ' ' {
			caesar += " "
			continue
		}

		// Hitung posisi baru karakter setelah pergeseran
		newPosition := int(char) - 'a' + offset
		newPosition = newPosition % 26

		// Tambahkan karakter baru pada string hasil pergeseran
		caesar += string('a' + newPosition)
	}

	return caesar
}

func main() {
	fmt.Println(caesar(3, "abc"))                           // def
	fmt.Println(caesar(2, "alta"))                          // def
	fmt.Println(caesar(10, "alterraacademy"))               // kvdobbkkmknowi
	fmt.Println(caesar(1, "abcdefghijklmnopqrstuvwxyz"))    // bcdefghijklmnopqrstuvwxyza
	fmt.Println(caesar(1000, "abcdefghijklmnopqrstuvwxyz")) // mnopqrstuvwxyzabcdefghijkl
}
