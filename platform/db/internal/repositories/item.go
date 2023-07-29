package repositories

import (
	"context"

	"github.com/go-kivik/kivik/v4"
	"github.com/google/uuid"
	"github.com/kitanoyoru/gigaservices/platform/db/internal/di"
	"github.com/kitanoyoru/gigaservices/platform/db/internal/models"
	"github.com/samber/do"
)

const ItemRepositoryName = "customers"

type ItemRepository struct {
	collection *kivik.DB
}

func NewItemRepository() *ItemRepository {
	conn := do.MustInvoke[*kivik.Client](di.Provider)
	coll := conn.DB(ItemRepositoryName)

	return &ItemRepository{
		collection: coll,
	}
}

func (r *ItemRepository) GetItemById(ctx context.Context, id string) (*models.Item, error) {
	row := r.collection.Get(ctx, id)
	defer row.Close()

	var item models.Item
	if err := row.ScanDoc(&item); err != nil {
		return nil, err
	}

	return &item, nil
}

func (r *ItemRepository) GetAllItems(ctx context.Context) ([]*models.Item, error) {
	rows := r.collection.AllDocs(ctx, kivik.Options{
		"include_docs": true,
	})
	defer rows.Close()

	var items []*models.Item
	for rows.Next() {
		var item models.Item
		if err := rows.ScanDoc(&item); err != nil {
			return nil, err
		}
		items = append(items, &item)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return items, nil
}

func (r *ItemRepository) AddItem(ctx context.Context, item *models.Item) (*models.Item, error) {
	id := uuid.NewString()

	if _, err := r.collection.Put(ctx, id, item); err != nil {
		return nil, err
	}

	item.Id = id

	return item, nil
}
