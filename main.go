package main

import (
	"log"

	c "edwinwalela/ordering/config"
	"edwinwalela/ordering/handlers"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Failed to load .env file: %v", err)
	}

	cfg := c.LoadConfig()

	r := gin.Default()

	r.POST("/customers", handlers.CreateCustomer)
	r.POST("/orders", handlers.CreateOrder)
	r.GET("/customers", handlers.GetCustomer)
	r.GET("/orders", handlers.GetOrder)

	r.Run(":" + cfg.Port)

}
