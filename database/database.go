package database

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	db, err := gorm.Open(sqlite.Open("app.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("❌ Erro ao conectar ao banco:", err)
	}

	DB = db
	log.Println("✅ Banco SQLite conectado com sucesso")
}
