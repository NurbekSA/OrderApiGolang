package service

import (
	"HalykProject/app/models"
	"HalykProject/app/repository"
	"HalykProject/util/exception"
	"context"
)

type OrderService struct {
	repo *repository.OrderRepo
}

func NewOrderService(repo *repository.OrderRepo) *OrderService {
	return &OrderService{repo: repo}
}

func (s *OrderService) GetAllOrders(ctx context.Context) ([]*models.Order, *exception.AppError) {
	return s.repo.GetAllOrders(ctx)
}

func (s *OrderService) GetOrderByID(ctx context.Context, id string) (*models.Order, *exception.AppError) {
	return s.repo.GetOrderByID(ctx, id)
}

func (s *OrderService) GetOrdersByUserId(ctx context.Context, id string) ([]*models.Order, *exception.AppError) {
	return s.repo.GetOrdersByUserId(ctx, id)
}

func (s *OrderService) GetOrderByBoxId(ctx context.Context, id string) ([]*models.Order, *exception.AppError) {
	return s.repo.GetOrderByBoxId(ctx, id)
}

func (s *OrderService) GetOrderByStatus(ctx context.Context, status string) ([]*models.Order, *exception.AppError) {
	return s.repo.GetOrderByStatus(ctx, status)
}

func (s *OrderService) CreateOrder(ctx context.Context, user models.Order) *exception.AppError {
	return s.repo.CreateOrder(ctx, user)
}

func (s *OrderService) Complete(ctx context.Context, id string) *exception.AppError {
	return s.repo.Complete(ctx, id)
}

func (s *OrderService) Update(ctx context.Context, order models.Order) *exception.AppError {
	return s.repo.Update(ctx, order)
}

func (s *OrderService) Extend(ctx context.Context, id string, rentalPeriod int) *exception.AppError {
	return s.repo.Extend(ctx, id, rentalPeriod)
}
