package application

import (
	"context"
	"errors"

	"github.com/machillka/shopping-system/internal/domain"
)

// 定义输入参数
type CreateOrderInput struct {
	UserID string
	Items []domain.OrderItem
}

// 定义订单服业务接口
type OrderUseCase interface {
	Create(ctx context.Context, input CreateOrderInput) (string, error)
	GetbyId(ctx context.Context, id string) (*domain.Order, error)
	Cancel(ctx context.Context, id string) error
}

type orderService struct {
	repo domain.OrderRepository
	domainSvc domain.OrderDomainService
}

func NewOrderService(repo domain.OrderRepository, domainSvc domain.OrderDomainService) OrderUseCase {
	return &orderService{
		repo: repo,
		domainSvc: domainSvc,
	}
}

// Create 实例
func (s *orderService) Create(ctx context.Context, input CreateOrderInput) (string, error) {
	// 创建订单聚合
	order := domain.NewOrder(input.UserID, input.Items)

	if err := s.repo.Save(ctx, order); err != nil {
		return "", err
	}

	return order.ID, nil
}

func (s *orderService) GetbyId(ctx context.Context, id string) (*domain.Order, error) {
	return  s.repo.FindById(ctx, id)
}

func (s *orderService) Cancel(ctx context.Context, id string) error {
	order, err := s.repo.FindById(ctx, id)
	if err != nil {
		return err
	}

	if order == nil {
		return errors.New("没有找到订单")
	}

	if err := s.domainSvc.ValidateCalcel(order); err != nil {
		return err
	}

	if err := order.Cancel(); err != nil {
		return err
	}

	return s.repo.Save(ctx, order)
}