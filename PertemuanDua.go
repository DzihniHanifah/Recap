package main

import "fmt"

func main() {

	var nama, alamat string = "Dzihni Hanifah", "Bogor"
	var berkacamata bool = true
	var umur int = 18
	var minuskiri, minuskanan float32 = 2.5, 2.5
	var minustotal float32 = minuskiri + minuskanan

	fmt.Println("Nama = ", nama)
	fmt.Println("Umur = ", umur)
	fmt.Println("Alamat = ", alamat)
	fmt.Println("Berkacamata = ", berkacamata)
	fmt.Println("Minus Kanan = ", minuskanan)
	fmt.Println("Minus Kiri = ", minuskiri)
	fmt.Println("Total Minus = ", minustotal)

}
