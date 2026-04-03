package model

import "time"

type Order struct {
	ID         uint          `gorm:"primaryKey" json:"id"`
	NomorMeja  string        `json:"nomorMeja"`
	TotalHarga int           `json:"totalHarga"`
	Status     string        `json:"status"` // "Pending", "Dimasak", "Selesai"
	CreatedAt  time.Time     `json:"createdAt"`
	Items      []OrderDetail `gorm:"foreignKey:OrderID" json:"items"`
}

type OrderDetail struct {
	ID       uint   `gorm:"primaryKey" json:"id"`
	OrderID  uint   `json:"orderId"`
	MenuID   uint   `json:"menuId"`
	NamaMenu string `json:"namaMenu"`
	Qty      int    `json:"qty"`
	Harga    int    `json:"harga"`
	Note     string `json:"note"` // Gin akan otomatis memetakan JSON "note" ke sini
}
