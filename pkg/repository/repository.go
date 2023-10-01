package repository

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"
	"wb_l0"
)

type Repository struct {
	db    *sqlx.DB
	cache *Cache
}

func NewRepository(db *sqlx.DB) *Repository {
	orders := make(map[int]wb_l0.Order)
	cache := Cache{data: orders}
	return &Repository{db: db, cache: &cache}
}

func (r *Repository) GetOrderFromDB(id int) (wb_l0.Order, error) {
	query := fmt.Sprintf("SELECT data FROM orders WHERE id = $1")
	data := make([]byte, 10)
	err := r.db.Get(&data, query, id)
	if err != nil {
		return wb_l0.Order{}, err
	}
	var order wb_l0.Order
	if err = json.Unmarshal(data, &order); err != nil {
		return order, err
	}
	return order, nil
}

func (r *Repository) GetOrderFromCache(id int) (wb_l0.Order, error) {
	var order wb_l0.Order
	var exist bool
	order, exist = r.cache.data[id]
	if !exist {
		return order, errors.New("Order doesnt exist")
	}
	return order, nil
}

func (r *Repository) GetOrders() ([]wb_l0.Data, error) {
	var orders []wb_l0.Data
	query := "SELECT * FROM orders"
	rows, err := r.db.Query(query)
	if err != nil {
		return orders, err
	}

	for rows.Next() {
		var id int
		var order wb_l0.Data
		data := make([]byte, 10)
		if err = rows.Scan(&id, &data); err != nil {
			return nil, err
		}
		order.Id = id
		err = json.Unmarshal(data, &order.Order)
		if err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}
	return orders, nil
}

func (r *Repository) PushOrder(order wb_l0.Order) (int, error) {
	var id int
	query := "INSERT INTO orders(data) VALUES($1) RETURNING id"
	row := r.db.QueryRow(query, order)
	err := row.Scan(&id)

	if err != nil {
		return 0, err
	}

	r.cache.Set(wb_l0.Data{Id: id, Order: order})

	return id, nil
}

func (r *Repository) RestoreCache() {
	data, err := r.GetOrders()
	if err != nil {
		log.Fatal("Couldnt restore cache:", err.Error())
	}
	for i := range data {
		r.cache.data[data[i].Id] = data[i].Order
	}
}
