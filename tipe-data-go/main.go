package main

func main() {
	/*
		# Tipe data numerik non-desimal
	*/
	// var positiveNumber uint8 = 89
	// var negativeNumber = -1243423644

	// => (%d) -> untuk memformat data numerik non-desimal
	// fmt.Printf("bilangan positif: %d\n", positiveNumber)
	// fmt.Printf("bilangan negative: %d\n", negativeNumber)

	/*
		# Tipe data desimal / floating point
		=> (%f) -> untuk memformat data numerik desimal menjadi string, digit desimal defaultnya adalah 6 digit
		=> tetapi kalian bisa atur jumlah digit yang muncul menggunakan (%.<angka>f)
	*/
	// decimalNumber := 3.14
	// fmt.Printf("bilangan desimal: %f\n", decimalNumber)
	// fmt.Printf("bilangan desimal: %.2f\n", decimalNumber)

	/*
		# Tipe data bool(Boolean)
		? hanya ada 2 value yaitu true dan false
		! biasanya digunakan di logical statement atau perulangan
	*/
	// var exist bool = false
	// => (%t) -> memformat data boolean menggunakan fungsi fmt.Printf()
	// fmt.Printf("exist? %t \n", exist)

	/*
		# Tipe data string
	*/
	// => string menggunakan double quote("")
	// var message string = "Hello"
	// fmt.Printf("message: %s \n", message)

	// => string menggunakan backticks
	// var message = `
	// 	Nama saya Daffa Naufal Fachrezi
	// 	Saya sekolah di SMK Gelora Bekasi
	// `
	// fmt.Println(message)

	/*
		# Nilai nil dan zero value
		? nil bukan merupakan tipe data, melainkan sebuah nilai. Variabel yang isi nilainya nil berarti memiliki nilai kosong.
		    => Artinya meskipun variabel dideklarasikan dengan tanpa nilai awal, tetap akan ada nilai default-nya
			    => Zero value dari string adalah "" (string kosong).
				=> Zero value dari bool adalah false.
				=> Zero value dari tipe numerik non-desimal adalah 0.
				=> Zero value dari tipe numerik desimal adalah 0.0

			? Beberapa tipe data yang bisa di-set nilainya dengan nil, di antaranya:
			=> pointer
			=> tipe data fungsi
			=> slice
			=> map
			=> channel
			=> interface kosong atau any (yang merupakan alias dari interface{})
	*/
}
