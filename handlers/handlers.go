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

	id, err := h.Repo.CreateCustomer(customer)

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
		"customer": id,
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
	orderId, err := h.Repo.CreateOrder(order)

	if err != nil {
		log.Printf("Failed to create order: %v", err)
		c.JSON(500, gin.H{
			"message": "Failed to create order",
			"error":   err.Error(),
		})
		return
	}

	customer := h.Repo.GetCustomerById(order.CustomerId)

	message := models.Message{
		Recipient: customer.Name,
		Item:      order.Item,
	}

	err = h.Repo.SmsService.SendMessage(message)
	fmt.Println(err)
	c.JSON(201, gin.H{
		"message": "Order created",
		"id":      orderId,
	})
}

func (h *Handlers) GetOrders(c *gin.Context) {
	orders, err := h.Repo.GetOrders()

	if err != nil {
		c.JSON(500, gin.H{
			"message": "Failed to create order",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"orders": orders,
	})
}
