package usecase

import (
	"bufio"
	"encoding/json"
	"fmt"
	"miniproject2/models"
	"os"
	"strings"
	"sync"
	"time"
)

var listBook []models.DaftarBuku
var inputUser = bufio.NewReader(os.Stdin)

func TambahBuku() {
	var kodebuku string
	pengarang := ""
	penerbit := ""
	jumlahhalaman := 0
	tahunterbit := 0
	fmt.Println("=================================")
	fmt.Println("Tambah Pesanan")
	fmt.Println("=================================")

	draftBuku := []models.DaftarBuku{}

	for {
		for {
			fmt.Print("Silahkan Masukan Kode Buku : ")
			_, err := fmt.Scanln(&kodebuku)
			if err != nil {
				fmt.Println("Terjadi Error:", err)
				return
			}
			if kodeBukuExists(kodebuku) {
				fmt.Println("Kode buku sudah digunakan. Masukkan kode buku yang berbeda.")
			} else {
				break
			}
		}
		fmt.Print("Silahkan Masukan Buku : ")

		// untuk user Windows, gunakan yang dicomment (\r) :
		// menuPelanggan, err := inputanUser.ReadString('\r')

		judulbuku, err := inputUser.ReadString('\n')
		if err != nil {
			fmt.Println("Terjadi Error:", err)
			return
		}

		judulbuku = strings.Replace(
			judulbuku,
			"\n",
			"",
			1)

		// special treatment untuk windows
		judulbuku = strings.Replace(
			judulbuku,
			"\r",
			"",
			1)

		fmt.Print("Silahkan Masukan pengarang : ")
		_, err = fmt.Scanln(&pengarang)
		if err != nil {
			fmt.Println("Terjadi Error:", err)
			return
		}

		fmt.Print("Silahkan Masukan Penerbit : ")
		_, err = fmt.Scanln(&penerbit)
		if err != nil {
			fmt.Println("Terjadi Error:", err)
			return
		}
		fmt.Print("Silahkan Masukan Jumlah Halaman : ")
		_, err = fmt.Scanln(&jumlahhalaman)
		if err != nil {
			fmt.Println("Terjadi Error:", err)
			return
		}
		fmt.Print("Silahkan Masukan Tahun Terbit : ")
		_, err = fmt.Scanln(&tahunterbit)
		if err != nil {
			fmt.Println("Terjadi Error:", err)
			return
		}

		// Simpan ID dan Tanggal
		draftBuku = append(draftBuku, models.DaftarBuku{
			Kode_Buku:      kodebuku,
			Judul_Buku:     judulbuku,
			Pengarang:      pengarang,
			Penerbit:       penerbit,
			Jumlah_Halaman: jumlahhalaman,
			Tahun_Terbit:   tahunterbit,
			Tanggal:        time.Now(),
		})
		var pilMenuBuku = 0
		fmt.Println("Ketik 1 untuk tambah Buku, ketik 0 untuk keluar")
		_, err = fmt.Scanln(&pilMenuBuku)
		if err != nil {
			fmt.Println("Terjadi Error:", err)
			return
		}

		if pilMenuBuku == 0 {
			break
		}
	}

	fmt.Println("Menambah buku...")

	_ = os.Mkdir("books", 0777)

	ch := make(chan models.DaftarBuku)

	wg := sync.WaitGroup{}

	jumlahAntrian := 5

	// Mendaftarkan receiver/pemroses data
	for i := 0; i < jumlahAntrian; i++ {
		wg.Add(1)
		go simpanBuku(ch, &wg, i)
	}

	// Mengirimkan data ke channel
	for _, book := range draftBuku {
		ch <- book
	}

	close(ch)

	wg.Wait()

	fmt.Println("Berhasil Menambah Buku!")
}

func simpanBuku(ch <-chan models.DaftarBuku, wg *sync.WaitGroup, noAntrian int) {

	for buku := range ch {
		dataJson, err := json.Marshal(buku)
		if err != nil {
			fmt.Println("Terjadi error:", err)
		}

		err = os.WriteFile(fmt.Sprintf("books/%s.json", buku.Kode_Buku), dataJson, 0644)
		if err != nil {
			fmt.Println("Terjadi error:", err)
		}

		fmt.Printf("Antrian No %d Memproses Kode Buku : %s!\n", noAntrian, buku.Kode_Buku)
	}
	wg.Done()
}
