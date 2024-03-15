package main

import (
	"fmt"
	"miniproject2/usecase"
	"os"
)

func main() {

	chooseOrder := 0
	fmt.Println("Daftar buku manajemen")
	fmt.Println("================================================")
	fmt.Println("Silahkan Pilih :")
	fmt.Println("1. Tambah Buku")
	fmt.Println("2. Edit Buku")
	fmt.Println("3. Delete Buku")
	fmt.Println("4. List Buku")
	fmt.Println("5. Generate PDF")
	fmt.Println("6. Out Program")
	fmt.Println("Tekan pilihanmu")
	_, err := fmt.Scanln(&chooseOrder)

	if err != nil {
		fmt.Println("error: ", err)
	}

	if chooseOrder == 1 {
		usecase.TambahBuku()
	} else if chooseOrder == 2 {
		usecase.EditBuku()
	} else if chooseOrder == 3 {
		usecase.DeleteBuku()
	} else if chooseOrder == 4 {
		usecase.ListBuku()
	} else if chooseOrder == 5 {
		usecase.GeneratePdf()
	} else if chooseOrder == 6 {
		os.Exit(0)
	}
	main()
}
