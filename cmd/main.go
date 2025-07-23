package main

import (
	"log"

	"github.com/valentinoanugerah/ERP/config"
	"github.com/valentinoanugerah/ERP/migration"
)

func main() {
	_, err := config.ConnectDB()
	if err != nil {
		log.Fatalf("Database error: %v", err)
	}

	migration.RunMigration()

	log.Println("🚀 Aplikasi ERP berjalan...")
}
