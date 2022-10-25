package main

import (
	"fmt"
	"os"
	"strconv"
)

type siswa struct {
	Name      string
	Alamat    string
	Pekerjaan string
	Alasan    string
}

func main() {

	xstr := os.Args[1]

	x, _ := strconv.Atoi(xstr)

	var dbSiswa = []siswa{
		{Name: "Surya", Alamat: "Eka Jaya", Pekerjaan: "Mahasiswa", Alasan: "Ingin tahu lebih banyak"},
		{Name: "Yuda", Alamat: "Paal Merah", Pekerjaan: "Mahasiswa", Alasan: "Supaya jadi programmer handal"},
		{Name: "Budi", Alamat: "Paal X", Pekerjaan: "Mahasiswa", Alasan: "Mau jadi backend engineer"},
		{Name: "Rehan", Alamat: "Paal V", Pekerjaan: "Mahasiswa", Alasan: "Modal buat cari kerja"},
		{Name: "Doni", Alamat: "Kasang", Pekerjaan: "PNS", Alasan: "Memperdalam ilmu coding"},
		{Name: "Rafli", Alamat: "Kumpeh", Pekerjaan: "Mahasiswa", Alasan: "Dapetin skill baru"},
		{Name: "Samsul", Alamat: "Handil", Pekerjaan: "Pengangguran", Alasan: "Nyari kesibukan"},
		{Name: "Bejo", Alamat: "Bandung barat", Pekerjaan: "Mahasiswa", Alasan: "Diajak teman"},
		{Name: "Tono", Alamat: "Solok", Pekerjaan: "Mahasiswa", Alasan: "Disuruh orang tua"},
		{Name: "Tedi", Alamat: "Citayem", Pekerjaan: "Freelancer", Alasan: "Dapetin sertifikat"},
	}

	if x > len(dbSiswa) || x < 1 {
		showError()
	} else {
		tampilData(dbSiswa[x-1])
	}
}

func tampilData(siswa siswa) {
	fmt.Println("Nama:", siswa.Name)
	fmt.Println("Alamat:", siswa.Alamat)
	fmt.Println("Pekerjaan: ", siswa.Pekerjaan)
	fmt.Println("Alasan:", siswa.Alasan)
}

func showError() {
	fmt.Println("No Absen Tidak Ditemukan")
}
