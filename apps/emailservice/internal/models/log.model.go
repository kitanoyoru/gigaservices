package models

import (
	"database/sql"
	"time"

	"github.com/guregu/null"
	uuid "github.com/satori/go.uuid"
)

const (
	TableName = "email_journal"
)

var (
	_ = time.Second
	_ = sql.LevelDefault
	_ = null.Bool{}
	_ = uuid.UUID{}
)

type LogModel struct {
	ID    string `gorm:"primary_key;column:id;type:TEXT" json:"id"`
	Email string `gorm:"column:email;type:TEXT" json:"email"`

	IsSended  bool        `gorm:"column:isSended;type:bool" json:"isSended"`
	ErrReason null.String `gorm:"column:errReson;type:TEXT" json:"errReason"`

	CreatedAt time.Time `gorm:"column:createdAt;type:timestamp" json:"createdAt"`
}

func (l *LogModel) TableName() string {
	return TableName
}
