package models

import "time"

const (
	TableName = "email_journal"
)

type JournalModel struct {
	Id    string `gorm:"primary_key;column:id;type:TEXT" json:"id"`
	Email string `gorm:"column:email;type:TEXT" json:"email"`

	IsSended bool `gorm:"column:isSended;type:bool" json:"isSended"`

	CreatedAt time.Time `gorm:"column:createdAt;type:timestamp" json:"createdAt"`
}

func (j *JournalModel) TableName() string {
	return TableName
}
