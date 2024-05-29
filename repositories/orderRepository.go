package repositories

import (
	"OrderPick/models"

	"github.com/gocql/gocql"
)

type OrderRepository struct {
	session *gocql.Session
}

func NewOrderRepository(session *gocql.Session) *OrderRepository {
	return &OrderRepository{session}
}

func (r *OrderRepository) CreateOrder(order models.Order) error {
	query := "INSERT INTO orders (order_id, item_id, created_at, packing_status) VALUES (?, ?, ?, ?)"
	return r.session.Query(query, order.Order_ID, order.Item_id, order.Created_at, order.Packing_status).Exec()
}

func (r *OrderRepository) GetOrderById(orderId string) (models.Order, error) {
	var order models.Order
	query := `SELECT order_id, item_id, created_at, packing_status FROM orders WHERE order_id = ? LIMIT 1`
	err := r.session.Query(query, orderId).Consistency(gocql.One).Scan(&order.Order_ID, &order.Item_id, &order.Created_at, &order.Packing_status)
	return order, err
}

func (r *OrderRepository) GetOrders(recordPerPage int, pagingState []byte) ([]models.Order, []byte, error) {
	query := "SELECT order_id, item_id, created_at, packing_status FROM orders LIMIT ?"
	iter := r.session.Query(query, recordPerPage).PageState(pagingState).Iter()

	var orders []models.Order
	for {
		var order models.Order
		if !iter.Scan(&order.Order_ID, &order.Item_id, &order.Created_at, &order.Packing_status) {
			break
		}
		orders = append(orders, order)
	}
	if err := iter.Close(); err != nil {
		return nil, nil, err
	}

	nextPageState := iter.PageState()
	return orders, nextPageState, nil
}
