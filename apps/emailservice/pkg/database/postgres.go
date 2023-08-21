package database

import (
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	MaxOpenConns = 10
	MaxIdleTime  = 1 * time.Minute
)

type Postgres struct {
	db *gorm.DB
}

func NewPostgres(databaseUrl string) (*Postgres, error) {
	db, err := gorm.Open(postgres.Open(databaseUrl))
	if err != nil {
		return nil, err
	}

	return &Postgres{
		db,
	}, nil
}

func (e *Postgres) Init() error {
	d, err := e.db.DB()
	if err != nil {
		return err
	}

	d.SetMaxOpenConns(MaxOpenConns)
	d.SetConnMaxIdleTime(MaxIdleTime)

	return nil
}

func (e *Postgres) Save(model interface{}) error {
	return e.db.Create(model).Error
}
