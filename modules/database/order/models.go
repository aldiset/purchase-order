package order

import (
	"time"
)

type Outlet struct {
	ID        uint64    `gorm:"type:int unsigned;primaryKey;autoIncrement;not null"`
	UserID    uint      `gorm:"not null"`
	Name      string    `gorm:"type:varchar(100);not null"`
	Location  string    `gorm:"type:varchar(100);not null"`
	CreatedAt time.Time `gorm:"type:datetime;not null"`
	UpdatedAt time.Time `gorm:"type:datetime;not null"`
	DeletedAt time.Time `gorm:"type:datetime"`
}
