package repositories

import (
	"OrderPick/models"

	"github.com/gocql/gocql"
)

type ItemRepository struct {
	session *gocql.Session
}

func NewItemRepository(session *gocql.Session) *ItemRepository {
	return &ItemRepository{session}
}

func (r *ItemRepository) CreateItem(item models.Item) error {
	query := "INSERT INTO item (item_id, name) VALUES (?, ?)"
	return r.session.Query(query, item.Item_id, *item.Name).Exec()
}

func (r *ItemRepository) GetItemById(itemId string) (models.Item, error) {
	var item models.Item
	query := "SELECT * FROM item WHERE item_id = ? LIMIT 1"
	err := r.session.Query(query, itemId).Consistency(gocql.One).Scan(
		&item.Item_id, &item.Name)
	return item, err
}

func (r *ItemRepository) GetItems(recordPerPage int, pagingState []byte) ([]models.Item, []byte, error) {
	query := "SELECT item_id, name FROM item LIMIT ?"
	iter := r.session.Query(query, recordPerPage).PageState(pagingState).Iter()

	var items []models.Item
	for {
		var item models.Item
		if !iter.Scan(&item.Item_id, &item.Name) {
			break
		}
		items = append(items, item)
	}
	if err := iter.Close(); err != nil {
		return nil, nil, err
	}

	nextPageState := iter.PageState()
	return items, nextPageState, nil
}
