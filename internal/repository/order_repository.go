package repository

import (
	"github.com/FabricioCosati/onfly-test/internal/domain/dao"
	"gorm.io/gorm"
)

type OrderRepository interface {
	CreateOrder(order *dao.Order) (dao.Order, error)
	UpdateStatusOrder(order *dao.Order, orderId int) (dao.Order, error)
	GetOrderById(orderId int) (dao.Order, error)
	GetOrders(status string) (dao.OrderCollection, error)
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

func (impl *OrderRepositoryImpl) UpdateStatusOrder(order *dao.Order, orderId int) (dao.Order, error) {
	err := impl.Database.Model(order).
		Where(&dao.Order{OrderId: orderId}).
		Update("status_order", order.Status).
		First(&order).
		Error

	if err != nil {
		return *order, err
	}

	return *order, nil
}

func (impl *OrderRepositoryImpl) GetOrderById(orderId int) (dao.Order, error) {
	var order dao.Order
	err := impl.Database.Model(order).
		Where(&dao.Order{OrderId: orderId}).
		First(&order).
		Error

	if err != nil {
		return order, err
	}

	return order, nil
}

func (impl *OrderRepositoryImpl) GetOrders(status string) (dao.OrderCollection, error) {
	var orders dao.OrderCollection

	query := impl.Database.Model(orders)

	if status != "" {
		query = query.Where(&dao.Order{Status: status})
	}

	if err := query.Find(&orders).Error; err != nil {
		return orders, err
	}

	return orders, nil
}

func OrderRepositoryInit(db *gorm.DB) *OrderRepositoryImpl {
	return &OrderRepositoryImpl{
		Database: db,
	}
}
