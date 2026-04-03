package model

type Menu struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	Nama      string `gorm:"type:varchar(100)" json:"nama"`
	Harga     int    `json:"harga"`
	Kategori  string `gorm:"type:varchar(50)" json:"kategori"`
	Gambar    string `json:"gambar"`
	Deskripsi string `json:"deskripsi"`
}
