package model

type Customer struct {
	ID      uint   `gorm:"primaryKey;autoIncrement"`
	Name    string `gorm:"type:varchar(100)"`
	Contact string `gorm:"type:varchar(100)"`

	Sales []Sale
}
