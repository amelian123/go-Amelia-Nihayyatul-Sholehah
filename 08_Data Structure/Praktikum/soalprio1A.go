package main

import "fmt"

func mergeArrays(arr1, arr2 []string) []string {
	merged := make(map[string]bool)

	// Tambahkan nilai dari arr1 ke dalam map merged
	for _, val := range arr1 {
		merged[val] = true
	}

	// Tambahkan nilai dari arr2 ke dalam map merged jika belum ada
	for _, val := range arr2 {
		if _, ok := merged[val]; !ok {
			merged[val] = true
		}
	}
	// Buat sebuah slice baru dari nilai-nilai yang sudah ada di dalam map merged
	var result []string
	for key := range merged {
		result = append(result, key)
	}

	// Kembalikan slice baru yang sudah digabungkan dan tidak memiliki duplikat
	return result
}

func main() {
	//tes case
	fmt.Println(mergeArrays([]string{"king", "devil jin", "akuma"}, []string{"eddie", "steve", "geese"}))
	fmt.Println(mergeArrays([]string{"sergei", "jin", "akuma"}, []string{"jin", "steve", "bryan"}))
	fmt.Println(mergeArrays([]string{"alisa", "yoshimitsu"}, []string{"devil jin", "yoshimitsu", "alisa", "law"}))
	fmt.Println(mergeArrays([]string{}, []string{"devil jin", "sergei"}))
	fmt.Println(mergeArrays([]string{"hwoarang"}, []string{}))
	fmt.Println(mergeArrays([]string{}, []string{}))
}
