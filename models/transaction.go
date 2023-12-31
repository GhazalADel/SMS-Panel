package models

import "time"

type Transaction struct {
	ID        uint      `gorm:"primary_key"`
	AccountID uint      `gorm:"not null"`
	Amount    int64     `gorm:"type:bigint"`
	Status    string    `gorm:"type:varchar(255)"`
	Authority string    `gorm:"type:varchar(255)"`
	CreatedAt time.Time `gorm:"type:datetime"`
}
