package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// deklarasi variabel
	var input int16
	targetTebakan := generateTebakan(1, 10)
	var attempts int16 = 0

	for {
		attempts++

		// lakukan input
		fmt.Print("Masukkan Angka: ")
		fmt.Scanf("%d\n", &input)

		// proses
		if input > int16(targetTebakan) {
			fmt.Println("Angka yang ditebak kebesaran")
		}
		if input < int16(targetTebakan) {
			fmt.Println("Angka yang ditebak kekecilan")
		}
		if input == int16(targetTebakan) {
			fmt.Println("Selamat target tebakan benar!")
			break
		}

		if attempts == 3 {
			fmt.Println("Selamat target tebakan benar!")
			break
		}
	}

	// tampilkan output
	fmt.Println("Output inputan anda:", input)
	fmt.Println("Angka dari program:", targetTebakan)
}

func generateTebakan(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min+1) + min
}
