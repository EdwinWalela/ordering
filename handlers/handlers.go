package handlers

import (
	"context"
	"edwinwalela/ordering/models"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
)

type Handlers struct {
	Conn *pgx.Conn
	Ctx  context.Context
}

func (h *Handlers) CreateCustomer(c *gin.Context) {
	var customer models.Customer
	if err := c.BindJSON(&customer); err != nil {
		c.JSON(500, gin.H{
			"message": "Failed to create customer",
			"error":   err.Error(),
		})
		return
	}

	sqlStatement := `
	 INSERT INTO customers(name)
	 VALUES($1);
	`
	_, err := h.Conn.Exec(h.Ctx, sqlStatement, customer.Name)
	if err != nil {
		fmt.Printf("Failed to insert customer to DB: %v", err)
		c.JSON(500, gin.H{
			"message": "Failed to create customer",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(201, gin.H{
		"message":  "Customer created",
		"customer": customer,
	})
}

func (h *Handlers) GetCustomers(c *gin.Context) {
	sqlStatement := `
		SELECT * FROM customers
	`
	rows, err := h.Conn.Query(h.Ctx, sqlStatement)
	if err != nil {
		log.Printf("Failed to query DB: %v", err)
		c.JSON(500, gin.H{
			"message": "Error",
			"error":   err.Error(),
		})
		return
	}

	var customers []models.Customer

	for rows.Next() {
		var customer models.Customer
		rows.Scan(
			&customer.Id,
			&customer.Name,
			&customer.Code,
		)
		customers = append(customers, customer)
	}
	c.JSON(200, gin.H{
		"customers": customers,
	})
}

func (h *Handlers) CreateOrder(c *gin.Context) {
	var order models.Order
	if err := c.BindJSON(&order); err != nil {
		c.JSON(500, gin.H{
			"message": "Failed to create order",
			"error":   err.Error(),
		})
		return
	}

	sqlStatement := `
		INSERT INTO orders (item,amount,customer_id)
		VALUES($1,$2,$3)
		returning id;
	`

	var id int64
	err := h.Conn.QueryRow(h.Ctx, sqlStatement, order.Item, order.Amount, order.CustomerId).Scan(&id)

	if err != nil {
		log.Printf("Failed to create order: %v", err)
		c.JSON(500, gin.H{
			"message": "Failed to create order",
			"error":   err.Error(),
		})
		return
	}
	// TODO: Send sms via AT
	c.JSON(201, gin.H{
		"message": "Order created",
		"id":      id,
	})
}

func (h *Handlers) GetOrders(c *gin.Context) {
	var orders []models.Order
	sqlStatement := `
		SELECT * FROM orders
	`
	rows, err := h.Conn.Query(h.Ctx, sqlStatement)
	if err != nil {
		log.Printf("Failed to query DB: %v", err)
		c.JSON(500, gin.H{
			"message": "Error",
			"error":   err.Error(),
		})
		return
	}

	for rows.Next() {
		var order models.Order

		err := rows.Scan(
			&order.Id,
			&order.Item,
			&order.Amount,
			&order.CustomerId,
			&order.Time,
		)

		log.Println(err)

		orders = append(orders, order)
	}

	c.JSON(200, gin.H{
		"orders": orders,
	})
}
