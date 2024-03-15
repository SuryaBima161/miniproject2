package models

import "time"

type DaftarBuku struct {
	Kode_Buku      string
	Judul_Buku     string
	Pengarang      string
	Penerbit       string
	Jumlah_Halaman int
	Tahun_Terbit   int
	Tanggal        time.Time
}
