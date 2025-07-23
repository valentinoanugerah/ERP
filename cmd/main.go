package main

import (
	"log"

	"github.com/valentinoanugerah/ERP/config"
)

func main() {
	// Inisialisasi koneksi database
	db, err := config.ConnectDB()
	if err != nil {
		log.Fatalf("Database error: %v", err)
	}
	defer db.Close()

	log.Println("✅ Koneksi ke database berhasil.")
	log.Println("🚀 Aplikasi ERP berjalan...")

	// TODO: Tambahkan inisialisasi router, migration, handler, dsb.
}
