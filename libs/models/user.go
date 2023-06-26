package models

import (
	"database/sql"
	"time"

	"github.com/guregu/null"
	uuid "github.com/satori/go.uuid"
)

const TableName = "users"

var (
	_ = time.Second
	_ = sql.LevelDefault
	_ = null.Bool{}
	_ = uuid.UUID{}
)

type User struct {
	ID      string         `gorm:"primary_key;column:id;type;TEXT;" json:"id"`
	TokenID sql.NullString `gorm:"column:tokenId;type:TEXT;" json:"tokenId"`
	IsAdmin bool           `gorm:"column:isAdmin;type:BOOL;default:false;" json:"isAdmin"`
	ApiKey  string         `gorm:"column:apiKey;type:TEXT;" json:"apiKey"`
	Token   *Token         `gorm:"foreignKey:TokenID" json:"token"`
	Stats   *UserStats     `gorm:"foreignKey:UserID" json:"stats"`
}

func (u *User) TableName() string {
	return TableName
}
