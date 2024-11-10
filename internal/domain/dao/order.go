package dao

import (
	"time"
)

type OrderCollection = []Order

type Order struct {
	OrderId       int       `gorm:"column:id_order; primary_key; not null"`
	RequesterName string    `gorm:"column:requester_name_order;"`
	Destination   string    `gorm:"column:destination_order;"`
	GoingDate     time.Time `gorm:"column:going_date_order;"`
	ReturnDate    time.Time `gorm:"column:return_date_order;"`
	Status        string    `gorm:"column:status_order;"`
	CreatedAt     time.Time `gorm:"column:created_at_order;"`
	UpdatedAt     time.Time `gorm:"column:updated_at_order;"`
}
