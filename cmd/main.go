package main

import (
	"log"
	"github.com/valentinoanugerah/ERP/config"
	"github.com/valentinoanugerah/ERP/migration"
)

func main() {
	db, err := config.ConnectDB()
	if err != nil {
		log.Fatalf("Database error: %v", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Gagal mengambil koneksi SQL: %v", err)
	}
	defer sqlDB.Close()

	migration.RunMigration(db)

	log.Println("✅ Koneksi ke database berhasil.")
	log.Println("🚀 Aplikasi ERP berjalan...")
}
