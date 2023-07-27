package database

import (
	kivik "github.com/go-kivik/kivik/v4"
	"github.com/kitanoyoru/gigaservices/platform/db/internal/di"
	"github.com/kitanoyoru/gigaservices/platform/db/internal/repositories"
	"github.com/kitanoyoru/gigaservices/platform/db/pkg/cfg"
	"github.com/samber/do"
)

type DatabaseConnection struct {
	customers *repositories.CustomerRepository
	items     *repositories.ItemRepository
}

func NewDatabaseConnection() (*DatabaseConnection, error) {
	config := do.MustInvoke[cfg.Config](di.Provider)

	client, err := kivik.New("couch", config.Database.ToUrl())
	if err != nil {
		return nil, err
	}

	do.ProvideValue[*kivik.Client](di.Provider, client)

	return &DatabaseConnection{
		customers: repositories.NewCustomerRepository(),
		items:     repositories.NewItemRepository(),
	}, nil
}
