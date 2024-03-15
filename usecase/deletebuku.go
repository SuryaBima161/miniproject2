package usecase

import (
	"fmt"
	"os"
)

var kode int

func DeleteBuku() {

	fmt.Println("=================================")
	fmt.Println("Delete Buku")
	fmt.Println("=================================")
	ListBuku()
	fmt.Println("=================================")
	fmt.Print("Masukan urutan Buku : ")
	_, err := fmt.Scanln(&kode)
	if err != nil {
		fmt.Println("Terjadi error:", err)
		return
	}

	if (kode-1) < 0 ||
		(kode-1) > len(listBook) {
		fmt.Println("Urutan Buku Tidak Sesuai")
		DeleteBuku()
		return
	}
	fmt.Println("Buku dengan Kode", listBook[kode-1].Kode_Buku, " dengan judul : ", listBook[kode-1].Judul_Buku, " Berhasil Dihapus.")

	err = os.Remove(fmt.Sprintf("books/%s.json", listBook[kode-1].Kode_Buku))
	if err != nil {
		fmt.Println("Terjadi error:", err)
	}
}
