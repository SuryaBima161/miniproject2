package usecase

import (
	"fmt"
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
