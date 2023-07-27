package models

import (
	"database/sql"
	"time"

	"github.com/guregu/null"
	uuid "github.com/satori/go.uuid"
)

const UserStatsTableName = "user_stats"

var (
	_ = time.Second
	_ = sql.LevelDefault
	_ = null.Bool{}
	_ = uuid.UUID{}
)

type UserStats struct {
	ID            string `gorm:"primary_key;column:id;type;TEXT;" json:"id"`
	UserID        string `gorm:"column:userId;type:TEXT;" json:"userId"`
	TotalMessages int32  `gorm:"column:totalMessages;type:INT4;default:0;" json:"totalMessages"`
	TotalWatched  int32  `gorm:"column:totalWatched;type:INT4;default:0;" json:"totalWatched"`
}

func (us *UserStats) TableName() string {
	return UserStatsTableName
}
