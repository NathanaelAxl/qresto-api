package database

import "github.com/Nathanaelaxl/qresto-api/model"

func RunMigrations() {
	// Menjalankan auto migrate untuk model Menu
	DB.AutoMigrate(&model.Menu{}, &model.Order{}, &model.OrderDetail{})
}
