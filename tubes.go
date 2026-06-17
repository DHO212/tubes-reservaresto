package main

import "fmt"

const NMAX = 100

type meja struct {
	id        string
	nomor     int
	kapasitas int
	tersedia  bool
}
type tabMeja [NMAX]meja

type pelanggan struct {
	id   string
	nama string
	noHP string
}
type tabPelanggan [NMAX]pelanggan

type reservasi struct {
	id          string
	idPelanggan string
	idMeja      string
	tanggal     string
	jam         string
}
type tabReservasi [NMAX]reservasi

// fungsi helper untuk validasi input angka
func inputAngka(prompt string) int {
	var input string
	var angka int
	var valid bool
	var i int

	valid = false
	for !valid {
		fmt.Print(prompt)
		fmt.Scan(&input)

		valid = true
		for i = 0; i < len(input); i++ {
			if input[i] < '0' || input[i] > '9' {
				valid = false
			}
		}

		if len(input) == 0 {
			valid = false
		}

		if !valid {
			fmt.Println(">> Input harus berupa angka! Silakan masukkan ulang.")
		} else {
			angka = 0
			for i = 0; i < len(input); i++ {
				angka = angka*10 + int(input[i]-'0')
			}
		}
	}

	return angka
}

// fungsi helper untuk validasi input nomor HP (harus angka semua)
func inputNoHP(prompt string) string {
	var input string
	var valid bool
	var i int

	valid = false
	for !valid {
		fmt.Print(prompt)
		fmt.Scan(&input)

		valid = true
		for i = 0; i < len(input); i++ {
			if input[i] < '0' || input[i] > '9' {
				valid = false
			}
		}

		if len(input) == 0 {
			valid = false
		}

		if !valid {
			fmt.Println(">> Nomor HP harus berupa angka! Silakan masukkan ulang.")
		}
	}

	return input
}

func main() {
	var pilihan int

	var dataMeja tabMeja
	var nMeja int

	var dataPelanggan tabPelanggan
	var nPelanggan int

	var dataReservasi tabReservasi
	var nReservasi int

	pilihan = -1

	for pilihan != 0 {

		fmt.Println("===================================")
		fmt.Println("         RESERVARESTO")
		fmt.Println("===================================")
		fmt.Println("| 1 . Kelola Meja                |")
		fmt.Println("| 2 . Kelola Pelanggan           |")
		fmt.Println("| 3 . Menu Reservasi             |")
		fmt.Println("| 4 . Menu Pencarian             |")
		fmt.Println("| 5 . Menu Pengurutan            |")
		fmt.Println("| 6 . Statistik reservasi        |")
		fmt.Println("| 0 . Keluar                     |")
		fmt.Println("===================================")
		fmt.Print("Pilihan : ")

		pilihan = inputAngka("")

		switch pilihan {

		case 1:
			menuKelolaMeja(&dataMeja, &nMeja)

		case 2:
			menuKelolaPelanggan(&dataPelanggan, &nPelanggan)

		case 3:
			menuReservasi(&dataReservasi, &nReservasi)

		case 4:
			menuPencarian(&dataMeja, nMeja)

		case 5:
			menuPengurutan(&dataMeja, nMeja)

		case 6:
			statistikReservasi(dataReservasi, nReservasi)

		}
	}

	fmt.Println("Terima kasih... anda puas kami pun senang :)")
}

// menambahkan data meja
func tambahMeja(T *tabMeja, n *int) {

	fmt.Print("ID Meja      : ")
	fmt.Scan(&T[*n].id)

	T[*n].nomor = inputAngka("Nomor Meja   : ")

	T[*n].kapasitas = inputAngka("Kapasitas    : ")

	T[*n].tersedia = true

	*n = *n + 1

	fmt.Println("Data berhasil ditambahkan")
}

// menampilkan data meja
func tampilMeja(T tabMeja, n int) {
	var i int

	fmt.Println("================================================")
	fmt.Println("| ID   | NOMOR MEJA | KAPASITAS | TERSEDIA     |")
	fmt.Println("================================================")

	for i = 0; i < n; i++ {
		fmt.Printf("| %-4s | %-11d | %-9d | %-12t |\n",
			T[i].id,
			T[i].nomor,
			T[i].kapasitas,
			T[i].tersedia)
	}

	fmt.Println("================================================")
}

// ubah data meja
func ubahMeja(T *tabMeja, n int) {
	var nomor, i int
	var ketemu bool

	nomor = inputAngka("Masukkan nomor meja yang mau kamu ubah: ")

	ketemu = false
	i = 0

	for i < n && !ketemu {
		if T[i].nomor == nomor {
			ketemu = true

			fmt.Print("ID Baru : ")
			fmt.Scan(&T[i].id)

			T[i].nomor = inputAngka("Nomor Meja Baru : ")

			T[i].kapasitas = inputAngka("Kapasitas Baru : ")

			fmt.Println("Data  telah berhasil diubah")
		}
		i = i + 1
	}

	if !ketemu {
		fmt.Println("Meja yang dicari tidak ditemukan")
	}
}

// hapus data meja
func hapusMeja(T *tabMeja, n *int) {
	var nomor, i, idx int
	var ketemu bool

	nomor = inputAngka("Masukkan nomor meja yang ingin kamu hapus: ")

	ketemu = false
	idx = -1

	for i = 0; i < *n; i = i + 1 {
		if T[i].nomor == nomor {
			ketemu = true
			idx = i
		}
	}

	if ketemu {
		for i = idx; i < *n-1; i = i + 1 {
			T[i] = T[i+1]
		}

		*n = *n - 1

		fmt.Println("Data berhasil dihapus")
	} else {
		fmt.Println("Meja tidak ditemukan")
	}
}

func menuKelolaMeja(T *tabMeja, n *int) {
	var pilih int

	pilih = -1

	for pilih != 0 {

		fmt.Println("=================================")
		fmt.Println("         KELOLA MEJA")
		fmt.Println("=================================")
		fmt.Println("| 1 | Tambah Meja              |")
		fmt.Println("| 2 | Ubah Meja                |")
		fmt.Println("| 3 | Hapus Meja               |")
		fmt.Println("| 4 | Tampilkan Meja           |")
		fmt.Println("| 0 | Kembali                  |")
		fmt.Println("=================================")

		fmt.Print("Pilih menu : ")
		pilih = inputAngka("")

		switch pilih {

		case 1:
			tambahMeja(T, n)

		case 2:
			ubahMeja(T, *n)

		case 3:
			hapusMeja(T, n)

		case 4:
			tampilMeja(*T, *n)

		case 0:
			fmt.Println("Kembali ke menu utama...")

		default:
			fmt.Println("Pilihan tidak valid!")
		}

		fmt.Println()
	}
}

// nambah data pelanggan
func tambahPelanggan(P *tabPelanggan, n *int) {
	fmt.Print("ID Pelanggan : ")
	fmt.Scan(&P[*n].id)

	fmt.Print("Nama : ")
	fmt.Scan(&P[*n].nama)

	P[*n].noHP = inputNoHP("No HP : ")

	*n = *n + 1

	fmt.Println("Data pelanggan berhasil ditambahkan")
}

// menampilkan data pelanggan
func tampilPelanggan(P tabPelanggan, n int) {
	var i int

	fmt.Println("====================================================")
	fmt.Println("| ID     | Nama                 | No HP           |")
	fmt.Println("====================================================")

	for i = 0; i < n; i++ {
		fmt.Printf("| %-6s | %-20s | %-15s |\n",
			P[i].id,
			P[i].nama,
			P[i].noHP)
	}

	fmt.Println("====================================================")
}

// ubah data pelanggan (gavino)
func ubahPelanggan(P *tabPelanggan, n int) {
	var id string
	var i int
	var ketemu bool

	fmt.Print("Masukkan ID Pelanggan yang ingin diubah: ")
	fmt.Scan(&id)

	ketemu = false
	i = 0

	for i < n && !ketemu {

		if P[i].id == id {

			ketemu = true

			fmt.Print("ID Baru : ")
			fmt.Scan(&P[i].id)

			fmt.Print("Nama Baru : ")
			fmt.Scan(&P[i].nama)

			P[i].noHP = inputNoHP("No HP Baru : ")

			fmt.Println("Data pelanggan berhasil diubah")
		}

		i = i + 1
	}

	if !ketemu {
		fmt.Println("Data pelanggan tidak ditemukan")
	}
}

// hapus data pelanggan
func hapusPelanggan(P *tabPelanggan, n *int) {
	var id string
	var i, idx int
	var ketemu bool

	fmt.Print("Masukkan ID Pelanggan yang ingin dihapus: ")
	fmt.Scan(&id)

	ketemu = false
	idx = -1

	for i = 0; i < *n; i = i + 1 {

		if P[i].id == id {

			ketemu = true
			idx = i
		}
	}

	if ketemu {

		for i = idx; i < *n-1; i = i + 1 {
			P[i] = P[i+1]
		}

		*n = *n - 1

		fmt.Println("Data pelanggan berhasil dihapus")

	} else {

		fmt.Println("Data pelanggan tidak ditemukan")
	}
}

func menuKelolaPelanggan(P *tabPelanggan, n *int) {
	var pilih int

	pilih = -1

	for pilih != 0 {

		fmt.Println("=================================")
		fmt.Println("      KELOLA PELANGGAN")
		fmt.Println("=================================")
		fmt.Println("| 1 | Tambah Pelanggan         |")
		fmt.Println("| 2 | Ubah Pelanggan           |")
		fmt.Println("| 3 | Hapus Pelanggan          |")
		fmt.Println("| 4 | Tampil Pelanggan         |")
		fmt.Println("| 0 | Kembali                  |")
		fmt.Println("=================================")

		fmt.Print("Pilih menu : ")
		pilih = inputAngka("")

		switch pilih {

		case 1:
			tambahPelanggan(P, n)

		case 2:
			ubahPelanggan(P, *n)

		case 3:
			hapusPelanggan(P, n)

		case 4:
			tampilPelanggan(*P, *n)

		case 0:
			fmt.Println("Kembali ke menu utama...")

		default:
			fmt.Println("Pilihan tidak valid!")
		}

		fmt.Println()
	}
}

// tambah reservasi
func tambahReservasi(R *tabReservasi, n *int) {

	fmt.Print("ID Reservasi : ")
	fmt.Scan(&R[*n].id)

	fmt.Print("ID Pelanggan : ")
	fmt.Scan(&R[*n].idPelanggan)

	fmt.Print("ID Meja : ")
	fmt.Scan(&R[*n].idMeja)

	fmt.Print("Tanggal : ")
	fmt.Scan(&R[*n].tanggal)

	R[*n].jam = inputNoHP("Jam : ")

	*n = *n + 1

	fmt.Println("Reservasi berhasil ditambahkan")
}

// tampilkan reservasi
func tampilReservasi(R tabReservasi, n int) {
	var i int

	fmt.Println("======================================================================")
	fmt.Println("| ID Reservasi | ID Pelanggan | ID Meja | Tanggal    | Jam         |")
	fmt.Println("======================================================================")

	for i = 0; i < n; i++ {
		fmt.Printf("| %-12s | %-12s | %-7s | %-10s | %-11s |\n",
			R[i].id,
			R[i].idPelanggan,
			R[i].idMeja,
			R[i].tanggal,
			R[i].jam)
	}

	fmt.Println("======================================================================")
}

//menu reservasi
func menuReservasi(R *tabReservasi, n *int) {
	var pilih int

	pilih = -1

	for pilih != 0 {

		fmt.Println("=================================")
		fmt.Println("          RESERVASI")
		fmt.Println("=================================")
		fmt.Println("| 1 | Tambah Reservasi         |")
		fmt.Println("| 2 | Tampil Reservasi         |")
		fmt.Println("| 0 | Kembali                  |")
		fmt.Println("=================================")

		fmt.Print("Pilih menu : ")
		pilih = inputAngka("")

		switch pilih {

		case 1:
			tambahReservasi(R, n)

		case 2:
			tampilReservasi(*R, *n)

		case 0:
			fmt.Println("Kembali ke menu utama...")

		default:
			fmt.Println("Pilihan tidak valid!")
		}

		fmt.Println()
	}
}

// sequential search meja
func sequentialSearchMeja(T tabMeja, n int) {
	var nomor int
	var i int
	var ketemu bool

	nomor = inputAngka("Masukkan nomor meja yang dicari: ")

	ketemu = false
	i = 0

	for i < n && !ketemu {

		if T[i].nomor == nomor {

			ketemu = true

			fmt.Println("Data ditemukan")
			fmt.Println("ID :", T[i].id)
			fmt.Println("Nomor :", T[i].nomor)
			fmt.Println("Kapasitas :", T[i].kapasitas)
		}

		i = i + 1
	}

	if !ketemu {
		fmt.Println("Data tidak ditemukan")
	}
}

//binary search meja
func binarySearchMeja(T tabMeja, n int) {
	var nomor int
	var kiri, kanan, tengah int
	var ketemu bool

	nomor = inputAngka("Masukkan nomor meja yang dicari: ")

	kiri = 0
	kanan = n - 1
	ketemu = false

	for kiri <= kanan && !ketemu {

		tengah = (kiri + kanan) / 2

		if T[tengah].nomor == nomor {

			ketemu = true

			fmt.Println("Data ditemukan")
			fmt.Println("ID :", T[tengah].id)
			fmt.Println("Nomor :", T[tengah].nomor)
			fmt.Println("Kapasitas :", T[tengah].kapasitas)

		} else if nomor < T[tengah].nomor {

			kanan = tengah - 1

		} else {

			kiri = tengah + 1
		}
	}

	if !ketemu {
		fmt.Println("Data tidak ditemukan")
	}
}

// binary search kapasitas meja
func binarySearchKapasitas(T tabMeja, n int) {
	var kapasitas int
	var kiri, kanan, tengah int
	var ketemu bool

	kapasitas = inputAngka("Masukkan kapasitas yang dicari: ")

	kiri = 0
	kanan = n - 1
	ketemu = false

	for kiri <= kanan && !ketemu {

		tengah = (kiri + kanan) / 2

		if T[tengah].kapasitas == kapasitas {

			ketemu = true

			fmt.Println("Data ditemukan")
			fmt.Println("ID :", T[tengah].id)
			fmt.Println("Nomor :", T[tengah].nomor)
			fmt.Println("Kapasitas :", T[tengah].kapasitas)

		} else if kapasitas < T[tengah].kapasitas {

			kanan = tengah - 1

		} else {

			kiri = tengah + 1
		}
	}

	if !ketemu {
		fmt.Println("Data tidak ditemukan")
	}
}

//sequential kapasitas meja
func sequentialSearchKapasitas(T tabMeja, n int) {
	var kapasitas int
	var i int
	var ketemu bool

	kapasitas = inputAngka("Masukkan kapasitas yang dicari: ")

	ketemu = false
	i = 0

	for i < n && !ketemu {

		if T[i].kapasitas == kapasitas {

			ketemu = true

			fmt.Println("Data ditemukan")
			fmt.Println("ID :", T[i].id)
			fmt.Println("Nomor :", T[i].nomor)
			fmt.Println("Kapasitas :", T[i].kapasitas)
		}

		i = i + 1
	}

	if !ketemu {
		fmt.Println("Data tidak ditemukan")
	}
}

//menu pencarian
func menuPencarian(T *tabMeja, n int) {
	var pilihData, pilihMetode int

	fmt.Println("========================================")
	fmt.Println("             PENCARIAN")
	fmt.Println("========================================")
	fmt.Println("| 1 | Cari berdasarkan Nomor Meja      |")
	fmt.Println("| 2 | Cari berdasarkan Kapasitas       |")
	fmt.Println("| 0 | Kembali                          |")
	fmt.Println("========================================")
	fmt.Print("Pilih jenis pencarian : ")
	pilihData = inputAngka("")

	if pilihData != 0 {

		fmt.Println()
		fmt.Println("========================================")
		fmt.Println("          METODE PENCARIAN")
		fmt.Println("========================================")
		fmt.Println("| 1 | Sequential Search               |")
		fmt.Println("| 2 | Binary Search                   |")
		fmt.Println("========================================")
		fmt.Print("Pilih metode : ")
		pilihMetode = inputAngka("")

		if pilihData == 1 {

			if pilihMetode == 1 {
				sequentialSearchMeja(*T, n)
			} else if pilihMetode == 2 {
				insertionSortKapasitas(T, n)
				binarySearchMeja(*T, n)
			}

		} else if pilihData == 2 {

			if pilihMetode == 1 {
				sequentialSearchKapasitas(*T, n)
			} else if pilihMetode == 2 {
				selectionSortKapasitas(T, n)
				binarySearchKapasitas(*T, n)
			}

		}
	}
}

//selection kapasitas
func selectionSortKapasitas(T *tabMeja, n int) {
	var i, j, idxMin int
	var temp meja

	for i = 0; i < n-1; i++ {
		idxMin = i

		for j = i + 1; j < n; j++ {
			if T[j].kapasitas < T[idxMin].kapasitas {
				idxMin = j
			}
		}

		temp = T[i]
		T[i] = T[idxMin]
		T[idxMin] = temp
	}
}

//insertion kapasitas
func insertionSortKapasitas(T *tabMeja, n int) {
	var i, j int
	var temp meja

	for i = 1; i < n; i++ {
		temp = T[i]
		j = i - 1

		for j >= 0 && T[j].kapasitas > temp.kapasitas {
			T[j+1] = T[j]
			j--
		}

		T[j+1] = temp
	}
}

//menu pengurutan
func menuPengurutan(T *tabMeja, n int) {
	var pilih int

	pilih = -1

	for pilih != 0 {

		fmt.Println("========================================")
		fmt.Println("             PENGURUTAN")
		fmt.Println("========================================")
		fmt.Println("| 1 | Selection Sort Kapasitas         |")
		fmt.Println("| 2 | Insertion Sort Kapasitas         |")
		fmt.Println("| 0 | Kembali                          |")
		fmt.Println("========================================")

		fmt.Print("Pilih metode : ")
		pilih = inputAngka("")

		switch pilih {

		case 1:
			selectionSortKapasitas(T, n)
			fmt.Println()
			fmt.Println(">>> Data berhasil diurutkan dengan Selection Sort")
			tampilMeja(*T, n)

		case 2:
			insertionSortKapasitas(T, n)
			fmt.Println()
			fmt.Println(">>> Data berhasil diurutkan dengan Insertion Sort")
			tampilMeja(*T, n)

		case 0:
			fmt.Println("Kembali ke menu utama...")

		default:
			fmt.Println("Pilihan tidak valid!")
		}

		fmt.Println()
	}
}

//meja yang sering di pesan
func mejaTerfavorit(R tabReservasi, n int) string {
	var i, j, max, hitung int
	var idFavorit string

	max = 0

	for i = 0; i < n; i++ {

		hitung = 0

		for j = 0; j < n; j++ {

			if R[i].idMeja == R[j].idMeja {
				hitung++
			}
		}

		if hitung > max {
			max = hitung
			idFavorit = R[i].idMeja
		}
	}

	return idFavorit
}

//Statistik
func statistikReservasi(R tabReservasi, n int) {
	var i, j, jumlah int
	var sudahAda bool

	fmt.Println("=================================")
	fmt.Println("          STATISTIK")
	fmt.Println("=================================")

	fmt.Println("Jumlah Reservasi per Hari:")

	for i = 0; i < n; i++ {

		sudahAda = false

		for j = 0; j < i; j++ {

			if R[i].tanggal == R[j].tanggal {
				sudahAda = true
			}
		}

		if !sudahAda {

			jumlah = 0

			for j = 0; j < n; j++ {

				if R[i].tanggal == R[j].tanggal {
					jumlah = jumlah + 1
				}
			}

			fmt.Println(R[i].tanggal, ":", jumlah, "reservasi")
		}
	}

	fmt.Println()
	fmt.Println("Meja paling sering dipesan :", mejaTerfavorit(R, n))
}
