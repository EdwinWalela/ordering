package handlers

import (
	"io/ioutil"

	"github.com/gin-gonic/gin"
)

func CreateCustomer(c *gin.Context) {
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "Failed to create customer",
			"error":   err.Error(),
		})
	}
	c.JSON(201, gin.H{
		"message": "Customer created",
		"body":    body,
	})
}

func GetCustomer(c *gin.Context) {

}

func CreateOrder(c *gin.Context) {

}

func GetOrder(c *gin.Context) {

}
