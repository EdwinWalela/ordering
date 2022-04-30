package repository

import (
	"edwinwalela/ordering/models"
	"log"
)

func (r *Repository) CreateCustomer(customer models.Customer) error {
	sqlStatement := `
	 INSERT INTO customers(name)
	 VALUES($1);
	`
	_, err := r.Conn.Exec(r.Ctx, sqlStatement, customer.Name)
	return err
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
