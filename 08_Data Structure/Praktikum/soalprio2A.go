package main

import "fmt"

func PairSum(arr []int, target int) []int {
	//tulis kode bree
	// Definisikan indeks awal dan akhir dalam array
	start := 0
	end := len(arr) - 1

	// Loop hingga ditemukan pasangan yang jumlahnya sama dengan target
	for start < end {
		// Hitung jumlah dari dua angka pada indeks start dan end
		sum := arr[start] + arr[end]

		// Jika jumlahnya sama dengan target, kembalikan indeks start dan end sebagai pasangan
		if sum == target {
			return []int{start, end}
		} else if sum < target {
			// Jika jumlahnya kurang dari target, tambahkan indeks start
			start++
		} else {
			// Jika jumlahnya lebih dari target, kurangi indeks end
			end--
		}
	}

	// Jika tidak ditemukan pasangan yang jumlahnya sama dengan target, kembalikan slice kosong
	return []int{}
}
func main() {
	fmt.Println(PairSum([]int{1, 2, 3, 4, 6}, 6))
	fmt.Println(PairSum([]int{2, 5, 9, 11}, 11))
	fmt.Println(PairSum([]int{1, 3, 5, 7}, 12))
	fmt.Println(PairSum([]int{1, 4, 6, 8}, 10))
	fmt.Println(PairSum([]int{1, 5, 6, 7}, 6))
}
