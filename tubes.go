package main

import (
	"fmt"
	"sort"
	"time"
)

type PaketMCU struct {
	ID     int
	Nama   string
	Harga  int
	Durasi time.Duration
}

type Pasien struct {
	ID         int
	Nama       string
	TanggalMCU time.Time
	PaketMCU   *PaketMCU
}

type RekapMCU struct {
	TanggalMCU     time.Time
	TotalPemasukan int
}

var dataPaketMCU = make(map[int]*PaketMCU)
var dataPasien = make(map[int]*Pasien)
var rekapMCU = make(map[time.Time]*RekapMCU)

func mainn() {
	for {
		fmt.Println("========== Aplikasi Pengelolaan Data MCU ==========")
		fmt.Println("1. Tambah Paket MCU")
		fmt.Println("2. Edit Paket MCU")
		fmt.Println("3. Hapus Paket MCU")
		fmt.Println("4. Tambah Data Pasien")
		fmt.Println("5. Edit Data Pasien")
		fmt.Println("6. Hapus Data Pasien")
		fmt.Println("7. Laporan Pemasukan")
		fmt.Println("8. Cari Pasien berdasarkan Paket")
		fmt.Println("9. Cari Pasien berdasarkan Tanggal")
		fmt.Println("10. Cari Pasien berdasarkan Nama")
		fmt.Println("11. Keluar")
		fmt.Println("===================================================")
		fmt.Print("Pilihan Anda: ")

		var pilihan int
		fmt.Scanln(&pilihan)

		switch pilihan {
		case 1:
			tambahPaketMCU()
		case 2:
			editPaketMCU()
		case 3:
			hapusPaketMCU()
		case 4:
			tambahDataPasien()
		case 5:
			editDataPasien()
		case 6:
			hapusDataPasien()
		case 7:
			laporanPemasukan()
		case 8:
			cariPasienBerdasarkanPaket()
		case 9:
			cariPasienBerdasarkanTanggal()
		case 10:
			cariPasienBerdasarkanNama()
		case 11:
			fmt.Println("Terima kasih! Sampai jumpa lagi.")
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

func tambahPaketMCU() {
	fmt.Println("===== Tambah Paket MCU =====")
	var paket PaketMCU

	fmt.Print("ID Paket: ")
	fmt.Scanln(&paket.ID)

	fmt.Print("Nama Paket: ")
	fmt.Scanln(&paket.Nama)

	fmt.Print("Harga Paket: ")
	fmt.Scanln(&paket.Harga)

	fmt.Print("Durasi (dalam menit): ")
	var durasi int
	fmt.Scanln(&durasi)
	paket.Durasi = time.Duration(durasi) * time.Minute

	dataPaketMCU[paket.ID] = &paket

	fmt.Println("Paket MCU berhasil ditambahkan.")
}

func editPaketMCU() {
	fmt.Println("===== Edit Paket MCU =====")
	var id int
	fmt.Print("Masukkan ID Paket yang akan diubah: ")
	fmt.Scanln(&id)

	paket, ok := dataPaketMCU[id]
	if !ok {
		fmt.Println("ID Paket tidak ditemukan.")
		return
	}

	fmt.Printf("Nama Paket (sebelumnya: %s): ", paket.Nama)
	fmt.Scanln(&paket.Nama)

	fmt.Printf("Harga Paket (sebelumnya: %d): ", paket.Harga)
	fmt.Scanln(&paket.Harga)

	fmt.Printf("Durasi (dalam menit) (sebelumnya: %d): ", int(paket.Durasi.Minutes()))
	var durasi int
	fmt.Scanln(&durasi)
	paket.Durasi = time.Duration(durasi) * time.Minute

	fmt.Println("Paket MCU berhasil diubah.")
}

func hapusPaketMCU() {
	fmt.Println("===== Hapus Paket MCU =====")
	var id int
	fmt.Print("Masukkan ID Paket yang akan dihapus: ")
	fmt.Scanln(&id)

	_, ok := dataPaketMCU[id]
	if !ok {
		fmt.Println("ID Paket tidak ditemukan.")
		return
	}

	delete(dataPaketMCU, id)

	fmt.Println("Paket MCU berhasil dihapus.")
}

func tambahDataPasien() {
	fmt.Println("===== Tambah Data Pasien =====")
	var pasien Pasien

	fmt.Print("ID Pasien: ")
	fmt.Scanln(&pasien.ID)

	fmt.Print("Nama Pasien: ")
	fmt.Scanln(&pasien.Nama)

	fmt.Print("Tanggal MCU (dd-mm-yyyy): ")
	var tanggal string
	fmt.Scanln(&tanggal)
	pasien.TanggalMCU, _ = time.Parse("02-01-2006", tanggal)

	fmt.Print("Pilih ID Paket MCU: ")
	var idPaket int
	fmt.Scanln(&idPaket)
	paket, ok := dataPaketMCU[idPaket]
	if !ok {
		fmt.Println("ID Paket tidak ditemukan.")
		return
	}
	pasien.PaketMCU = paket

	dataPasien[pasien.ID] = &pasien

	fmt.Println("Data pasien berhasil ditambahkan.")
}

func editDataPasien() {
	fmt.Println("===== Edit Data Pasien =====")
	var id int
	fmt.Print("Masukkan ID Pasien yang akan diubah: ")
	fmt.Scanln(&id)

	pasien, ok := dataPasien[id]
	if !ok {
		fmt.Println("ID Pasien tidak ditemukan.")
		return
	}

	fmt.Printf("Nama Pasien (sebelumnya: %s): ", pasien.Nama)
	fmt.Scanln(&pasien.Nama)

	fmt.Printf("Tanggal MCU (dd-mm-yyyy) (sebelumnya: %s): ", pasien.TanggalMCU.Format("02-01-2006"))
	var tanggal string
	fmt.Scanln(&tanggal)
	pasien.TanggalMCU, _ = time.Parse("02-01-2006", tanggal)

	fmt.Printf("Pilih ID Paket MCU (sebelumnya: %d): ", pasien.PaketMCU.ID)
	var idPaket int
	fmt.Scanln(&idPaket)
	paket, ok := dataPaketMCU[idPaket]
	if !ok {
		fmt.Println("ID Paket tidak ditemukan.")
		return
	}
	pasien.PaketMCU = paket

	fmt.Println("Data pasien berhasil diubah.")
}

func hapusDataPasien() {
	fmt.Println("===== Hapus Data Pasien =====")
	var id int
	fmt.Print("Masukkan ID Pasien yang akan dihapus: ")
	fmt.Scanln(&id)

	_, ok := dataPasien[id]
	if !ok {
		fmt.Println("ID Pasien tidak ditemukan.")
		return
	}

	delete(dataPasien, id)

	fmt.Println("Data pasien berhasil dihapus.")
}

func laporanPemasukan() {
	fmt.Println("===== Laporan Pemasukan =====")
	var tanggalAwal, tanggalAkhir time.Time

	fmt.Print("Tanggal Awal (dd-mm-yyyy): ")
	var tanggal string
	fmt.Scanln(&tanggal)
	tanggalAwal, _ = time.Parse("02-01-2006", tanggal)

	fmt.Print("Tanggal Akhir (dd-mm-yyyy): ")
	fmt.Scanln(&tanggal)
	tanggalAkhir, _ = time.Parse("02-01-2006", tanggal)

	totalPemasukan := 0
	for _, rekap := range rekapMCU {
		if rekap.TanggalMCU.After(tanggalAwal) && rekap.TanggalMCU.Before(tanggalAkhir) {
			totalPemasukan += rekap.TotalPemasukan
		}
	}

	fmt.Printf("Total Pemasukan: %d\n", totalPemasukan)
}

func cariPasienBerdasarkanPaket() {
	fmt.Println("===== Cari Pasien berdasarkan Paket =====")
	var idPaket int
	fmt.Print("Masukkan ID Paket: ")
	fmt.Scanln(&idPaket)

	paket, ok := dataPaketMCU[idPaket]
	if !ok {
		fmt.Println("ID Paket tidak ditemukan.")
		return
	}

	var pasienMCU []*Pasien
	for _, pasien := range dataPasien {
		if pasien.PaketMCU.ID == idPaket {
			pasienMCU = append(pasienMCU, pasien)
		}
	}

	sort.Slice(pasienMCU, func(i, j int) bool {
		return pasienMCU[i].TanggalMCU.Before(pasienMCU[j].TanggalMCU)
	})

	fmt.Printf("Daftar Pasien dengan Paket %s:\n", paket.Nama)
	for _, pasien := range pasienMCU {
		fmt.Printf("ID: %d, Nama: %s, Tanggal MCU: %s\n", pasien.ID, pasien.Nama, pasien.TanggalMCU.Format("02-01-2006"))
	}
}

func cariPasienBerdasarkanTanggal() {
	fmt.Println("===== Cari Pasien berdasarkan Tanggal =====")
	fmt.Print("Tanggal MCU (dd-mm-yyyy): ")
	var tanggal string
	fmt.Scanln(&tanggal)

	tanggalMCU, _ := time.Parse("02-01-2006", tanggal)

	var pasienMCU []*Pasien
	for _, pasien := range dataPasien {
		if pasien.TanggalMCU.Equal(tanggalMCU) {
			pasienMCU = append(pasienMCU, pasien)
		}
	}

	sort.Slice(pasienMCU, func(i, j int) bool {
		return pasienMCU[i].TanggalMCU.Before(pasienMCU[j].TanggalMCU)
	})

	fmt.Printf("Daftar Pasien dengan Tanggal MCU %s:\n", tanggalMCU.Format("02-01-2006"))
	for _, pasien := range pasienMCU {
		fmt.Printf("ID: %d, Nama: %s, Paket: %s\n", pasien.ID, pasien.Nama, pasien.PaketMCU.Nama)
	}
}

func cariPasienBerdasarkanNama() {
	fmt.Println("===== Cari Pasien berdasarkan Nama =====")
	fmt.Print("Nama Pasien: ")
	var nama string
	fmt.Scanln(&nama)

	var pasienMCU []*Pasien
	for _, pasien := range dataPasien {
		if pasien.Nama == nama {
			pasienMCU = append(pasienMCU, pasien)
		}
	}

	sort.Slice(pasienMCU, func(i, j int) bool {
		return pasienMCU[i].TanggalMCU.Before(pasienMCU[j].TanggalMCU)
	})

	fmt.Printf("Daftar Pasien dengan Nama %s:\n", nama)
	for _, pasien := range pasienMCU {
		fmt.Printf("ID: %d, Tanggal MCU: %s, Paket: %s\n", pasien.ID, pasien.TanggalMCU.Format("02-01-2006"), pasien.PaketMCU.Nama)
	}
}
