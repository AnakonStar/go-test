package mysql

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	// Carrega o arquivo .env
	err := godotenv.Load()
	if err != nil {
		log.Println("⚠️ .env não encontrado, usando variáveis de ambiente do sistema")
	}

	// Lê as variáveis
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	maxIdle := 10
	maxOpen := 100

	// DSN do MySQL (Data Source Name)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user, password, host, port, dbname)

	// Conecta com GORM
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("❌ Erro ao conectar ao MySQL:", err)
	}

	sqlDB, _ := DB.DB()
	sqlDB.SetMaxIdleConns(maxIdle)
	sqlDB.SetMaxOpenConns(maxOpen)

	DB = db
	log.Println("✅ Conectado ao MySQL com sucesso")
}
