package main

import (
	"context"
	"log"

	c "edwinwalela/ordering/config"
	"edwinwalela/ordering/handlers"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Failed to load .env file: %v", err)
	}

	cfg := c.LoadConfig()
	r := gin.Default()

	conn, err := pgx.Connect(context.Background(), cfg.DbURl)
	if err != nil {
		log.Fatalf("Failed to connect to DB: %v", err)
	}
	log.Println("Connected to DB")
	handlers := handlers.Handlers{
		Conn: conn,
		Ctx:  context.Background(),
	}

	r.POST("/customers", handlers.CreateCustomer)
	r.POST("/orders", handlers.CreateOrder)
	r.GET("/customers", handlers.GetCustomers)
	r.GET("/orders", handlers.GetOrder)

	r.Run(":" + cfg.Port)

}
