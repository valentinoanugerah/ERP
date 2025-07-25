package model

import "time"

type Transaction struct {
	ID             uint      `gorm:"primaryKey;autoIncrement"`
	TransactionType string    `gorm:"type:varchar(50)"`
	Amount         float64   `gorm:"type:numeric(12,2)"`
	Description    string    `gorm:"type:text"`
	CreatedAt      time.Time
	CreatedBy      uint
	User           User
}
