package main

import "fmt"

func main() {
	// # Contoh penggunaan konstanta
	// const panjang = 18
	// const lebar = 8
	// const tinggi = 5
	// var hasil = panjang * lebar * tinggi
	// fmt.Println("Nilai p x l x y adalah", hasil)

	/*
		# Deklarasi Multi Konstanta
	*/
	// const (
	// 	square         = "Kotak" // Type Inference
	// 	isToday  bool  = true    // Manifest Typing
	// 	numeric  uint8 = 9       // Manifest Typing
	// 	floatNum       = 3.14    // Type Inference
	// )

	// const (
	// 	a = "Konstanta"
	// 	b // Tipe data dan value akan mengikuti variable yang sudah di deklarasikan diatas / sebelumnya
	// )
	// fmt.Println(a, b)

	/*
		# Deklarasi variable konstanta satu baris
	*/
	const satu, dua = 1, 2                     // Type inference
	const three, four string = "tiga", "empat" // Manifest Typing
	fmt.Println(satu, dua)
	fmt.Println(three, four)
}
