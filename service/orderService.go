package service

import (
	"github.com/Nathanaelaxl/qresto-api/model"
	"github.com/Nathanaelaxl/qresto-api/repository"
)

type OrderService interface {
	PlaceOrder(orderRequest model.Order) (model.Order, error)
}

type orderService struct {
	orderRepo repository.OrderRepository
}

func NewOrderService(repo repository.OrderRepository) OrderService {
	return &orderService{orderRepo: repo}
}

func (s *orderService) PlaceOrder(orderRequest model.Order) (model.Order, error) {
	// Hitung total harga ulang di backend untuk keamanan
	total := 0
	for _, item := range orderRequest.Items {
		total += item.Harga * item.Qty
	}
	orderRequest.TotalHarga = total
	orderRequest.Status = "Pending"

	err := s.orderRepo.CreateOrder(&orderRequest)
	return orderRequest, err
}
