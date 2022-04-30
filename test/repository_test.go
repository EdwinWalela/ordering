package test

import (
	"edwinwalela/ordering/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetCustomers(t *testing.T) {
	customers, err := r.GetCustomers()
	if err != nil {
		t.Fatalf("Failed getting customers: %v", err)
	}
	assert.GreaterOrEqual(t, len(customers), 1)
}
func TestGetOrders(t *testing.T) {
	orders, err := r.GetOrders()
	if err != nil {
		t.Fatalf("Failed getting orders: %v", err)
	}
	assert.GreaterOrEqual(t, len(orders), 1)
}
func TestCreateCustomer(t *testing.T) {
	customer := models.Customer{Name: "test-customer"}
	_, err := r.CreateCustomer(customer)
	if err != nil {
		t.Fatalf("Failed to create customer: %v", err)
	}
	customers, err := r.GetCustomers()
	if err != nil {
		t.Fatalf("Failed to retrive customers: %v", err)
	}
	exists := false
	for _, cust := range customers {
		if cust.Name == customer.Name {
			exists = true
			break
		}
	}
	assert.Equal(t, true, exists)
}
func TestCreaterOrder(t *testing.T) {
	newOrder := models.Order{
		Item:       "test-item",
		Amount:     500,
		CustomerId: 1,
	}
	_, err := r.CreateOrder(newOrder)
	if err != nil {
		t.Fatalf("Failed to create customer: %v", err)
	}
	orders, err := r.GetOrders()
	if err != nil {
		t.Fatalf("Failed to retrive customers: %v", err)
	}
	exists := false
	for _, order := range orders {
		if order.Item == newOrder.Item && order.Amount == newOrder.Amount && order.CustomerId == newOrder.CustomerId {
			{
				exists = true
				break
			}
		}
	}
	assert.Equal(t, true, exists)
}
