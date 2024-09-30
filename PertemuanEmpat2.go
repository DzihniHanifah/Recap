package main

import "fmt"

func main() {

	var angka int = 5

	switch {
	case angka > 10:
		fmt.Println("angka lebih dari 10, angka anda:", angka)
	default:
		fmt.Println("angka kurang dari 10, angka anda:", angka)
	}
}
