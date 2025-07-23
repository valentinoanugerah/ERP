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
	defer func() {
		sqlDB, _ := db.DB()
		sqlDB.Close()
	}()

	log.Println("✅ Koneksi ke database berhasil.")

	migration.RunMigration(db)

	log.Println("🚀 Aplikasi ERP berjalan...")
}
