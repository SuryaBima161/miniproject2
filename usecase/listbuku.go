package usecase

import (
	"encoding/json"
	"fmt"
	"miniproject2/models"
	"os"
	"sort"
	"sync"
)

// func ListBuku() {
// 	fmt.Println("List Book")
// 	for index, list := range listBook {
// 		fmt.Printf("%d. Kode Buku: %s Judul Buku: %s Pengarang: %s Penerbit: %s Jumlah Halaman: %d : Tahun Terbit: %d\n",
// 			index+1, list.Kode_Buku, list.Judul_Buku, list.Pengarang, list.Penerbit, list.Jumlah_Halaman, list.Tahun_Terbit)
// 	}
// }

func LihatBuku(ch <-chan string, chPesanan chan models.DaftarBuku, wg *sync.WaitGroup) {
	var listBuku models.DaftarBuku
	for idBuku := range ch {
		dataJSON, err := os.ReadFile(fmt.Sprintf("Books/%s", idBuku))
		if err != nil {
			fmt.Println("Terjadi error:", err)
		}

		err = json.Unmarshal(dataJSON, &listBuku)
		if err != nil {
			fmt.Println("Terjadi error:", err)
		}

		chPesanan <- listBuku
	}
	wg.Done()
}

func ListBuku() {
	fmt.Println("Lihat List Buku")
	fmt.Println("=================================")
	fmt.Println("Memuat data ...")
	listBook = []models.DaftarBuku{}

	listJsonBuku, err := os.ReadDir("Books")
	if err != nil {
		fmt.Println("Terjadi error: ", err)
	}

	wg := sync.WaitGroup{}

	ch := make(chan string)
	chPesanan := make(chan models.DaftarBuku, len(listJsonBuku))

	jumlahPelayan := 5

	for i := 0; i < jumlahPelayan; i++ {
		wg.Add(1)
		go LihatBuku(ch, chPesanan, &wg)
	}

	for _, filePesanan := range listJsonBuku {
		ch <- filePesanan.Name()
	}

	close(ch)

	wg.Wait()

	close(chPesanan)

	for dataPesanan := range chPesanan {
		listBook = append(listBook, dataPesanan)
	}

	// kita urutkan list Pesanan sesuai waktu dibuat
	sort.Slice(listBook, func(i, j int) bool {
		return listBook[i].Tanggal.Before(listBook[j].Tanggal)
	})

	for urutan, book := range listBook {
		fmt.Printf("%d. Kode Buku : %s, Judul Buku : %s, Pengarang : %s, Penerbit : %s, Jumlah Halaman : %d, Tahun Terbit : %d\n",
			urutan+1,
			book.Kode_Buku,
			book.Judul_Buku,
			book.Pengarang,
			book.Penerbit,
			book.Jumlah_Halaman,
			book.Tahun_Terbit,
		)
	}

}
