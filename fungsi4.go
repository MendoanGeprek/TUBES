package main

import (
	"fmt"
)

type Pabrikan struct {
	ID          int
	Nama        string
	Negara      string
	JumlahMobil int
}

type Mobil struct {
	ID         int
	Nama       string
	IDPabrikan int
	Tahun      int
	Harga      float64
	Penjualan  int
}

var pabrikanList = []Pabrikan{
	{ID: 1, Nama: "Toyota", Negara: "Jepang", JumlahMobil: 2},
	{ID: 2, Nama: "Ford", Negara: "Amerika Serikat", JumlahMobil: 1},
	{ID: 3, Nama: "Honda", Negara: "Jepang", JumlahMobil: 1},
}

var mobilList = []Mobil{
	{ID: 1, Nama: "Corolla", IDPabrikan: 1, Tahun: 2020, Harga: 300, 000, 000, Penjualan: 5, 000},
	{ID: 2, Nama: "Mustang", IDPabrikan: 2, Tahun: 2021, Harga: 800, 000, 000, Penjualan: 6, 000},
	{ID: 3, Nama: "Civic", IDPabrikan: 3, Tahun: 2022, Harga: 400, 000, 000, Penjualan: 4, 000},
}

func daftarMobilTerlaris() {
	var pilihanUrutan int
	fmt.Println("Pilih :")
	fmt.Println("1. Tertinggi ke Terendah")
	fmt.Println("2. Terendah ke Tertinggi")
	fmt.Print("Masukkan pilihan: ")
	fmt.Scan(&pilihanUrutan)

	for i := 0; i < len(mobilList)-1; i++ {
		indexTerpilih := i
		for j := i + 1; j < len(mobilList); j++ {
			if (pilihanUrutan == 1 && mobilList[j].Penjualan > mobilList[indexTerpilih].Penjualan) ||
				(pilihanUrutan == 2 && mobilList[j].Penjualan < mobilList[indexTerpilih].Penjualan) {
				indexTerpilih = j
			}
		}
		mobilList[i], mobilList[indexTerpilih] = mobilList[indexTerpilih], mobilList[i]
	}

	fmt.Println("\n3 Mobil Terlaris:")
	for i := 0; i < 3 && i < len(mobilList); i++ {
		mobil := mobilList[i]
		var namaPabrikan string
		for _, pabrikan := range pabrikanList {
			if pabrikan.ID == mobil.IDPabrikan {
				namaPabrikan = pabrikan.Nama
				break
			}
		}
		fmt.Printf("Nama: %s, Pabrikan: %s, Penjualan: %d\n",
			mobil.Nama, namaPabrikan, mobil.Penjualan)
	}
}

func main() {
	for {
		fmt.Println("\nMenu:")
		fmt.Println("1. Lihat Mobil Terlaris")
		fmt.Println("2. Keluar")
		fmt.Print("Masukkan pilihan: ")

		var pilihan int
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			daftarMobilTerlaris()
		case 2:
			fmt.Println("Keluar dari program.")
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}
