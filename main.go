package main

import (
	"log"

	"github.com/AnakonStar/go-api/database"
	"github.com/AnakonStar/go-api/router"
)

func main() {
	database.Connect() // Conecta ao MySQL

	r := router.SetupRoutes()

	log.Println("🚀 Servidor rodando em http://localhost:8080")
	r.Run(":8080")
}
