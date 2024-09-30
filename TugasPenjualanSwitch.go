package main

import "fmt"

func main() {

	// deklarasi variabel
	var namaCustomer, namaBarang string
	var harga float32 = 25000
	var quantity int32
	var hasilDiscount, totalHarga float32

	// input nama customer
	fmt.Print("Input nama customer = ")
	fmt.Scanf("%s\n", &namaCustomer)

	// input nama barang
	fmt.Print("Input nama barang = ")
	fmt.Scanf("%s\n", &namaBarang)

	// input quantity barang
	fmt.Print("Input quantity barang = ")
	fmt.Scanf("%d\n", &quantity)

	// kondisi discount, kalau lebih dari 5 dapat 10% selain itu 2%
	switch {
	case quantity > 5:
		hasilDiscount = 0.1
	default:
		hasilDiscount = 0.02
	}

	// proses
	subTotal := float32(quantity) * harga
	totalHarga = subTotal - (hasilDiscount * subTotal)

	//tampilkan hasil harga
	fmt.Println("Hasil Discount = ", hasilDiscount)
	fmt.Println("quantity = ", quantity)
	fmt.Println("harga = ", harga)
	fmt.Println("Total harga = ", totalHarga)
}
