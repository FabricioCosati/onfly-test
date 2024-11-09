package repository

import (
	"github.com/FabricioCosati/onfly-test/internal/domain/dao"
	"gorm.io/gorm"
)

type OrderRepository interface {
	CreateOrder(order *dao.Order) (dao.Order, error)
}

type OrderRepositoryImpl struct {
	Database *gorm.DB
}

func (impl *OrderRepositoryImpl) CreateOrder(order *dao.Order) (dao.Order, error) {
	if err := impl.Database.Create(order).Error; err != nil {
		return *order, err
	}

	return *order, nil
}

func OrderRepositoryInit(db *gorm.DB) *OrderRepositoryImpl {
	return &OrderRepositoryImpl{
		Database: db,
	}
}
