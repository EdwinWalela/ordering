package handlers

import (
	"edwinwalela/ordering/models"

	"github.com/gin-gonic/gin"
)

func CreateCustomer(c *gin.Context) {
	var customer models.Customer
	if err := c.BindJSON(&customer); err != nil {
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

func GetCustomer(c *gin.Context) {
}

func CreateOrder(c *gin.Context) {
	var order models.Order
	if err := c.BindJSON(&order); err != nil {
		c.JSON(500, gin.H{
			"message": "Failed to create order",
			"error":   err.Error(),
		})
		return
	}
	c.JSON(201, gin.H{
		"message": "Order created",
		"order":   order,
	})
}

func GetOrder(c *gin.Context) {

}
