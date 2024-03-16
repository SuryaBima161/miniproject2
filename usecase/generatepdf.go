package usecase

import (
	"fmt"
	"miniproject2/models"
	"time"

	"github.com/go-pdf/fpdf"
)

func GeneratePdf() {
	ListBuku()
	fmt.Println("=================================")
	fmt.Println("Membuat Daftar Buku ...")
	fmt.Println("=================================")
	pdf := fpdf.New("P", "mm", "A4", "")
	pdf.AddPage()

	pdf.SetFont("Arial", "", 12)
	pdf.SetLeftMargin(10)
	pdf.SetRightMargin(10)

	for i, buku := range listBook {
		bukuText := fmt.Sprintf(
			"buku  #%d:\nKodeBuku : %s\nJudulBuku : %s\nPengarang : %s\nPenerbit : %s\nJumlahHalaman : %d\nTahunTerbit :  %d\nTanggal : %s\n",
			i+1, buku.Kode_Buku, buku.Judul_Buku,
			buku.Pengarang, buku.Penerbit, buku.Jumlah_Halaman, buku.Tahun_Terbit,
			buku.Tanggal.Format("2006-01-02 15:04:05"))

		pdf.MultiCell(0, 10, bukuText, "0", "L", false)
		pdf.Ln(5)
	}

	err := pdf.OutputFileAndClose(
		fmt.Sprintf("daftar_buku _%s.pdf",
			time.Now().Format("2006-01-02-15-04-05")))

	if err != nil {
		fmt.Println("Terjadi error:", err)
	}
}

func GenerateSelected(selectedBook models.DaftarBuku) {
	fmt.Println("=================================")
	fmt.Println("Membuat Daftar Buku ...")
	fmt.Println("=================================")

	pdf := fpdf.New("P", "mm", "A4", "")
	pdf.AddPage()

	pdf.SetFont("Arial", "", 12)
	pdf.SetLeftMargin(10)
	pdf.SetRightMargin(10)

	bukuText := fmt.Sprintf(
		"KodeBuku : %s\nJudulBuku : %s\nPengarang : %s\nPenerbit : %s\nJumlahHalaman : %d\nTahunTerbit :  %d\nTanggal : %s\n",
		selectedBook.Kode_Buku, selectedBook.Judul_Buku,
		selectedBook.Pengarang, selectedBook.Penerbit, selectedBook.Jumlah_Halaman, selectedBook.Tahun_Terbit,
		selectedBook.Tanggal.Format("2006-01-02 15:04:05"))

	pdf.MultiCell(0, 10, bukuText, "0", "L", false)

	err := pdf.OutputFileAndClose(
		fmt.Sprintf("daftar_buku_%s.pdf",
			time.Now().Format("2006-01-02-15-04-05")))

	if err != nil {
		fmt.Println("Terjadi error:", err)
	}

}
func PrintSelectedBook() {
	ListBuku()

	// Prompt user to select a book
	fmt.Print("Masukkan nomor urut buku yang ingin dicetak: ")
	var selectedNumber int
	_, err := fmt.Scanln(&selectedNumber)
	if err != nil {
		fmt.Println("Terjadi error:", err)
		return
	}

	// Ensure the selected number is within the valid range
	if selectedNumber < 1 || selectedNumber > len(listBook) {
		fmt.Println("Nomor urut buku tidak valid.")
		return
	}

	// Get the selected book from the list
	selectedBook := listBook[selectedNumber-1]

	// Generate PDF for the selected book
	Selected(selectedBook)
}

func Selected(selectedBook models.DaftarBuku) {
	pdf := fpdf.New("P", "mm", "A4", "")
	pdf.AddPage()

	pdf.SetFont("Arial", "", 12)
	pdf.SetLeftMargin(10)
	pdf.SetRightMargin(10)

	bukuText := fmt.Sprintf(
		"KodeBuku : %s\nJudulBuku : %s\nPengarang : %s\nPenerbit : %s\nJumlahHalaman : %d\nTahunTerbit :  %d\nTanggal : %s\n",
		selectedBook.Kode_Buku, selectedBook.Judul_Buku,
		selectedBook.Pengarang, selectedBook.Penerbit, selectedBook.Jumlah_Halaman, selectedBook.Tahun_Terbit,
		selectedBook.Tanggal.Format("2006-01-02 15:04:05"))

	pdf.MultiCell(0, 10, bukuText, "0", "L", false)

	err := pdf.OutputFileAndClose(
		fmt.Sprintf("daftar_buku_%s.pdf",
			time.Now().Format("2006-01-02-15-04-05")))

	if err != nil {
		fmt.Println("Terjadi error:", err)
	}
	fmt.Println("Buku berhasil dicetak dalam file PDF.")
}
