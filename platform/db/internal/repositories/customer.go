package repositories

import (
	"context"

	"github.com/go-kivik/kivik/v4"
	"github.com/google/uuid"
	"github.com/kitanoyoru/gigaservices/platform/db/internal/di"
	"github.com/kitanoyoru/gigaservices/platform/db/internal/models"
	"github.com/samber/do"
)

const CustomerCollectionName = "customers"

type CustomerRepository struct {
	collection *kivik.DB
}

func NewCustomerRepository() *CustomerRepository {
	conn := do.MustInvoke[*kivik.Client](di.Provider)
	coll := conn.DB(CustomerCollectionName)

	return &CustomerRepository{
		collection: coll,
	}
}

func (r *CustomerRepository) GetCustomerById(ctx context.Context, id string) (*models.Customer, error) {
	row := r.collection.Get(ctx, id)
	defer row.Close()

	var customer models.Customer
	if err := row.ScanDoc(&customer); err != nil {
		return nil, err
	}

	return &customer, nil
}

func (r *CustomerRepository) GetCustomerByEmail(ctx context.Context, email string) (*models.Customer, error) {
	query := map[string]any{
		"selector": map[string]any{
			"email": email,
		},
	}

	row := r.collection.Find(ctx, query)
	defer row.Close()

	var customer models.Customer
	if err := row.ScanDoc(&customer); err != nil {
		return nil, err
	}

	return &customer, nil
}

func (r *CustomerRepository) GetAllCustomers(ctx context.Context) ([]*models.Customer, error) {
	rows := r.collection.AllDocs(ctx, kivik.Options{
		"include_docs": true,
	})
	defer rows.Close()

	var customers []*models.Customer
	for rows.Next() {
		var customer models.Customer
		if err := rows.ScanDoc(&customer); err != nil {
			return nil, err
		}
		customers = append(customers, &customer)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return customers, nil
}

func (r *CustomerRepository) AddCustomer(ctx context.Context, customer *models.Customer) (*models.Customer, error) {
	id := uuid.NewString()

	if _, err := r.collection.Put(ctx, id, customer); err != nil {
		return nil, err
	}

	customer.Id = id

	return customer, nil
}
