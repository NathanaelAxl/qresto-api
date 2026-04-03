package main

import (
	"log"

	"github.com/Nathanaelaxl/qresto-api/controller"
	"github.com/Nathanaelaxl/qresto-api/database"
	"github.com/Nathanaelaxl/qresto-api/repository"
	"github.com/Nathanaelaxl/qresto-api/service"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// 1. Load file .env untuk konfigurasi database
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// 2. Jalankan Database & Automasi Tabel
	database.ConnectDatabase()
	database.RunMigrations()
	database.SeedMenus() // Mengisi data menu awal jika tabel kosong

	r := gin.Default()

	// 3. Setup CORS (Agar React di port 5173 bisa mengakses API ini)
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	// 4. Inisialisasi Layer Menu (Dependency Injection)
	menuRepo := repository.NewMenuRepository(database.DB)
	menuSvc := service.NewMenuService(menuRepo)
	menuCtrl := controller.NewMenuController(menuSvc)

	// 5. Inisialisasi Layer Order (Dependency Injection)
	orderRepo := repository.NewOrderRepository(database.DB)
	orderSvc := service.NewOrderService(orderRepo)
	orderCtrl := controller.NewOrderController(orderSvc)

	// 6. Route Grouping
	api := r.Group("/api")
	{
		// Endpoint untuk mengambil daftar menu
		api.GET("/menus", menuCtrl.GetMenus)

		// Endpoint untuk mengirim pesanan baru
		api.POST("/orders", orderCtrl.CreateOrder)
	}

	// Menjalankan server di port 8080
	r.Run(":8080")
}
