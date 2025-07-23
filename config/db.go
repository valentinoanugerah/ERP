package config

import (
	"github.com/jmoiron/sqlx"
	_"github.com/lib/pq"
	"fmt"
	"log"
)

var DB *sqlx.DB

func connect(){
	db, err := sqlx.Connect("postgres", "user=valentinoanugerah dbname=erp sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Koneksi berhasil!")
	defer db.Close()
}