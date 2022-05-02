package repository

import (
	"edwinwalela/ordering/models"
	"log"
)

func (r *Repository) CreateCustomer(customer models.Customer) (int64, error) {
	sqlStatement := `
	 INSERT INTO customers(name)
	 VALUES($1)
	 returning id;
	`
	row := r.Conn.QueryRow(r.Ctx, sqlStatement, customer.Name)

	var id int64

	err := row.Scan(&id)

	return id, err
}

func (r *Repository) GetCustomerById(id int64) models.Customer {
	sqlStatement := `
	SELECT * FROM customers
	WHERE id=$1
`
	rows := r.Conn.QueryRow(r.Ctx, sqlStatement, id)

	var customer models.Customer
	rows.Scan(
		&customer.Id,
		&customer.Name,
		&customer.Code,
	)

	return customer
}

func (r *Repository) GetCustomers() ([]models.Customer, error) {
	sqlStatement := `
	SELECT * FROM customers
`
	rows, err := r.Conn.Query(r.Ctx, sqlStatement)

	if err != nil {
		log.Printf("Failed to query DB: %v", err)

		return nil, err
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
	return customers, nil
}
