package repository

import (
	"github.com/Nathanaelaxl/qresto-api/model"
	"gorm.io/gorm"
)

type OrderRepository interface {
	CreateOrder(order *model.Order) error
}

type orderRepository struct {
	dbConn *gorm.DB
}

func NewOrderRepository(db *gorm.DB) OrderRepository {
	return &orderRepository{dbConn: db}
}

func (r *orderRepository) CreateOrder(order *model.Order) error {
	return r.dbConn.Create(order).Error
}
