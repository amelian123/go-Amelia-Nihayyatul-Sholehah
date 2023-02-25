package main

import "fmt"

// Fungsi untuk mencari angka yang hanya muncul satu kali dalam sebuah string
func munculSekali(input string) []int {
	// Buat sebuah map untuk menyimpan jumlah kemunculan dari setiap angka dalam string
	m := make(map[int]int)

	// Tambahkan setiap angka dalam string ke dalam map dan tingkatkan jumlah kemunculannya
	for _, s := range input {
		m[int(s-'0')]++
	}

	// Buat sebuah slice untuk menyimpan angka yang hanya muncul satu kali dalam string
	var result []int
	for k, v := range m {
		if v == 1 {
			result = append(result, k)
		}
	}

	// Kembalikan slice yang berisi angka yang hanya muncul satu kali dalam string
	return result
}

func main() {
	fmt.Println(munculSekali("1234123"))
	fmt.Println(munculSekali("76523752"))
	fmt.Println(munculSekali("12345"))
	fmt.Println(munculSekali("1122334455"))
	fmt.Println(munculSekali("0872504"))
}
