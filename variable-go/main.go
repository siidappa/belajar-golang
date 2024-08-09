package main

import "fmt"

func main() {
	/*
		# Penulisan Variable dengan tipe datanya (manifest typing)
	*/
	// var firstName string = "Dappa"
	// var lastName string = "Naufal"

	/*
		# Penulisan Variable tanpa tipe data (type inference)
		=> (:=) hanya digunakan sekali pada saat deklarasi awal, untuk assginment selanjutnya wajib menggunakan (=)
	*/
	// ? menggunakan keyword var, tanpa tipe-data, menggunakan perantara (=)
	// var firstName = "Dappa"

	// ? tanpa keyword var, tanpa tipe-data, menggunakan perantara (:=)
	// lastName := "Naufal"
	// lastName := "Fachrezi" => Error
	// lastName = "Fachrezi"

	// secondName := "S.Kom"

	/*
		# (%s) => akan digantikan dengan value dari nilai variable
		# (+) => digunakan untuk string concatenation
	*/
	//fmt.Printf("Halo %s %s %s !\n", firstName, lastName, secondName) // Output: Halo Dappa Fachrezi ! S.Kom
	//fmt.Println("Halo", firstName, lastName+" 😊", secondName)        // Output: Halo Dappa Fachezi 😊 S.Kom

	/*
		# Deklarasi multi variable
	*/
	// first, second, third := "Dappa", "Naufal", "Fachrezi"
	//fmt.Printf("Hallo 😊, %s %s %s", first, second, third) // Output: Hallo 😊 Daffa Naufal Fachrezi

	// Menggunakan metode type inference kita bisa assign value nya dengan berbagai tipe data
	// one, isSunday, twoPointTwo, say := 1, true, 2.2, "Hello"
	// fmt.Println(one, isSunday, twoPointTwo, say)

	/*
		# Deklarasi Variable Underscore (_)
		? Variable Underscore adalah predefined, jadi tidak perlu menggunakan (:=) untuk assign value nya cukup menggunakan (=)
		? (_) adalah reserved variable yang bisa dimanfaatkan untuk menampung nilai yang tidak dipakai. Bisa dibilang variabel ini merupakan keranjang sampah
		? Namun khusus untuk pengisian nilai multi variabel yang dilakukan dengan metode type inference, boleh di dalamnya terdapat variabel underscore.
		? Biasanya variabel underscore sering dimanfaatkan untuk menampung nilai balik fungsi yang tidak digunakan
	*/
	// _ = "Belajar Golang"
	// _ = "Golang itu mudah"
	// name, _ := "Dappa", "Naufal"

	//fmt.Println("Hallo", name) // => _ tidak bisa dicetak ke layar

	/*
		# Deklarasi Variable menggunakan fungsi new()
		=> Fungsi new() digunakan untuk membuat variabel pointer dengan tipe data tertentu. Nilai data default-nya akan menyesuaikan tipe datanya
	*/
	/*
		# Penjelasan Kode dibawah
		=> Variabel name menampung data bertipe pointer string. Jika ditampilkan yang muncul bukanlah nilainya melainkan alamat memori nilai tersebut (dalam bentuk notasi heksadesimal)
		=> Untuk menampilkan nilai aslinya, variabel tersebut perlu di-dereference terlebih dahulu, caranya dengan menuliskan tanda asterisk (*) sebelum nama variabel.
	*/
	name := new(string)
	fmt.Println(name) // Output: 0xc000028070 -> hexadecimal
	fmt.Println(*name)
}
