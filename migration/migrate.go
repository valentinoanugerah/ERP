package migration

import (
	_"github.com/valentinoanugerah/ERP/config"
	"github.com/valentinoanugerah/ERP/model"
	"log"
	"gorm.io/gorm"
)

func RunMigration(db *gorm.DB) {
	err := db.AutoMigrate(
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
