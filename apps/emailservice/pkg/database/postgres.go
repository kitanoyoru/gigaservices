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

	d, err := db.DB()
	if err != nil {
		return nil, err
	}

	d.SetMaxOpenConns(MaxOpenConns)
	d.SetConnMaxIdleTime(MaxIdleTime)

	return &Postgres{
		db,
	}, nil
}
