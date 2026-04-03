package database

import "github.com/Nathanaelaxl/qresto-api/model"

func SeedMenus() {
	var count int64
	DB.Model(&model.Menu{}).Count(&count)

	// Jika tabel masih kosong, isi dengan data dummy
	if count == 0 {
		menus := []model.Menu{
			{Nama: "Ayam Goreng Penyet", Harga: 25000, Kategori: "Makanan", Gambar: "ayampenyet.png", Deskripsi: "Ayam gurih sambal pedas"},
			{Nama: "Nasi Goreng Special", Harga: 22000, Kategori: "Makanan", Gambar: "nasgor.png", Deskripsi: "Nasi goreng telur mata sapi"},
			{Nama: "Es Teh Manis", Harga: 5000, Kategori: "Minuman", Gambar: "esteh.png", Deskripsi: "Segar dan manis"},
		}
		DB.Create(&menus)
	}
}
