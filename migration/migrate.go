package migration

import (
	"log"

	"github.com/valentinoanugerah/ERP/config"
	"github.com/valentinoanugerah/ERP/model"
)

func RunMigration() {
	db, err := config.ConnectDB()
	if err != nil {
		log.Fatal("Gagal koneksi ke database:", err)
	}

	err = db.AutoMigrate(
		&model.AuditLog{},
		&model.Customer{},
		&model.Product{},
		&model.PurchaseItem{},
		&model.Purchase{},
		&model.Role{},
		&model.SaleItem{},
		&model.Sale{},
		&model.Supplier{},
		&model.Transaction{},
		&model.User{},
	)

	if err != nil {
		log.Fatal("Migrasi gagal:", err)
	}

	log.Println("Migrasi berhasil dijalankan!")
}
