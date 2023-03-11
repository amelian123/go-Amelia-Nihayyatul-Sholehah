package main

import "fmt"

type kendaraan struct {
	totalroda       int
	kecepatanperjam int
}

type mobil struct {
	kendaraan
}

func (m *mobil) berjalan() {
	m.tambahkecepatan(10)
}

func (m *mobil) tambahkecepatan(kecepatanbaru int) {
	m.kecepatanperjam += kecepatanbaru
}

func main() {
	mobilcepat := mobil{}
	mobilcepat.berjalan()
	mobilcepat.berjalan()
	mobilcepat.berjalan()

	mobillamban := mobil{}
	mobillamban.berjalan()

	fmt.Printf("Kecepatan mobil laju cepat: %d km/jam\n", mobilcepat.kecepatanperjam)
	fmt.Printf("Kecepatan mobil laju lamban: %d km/jam\n", mobillamban.kecepatanperjam)
}
