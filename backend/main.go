package main

import (
	"log"

	"trader-backend/api"
	"trader-backend/services"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Инициализация Finam-клиента (пока заглушка)
	finamClient := services.NewFinamClient()

	// Подключаем маршруты
	api.RegisterRoutes(r, finamClient)

	log.Println("Server started on :8080")
	r.Run(":8080")
}

//мой коммит
