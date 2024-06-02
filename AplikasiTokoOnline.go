package main

import (
	"fmt"
	"os"
)

const adminPassword = "12345"
const maxProduk = 100
const maxTransaksi = 1000

type Transaksi struct {
	PembeliNama string
	ProdukNama  string
	Harga       float64
}

type Produk struct {
	Nama      string
	Harga     float64
	Deskripsi string
}

type Penjual struct {
	Nama      string
	Password  string // Menyimpan password penjual
	Terdaftar bool   // Menyimpan status registrasi
	Diterima  bool   // Menyimpan status persetujuan admin
}

type Pembeli struct {
	Nama      string
	Password  string // Menyimpan password pembeli
	Terdaftar bool   // Menyimpan status registrasi
	Diterima  bool   // Menyimpan status persetujuan admin
}

var (
	daftarTransaksi [maxTransaksi]Transaksi
	daftarProduk    [maxProduk]Produk
	daftarPenjual   [maxProduk]Penjual
	daftarPembeli   [maxProduk]Pembeli
	numProduk       int
	numPenjual      int
	numPembeli      int
	numTransaksi    int
)

func main() {
	showMainMenu()
}

func showMainMenu() {
	for {
		fmt.Println("=== Menu Login ===")
		fmt.Println("1. Login sebagai Admin")
		fmt.Println("2. Login sebagai Penjual")
		fmt.Println("3. Login sebagai Pembeli")
		fmt.Println("4. Registrasi Penjual")
		fmt.Println("5. Registrasi Pembeli")
		fmt.Println("6. Keluar")
		fmt.Print("Pilih opsi: ")

		var option int
		fmt.Scanln(&option)

		switch option {
		case 1:
			loginAdmin()
		case 2:
			loginPenjual()
		case 3:
			loginPembeli()
		case 4:
			registerPenjual()
		case 5:
			registerPembeli()
		case 6:
			fmt.Println("Keluar dari program.")
			os.Exit(0)
		default:
			fmt.Println("Opsi tidak valid, coba lagi.")
		}
	}
}

func loginAdmin() {
	fmt.Print("Masukkan password Admin: ")
	var password string
	fmt.Scanln(&password)

	if password == adminPassword {
		showAdminMenu()
	} else {
		fmt.Println("Password salah. Anda tidak memiliki akses sebagai Admin.")
	}
}

func showAdminMenu() {
	for {
		fmt.Println("=== Menu Admin ===")
		fmt.Println("1. Lihat Data Penjual dan Setujui Registrasi")
		fmt.Println("2. Lihat Data Pembeli dan Setujui Registrasi")
		fmt.Println("3. Logout")
		fmt.Print("Pilih opsi: ")

		var option int
		fmt.Scanln(&option)

		switch option {
		case 1:
			showPenjualData()
		case 2:
			showPembeliData()
		case 3:
			return
		default:
			fmt.Println("Opsi tidak valid, coba lagi.")
		}
	}
}

func showPenjualData() {
	fmt.Println("=== Data Penjual ===")
	for i := 0; i < numPenjual; i++ {
		if !daftarPenjual[i].Diterima {
			fmt.Printf("%d. %s\n", i+1, daftarPenjual[i].Nama)
		}
	}

	fmt.Println("Apakah Anda ingin menyetujui atau menolak registrasi penjual? (y/n): ")
	var response string
	fmt.Scanln(&response)
	if response == "y" || response == "Y" {
		fmt.Println("Masukkan nomor penjual yang akan disetujui atau ditolak: ")
		var index int
		fmt.Scanln(&index)
		if index > 0 && index <= numPenjual {
			fmt.Println("Apakah Anda ingin menyetujui (A) atau menolak (T)?")
			var decision string
			fmt.Scanln(&decision)
			if decision == "A" {
				daftarPenjual[index-1].Diterima = true
				fmt.Println("Registrasi penjual berhasil disetujui.")
			} else if decision == "T" {
				rejectPenjual(index)
				fmt.Println("Registrasi penjual berhasil ditolak.")
			} else {
				fmt.Println("Pilihan tidak valid.")
			}
		} else {
			fmt.Println("Nomor penjual tidak valid.")
		}
	}
}

func rejectPenjual(index int) {
	for i := index - 1; i < numPenjual-1; i++ {
		daftarPenjual[i] = daftarPenjual[i+1]
	}
	numPenjual--
}

func showPembeliData() {
	fmt.Println("=== Data Pembeli ===")
	for i := 0; i < numPembeli; i++ {
		if !daftarPembeli[i].Diterima {
			fmt.Printf("%d. %s\n", i+1, daftarPembeli[i].Nama)
		}
	}

	fmt.Println("Apakah Anda ingin menyetujui atau menolak registrasi pembeli? (y/n): ")
	var response string
	fmt.Scanln(&response)
	if response == "y" || response == "Y" {
		fmt.Println("Masukkan nomor pembeli yang akan disetujui atau ditolak: ")
		var index int
		fmt.Scanln(&index)
		if index > 0 && index <= numPembeli {
			fmt.Println("Apakah Anda ingin menyetujui (A) atau menolak (T)?")
			var decision string
			fmt.Scanln(&decision)
			if decision == "A" {
				daftarPembeli[index-1].Diterima = true
				fmt.Println("Registrasi pembeli berhasil disetujui.")
			} else if decision == "T" {
				rejectPembeli(index)
				fmt.Println("Registrasi pembeli berhasil ditolak.")
			} else {
				fmt.Println("Pilihan tidak valid.")
			}
		} else {
			fmt.Println("Nomor pembeli tidak valid.")
		}
	}
}

func rejectPembeli(index int) {
	for i := index - 1; i < numPembeli-1; i++ {
		daftarPembeli[i] = daftarPembeli[i+1]
	}
	numPembeli--
}

func loginPenjual() {
	fmt.Print("Masukkan nama Penjual: ")
	var nama string
	fmt.Scanln(&nama)

	fmt.Print("Masukkan password Penjual: ")
	var password string
	fmt.Scanln(&password)

	for i := 0; i < numPenjual; i++ {
		if daftarPenjual[i].Nama == nama && daftarPenjual[i].Password == password && daftarPenjual[i].Diterima {
			fmt.Println("Login berhasil sebagai Penjual.")
			// Panggil menu Penjual
			showPenjualMenu()
			return
		}
	}
	fmt.Println("Penjual tidak ditemukan, password salah, atau belum disetujui oleh admin.")
}

func loginPembeli() {
	fmt.Print("Masukkan nama Pembeli: ")
	var nama string
	fmt.Scanln(&nama)

	fmt.Print("Masukkan password Pembeli: ")
	var password string
	fmt.Scanln(&password)

	for i := 0; i < numPembeli; i++ {
		if daftarPembeli[i].Nama == nama && daftarPembeli[i].Password == password && daftarPembeli[i].Diterima {
			fmt.Println("Login berhasil sebagai Pembeli.")
			// Panggil menu Pembeli
			showPembeliMenu()
			return
		}
	}
	fmt.Println("Pembeli tidak ditemukan, password salah, atau belum disetujui oleh admin.")
}

func showPenjualMenu() {
	for {
		fmt.Println("=== Menu Penjual ===")
		fmt.Println("1. Tambahkan Produk")
		fmt.Println("2. Lihat Produk")
		fmt.Println("3. Edit Harga Produk")
		fmt.Println("4. Hapus Produk")
		fmt.Println("5. Logout")
		fmt.Print("Pilih opsi: ")

		var option int
		fmt.Scanln(&option)

		switch option {
		case 1:
			tambahProduk()
		case 2:
			lihatProduk()
		case 3:
			var namaProduk string
			fmt.Print("Masukkan nama produk yang akan diedit harganya: ")
			fmt.Scanln(&namaProduk)
			editProduk(namaProduk)
		case 4:
			hapusProduk()
		case 5:
			return
		default:
			fmt.Println("Opsi tidak valid, coba lagi.")
		}
	}
}

// Fungsi untuk mencari indeks produk berdasarkan nama menggunakan binary search
func binarySearchProduk(nama string) (int, bool) {
	// Mengurutkan array produk berdasarkan nama
	for i := 1; i < numProduk; i++ {
		key := daftarProduk[i]
		j := i - 1
		for j >= 0 && daftarProduk[j].Nama > key.Nama {
			daftarProduk[j+1] = daftarProduk[j]
			j = j - 1
		}
		daftarProduk[j+1] = key
	}

	// Melakukan binary search
	low := 0
	high := numProduk - 1
	for low <= high {
		mid := (low + high) / 2
		if daftarProduk[mid].Nama == nama {
			return mid, true // Produk ditemukan
		} else if daftarProduk[mid].Nama < nama {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	// Produk tidak ditemukan
	return -1, false
}

// Fungsi untuk mengedit harga produk berdasarkan namanya
func editProduk(namaProduk string) {
	index, found := binarySearchProduk(namaProduk)
	if found {
		// Meminta pengguna untuk memasukkan harga baru
		var hargaBaru float64
		fmt.Print("Masukkan harga baru: ")
		fmt.Scanln(&hargaBaru)

		// Memperbarui harga produk
		daftarProduk[index].Harga = hargaBaru

		fmt.Println("Harga produk berhasil diperbarui.")
	} else {
		fmt.Println("Produk tidak ditemukan.")
	}
}

func lihatProduk() {
	fmt.Println("=== Lihat Produk ===")
	if numProduk == 0 {
		fmt.Println("Tidak ada produk yang tersedia.")
		return
	}

	// Salin array produk untuk menghindari pengurutan di tempat
	daftarProdukCopy := make([]Produk, numProduk)
	copy(daftarProdukCopy, daftarProduk[:numProduk])

	// Selection sort untuk mengurutkan produk berdasarkan harga menurun
	for i := 0; i < numProduk-1; i++ {
		maxIndex := i
		for j := i + 1; j < numProduk; j++ {
			if daftarProdukCopy[j].Harga > daftarProdukCopy[maxIndex].Harga {
				maxIndex = j
			}
		}
		// Tukar posisi elemen
		daftarProdukCopy[i], daftarProdukCopy[maxIndex] = daftarProdukCopy[maxIndex], daftarProdukCopy[i]
	}

	// Menampilkan produk yang telah diurutkan
	fmt.Println("Daftar Produk:")
	for i, produk := range daftarProdukCopy {
		fmt.Printf("%d. %s - Rp %.2f\n", i+1, produk.Nama, produk.Harga)
	}
}

func tambahProduk() {
	fmt.Println("=== Tambah Produk ===")
	var produk Produk
	fmt.Print("Nama produk: ")
	fmt.Scanln(&produk.Nama)
	fmt.Print("Harga produk: ")
	fmt.Scanln(&produk.Harga)
	fmt.Print("Deskripsi produk: ")
	fmt.Scanln(&produk.Deskripsi)

	daftarProduk[numProduk] = produk
	numProduk++

	fmt.Println("Produk berhasil ditambahkan.")
}

func hapusProduk() {
	fmt.Println("=== Hapus Produk ===")
	if numProduk == 0 {
		fmt.Println("Tidak ada produk yang tersedia.")
		return
	}

	// Menampilkan daftar produk
	fmt.Println("Daftar Produk:")
	for i := 0; i < numProduk; i++ {
		fmt.Printf("%d. %s - Rp %.2f\n", i+1, daftarProduk[i].Nama, daftarProduk[i].Harga)
	}

	// Meminta pengguna untuk memasukkan nama produk yang ingin dihapus
	fmt.Print("Masukkan nama produk yang ingin dihapus: ")
	var namaProduk string
	fmt.Scanln(&namaProduk)

	// Sequential search untuk menemukan produk
	produkDitemukan := false
	var index int
	for i := 0; i < numProduk; i++ {
		if daftarProduk[i].Nama == namaProduk {
			produkDitemukan = true
			index = i
			break
		}
	}

	// Jika produk ditemukan
	if produkDitemukan {
		// Geser elemen-elemen di belakang produk yang dihapus
		for j := index; j < numProduk-1; j++ {
			daftarProduk[j] = daftarProduk[j+1]
		}
		numProduk--

		fmt.Println("Produk berhasil dihapus.")
	} else {
		fmt.Println("Produk tidak ditemukan.")
	}
}

func registerPenjual() {
	var penjual Penjual
	fmt.Print("Masukkan nama Penjual: ")
	fmt.Scanln(&penjual.Nama)

	fmt.Print("Masukkan password Penjual: ")
	fmt.Scanln(&penjual.Password)

	// Implementasi registrasi penjual
	daftarPenjual[numPenjual] = penjual
	numPenjual++
	fmt.Println("Registrasi Penjual berhasil.")
}

func showPembeliMenu() {
	var pembeliNama string

	for {
		fmt.Println("=== Menu Pembeli ===")
		fmt.Println("1. Lihat Produk")
		fmt.Println("2. Beli Produk")
		fmt.Println("3. Lihat Transaksi")
		fmt.Println("4. Logout")
		fmt.Print("Pilih opsi: ")
		var option int
		fmt.Scanln(&option)

		switch option {
		case 1:
			lihatProdukPembeli()
		case 2:
			beliProduk(pembeliNama)
		case 3:
			lihatTransaksi(pembeliNama)
		case 4:
			return
		default:
			fmt.Println("Opsi tidak valid, coba lagi.")
		}
	}
}

func lihatProdukPembeli() {
	if numProduk == 0 {
		fmt.Println("Tidak ada produk yang tersedia.")
		return
	}

	// Insertion Sort untuk mengurutkan produk berdasarkan harga secara menaik
	for i := 1; i < numProduk; i++ {
		key := daftarProduk[i]
		j := i - 1

		// Memindahkan elemen yang lebih besar dari key ke satu posisi di depan posisi saat ini
		for j >= 0 && daftarProduk[j].Harga > key.Harga {
			daftarProduk[j+1] = daftarProduk[j]
			j = j - 1
		}
		daftarProduk[j+1] = key
	}

	// Menampilkan daftar produk yang telah diurutkan
	fmt.Println("=== Daftar Produk ===")
	for i := 0; i < numProduk; i++ {
		produk := daftarProduk[i]
		fmt.Printf("%d. %s - Harga: %.2f\n", i+1, produk.Nama, produk.Harga)
		fmt.Printf("   Deskripsi: %s\n", produk.Deskripsi)
	}
}

// Fungsi untuk membeli produk
func beliProduk(pembeliNama string) {
	if numProduk == 0 {
		fmt.Println("Tidak ada produk yang tersedia untuk dibeli.")
		return
	}

	// Menampilkan daftar produk
	fmt.Println("=== Daftar Produk ===")
	for i := 0; i < numProduk; i++ {
		fmt.Printf("%d. %s - Rp %.2f\n", i+1, daftarProduk[i].Nama, daftarProduk[i].Harga)
	}

	// Meminta nama produk yang ingin dibeli
	fmt.Print("Masukkan nama produk yang ingin dibeli: ")
	var namaProduk string
	fmt.Scanln(&namaProduk)

	// Cari produk berdasarkan nama
	var produkIndex int = -1
	for i := 0; i < numProduk; i++ {
		if daftarProduk[i].Nama == namaProduk {
			produkIndex = i
			break
		}
	}

	// Validasi apakah produk ditemukan
	if produkIndex == -1 {
		fmt.Println("Produk tidak ditemukan.")
		return
	}

	produk := daftarProduk[produkIndex]

	if numTransaksi >= maxTransaksi {
		fmt.Println("Transaksi tidak dapat dilakukan karena batas maksimum transaksi tercapai.")
		return
	}

	// Menambahkan transaksi baru ke daftar transaksi
	daftarTransaksi[numTransaksi] = Transaksi{
		PembeliNama: pembeliNama,
		ProdukNama:  produk.Nama,
		Harga:       produk.Harga,
	}
	numTransaksi++
	fmt.Printf("Produk %s berhasil dibeli dengan harga %.2f\n", produk.Nama, produk.Harga)
}

// Fungsi untuk melihat transaksi pembeli
func lihatTransaksi(pembeliNama string) {
	fmt.Println("=== Daftar Transaksi ===")
	adaTransaksi := false
	// Menampilkan transaksi yang dilakukan oleh pembeli
	for i := 0; i < numTransaksi; i++ {
		transaksi := daftarTransaksi[i]
		if transaksi.PembeliNama == pembeliNama {
			fmt.Printf("Produk: %s, Harga: %.2f\n", transaksi.ProdukNama, transaksi.Harga)
			adaTransaksi = true
		}
	}
	if !adaTransaksi {
		fmt.Println("Belum ada transaksi yang dilakukan.")
	}
}

func registerPembeli() {
	var pembeli Pembeli
	fmt.Print("Masukkan nama Pembeli: ")
	fmt.Scanln(&pembeli.Nama)
	fmt.Print("Masukkan password Pembeli: ")
	fmt.Scanln(&pembeli.Password)

	// Implementasi registrasi pembeli
	daftarPembeli[numPembeli] = pembeli
	numPembeli++
	fmt.Println("Registrasi Pembeli berhasil.")
}
