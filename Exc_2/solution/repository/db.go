package repository

import (
	"ordersystem/model"
	"time"
)

type DatabaseHandler struct {
	// drinks represent all available drinks
	drinks []model.Drink
	// orders serves as order history
	orders []model.Order
}

// todo
func NewDatabaseHandler() *DatabaseHandler {
	// Init the drinks slice with some test data
	// drinks := ...
	drinks := []model.Drink{
		{ID: 1, Name: "Water", Price: 1, Description: "Still mineral water"},
		{ID: 2, Name: "Espresso", Price: 2, Description: "Just espresso"},
		{ID: 3, Name: "Capuccino", Price: 4, Description: "Espresso with milk and foam"},
		{ID: 4, Name: "Cafe au lait", Price: 3, Description: "Espresso with milk"},
		{ID: 5, Name: "Chocolate", Price: 3, Description: "Just chocolate"},
		{ID: 6, Name: "Tea", Price: 2., Description: "Just tea"},
	}

	// Init orders slice with some test data
	orders := []model.Order{
		{Amount: 2, CreatedAt: time.Date(2025, 10, 14, 9, 5, 0, 0, time.UTC), DrinkID: 6},
		{Amount: 1, CreatedAt: time.Date(2025, 10, 14, 9, 10, 0, 0, time.UTC), DrinkID: 2},
		{Amount: 2, CreatedAt: time.Date(2025, 10, 14, 9, 15, 0, 0, time.UTC), DrinkID: 3},
		{Amount: 1, CreatedAt: time.Date(2025, 10, 14, 9, 20, 0, 0, time.UTC), DrinkID: 5},
	}

	return &DatabaseHandler{
		drinks: drinks,
		orders: orders,
	}
}

func (db *DatabaseHandler) GetDrinks() []model.Drink {
	return db.drinks
}

func (db *DatabaseHandler) GetOrders() []model.Order {
	return db.orders
}

// todo
func (db *DatabaseHandler) GetTotalledOrders() map[uint64]uint64 {
	// calculate total orders
	// key = DrinkID, value = Amount of orders
	// totalledOrders map[uint64]uint64
	totalledOrders := make(map[uint64]uint64)
	for _, order := range db.orders {
		totalledOrders[order.DrinkID] += order.Amount
	}
	return totalledOrders
}

func (db *DatabaseHandler) AddOrder(order *model.Order) {
	// todo
	// add order to db.orders slice
	db.orders = append(db.orders, *order)
}
