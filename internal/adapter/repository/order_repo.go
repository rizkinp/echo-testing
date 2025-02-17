package repository

import (
	"database/sql"
	"my-echo-app/internal/core/model"
	"my-echo-app/internal/interfaces/repository" // Import yang benar
)

// NewOrderRepository mengembalikan `OrderRepository` sebagai interface
func NewOrderRepository(db *sql.DB) repository.OrderRepository {
	return &orderRepository{db: db}
}

type orderRepository struct {
	db *sql.DB
}

// CreateOrder implements repository.OrderRepository.
func (o *orderRepository) CreateOrder(order model.Order) (model.Order, error) {
	query := "INSERT INTO orders_staging (order_id, product_code, product_name, quantity, order_date, price, customer_id) VALUES (?, ?, ?, ?, ?, ?, ?)"
	_, err := o.db.Exec(query, order.OrderID, order.ProductCode, order.ProductName, order.Quantity, order.OrderDate, order.Price, order.CustomerID)
	if err != nil {
		return model.Order{}, err
	}
	return order, nil
}

func (o *orderRepository) GetAllOrders() ([]model.Order, error) {
	query := "SELECT * FROM orders_staging"
	rows, err := o.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []model.Order
	for rows.Next() {
		var order model.Order
		if err := rows.Scan(&order.OrderID, &order.ProductCode, &order.ProductName, &order.Quantity, &order.OrderDate, &order.Price, &order.CustomerID); err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}

	return orders, nil
}

func (o *orderRepository) GetOrder(id string) (model.Order, error) {
	query := "SELECT * FROM orders_staging WHERE order_id = ?"
	row := o.db.QueryRow(query, id)
	var order model.Order
	if err := row.Scan(&order.OrderID, &order.ProductCode, &order.ProductName, &order.Quantity, &order.OrderDate, &order.Price, &order.CustomerID); err != nil {
		return model.Order{}, err
	}
	return order, nil
}
