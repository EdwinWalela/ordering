package handlers

import (
	"edwinwalela/ordering/models"
	r "edwinwalela/ordering/repository"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

type Handlers struct {
	Repo r.Repository
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

	err := h.Repo.CreateCustomer(customer)

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
	customers, err := h.Repo.GetCustomers()
	if err != nil {
		log.Printf("Failed to query DB: %v", err)
		c.JSON(500, gin.H{
			"message": "Error",
			"error":   err.Error(),
		})
		return
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
