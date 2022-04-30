package repository

import (
	"edwinwalela/ordering/models"
	"log"
)

func (r *Repository) CreateOrder(order models.Order) (int64, error) {

	sqlStatement := `
	INSERT INTO orders (item,amount,customer_id)
	VALUES($1,$2,$3)
	returning id;
`
	var id int64
	err := r.Conn.QueryRow(r.Ctx, sqlStatement, order.Item, order.Amount, order.CustomerId).Scan(&id)
	return id, err
}

func (r *Repository) GetOrders() ([]models.Order, error) {
	var orders []models.Order
	sqlStatement := `
		SELECT * FROM orders
	`
	rows, err := r.Conn.Query(r.Ctx, sqlStatement)

	if err != nil {
		return nil, err
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
	return orders, nil
}
