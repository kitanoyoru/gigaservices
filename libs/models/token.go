package models

import (
	"database/sql"
	"time"

	"github.com/lib/pq"
	uuid "github.com/satori/go.uuid"

	"github.com/guregu/null"
)

const TokenTableName = "tokens"

var (
	_ = time.Second
	_ = sql.LevelDefault
	_ = null.Bool{}
	_ = uuid.UUID{}
)

type Token struct {
	ID           string         `gorm:"primary_key;AUTO_INCREMENT;column:id;type;TEXT;" json:"id"`
	AccessToken  string         `gorm:"column:accessToken;type:TEXT;" json:"accessToken"`
	RefreshToken string         `gorm:"column:refreshToken;type:TEXT;" json:"refreshToken"`
	CreatedAt    time.Time      `gorm:"column:createdAt;type:TIMESTAMP;" json:"createdAt"`
	ExpiresIn    int32          `gorm:"column:expiresIn;type:INT4;" json:"expiresIn"`
	Scopes       pq.StringArray `gorm:"column:scopes;type:text[]" json:"scopes"`
}

func (u *Token) TableName() string {
	return TokenTableName
}
