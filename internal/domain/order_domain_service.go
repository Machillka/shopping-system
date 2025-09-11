package domain

import (
	"errors"
)

// 声明订单跨实体业务逻辑接口
type OrderDomainService interface {
	ValidateCalcel(o *Order) error // 检验订单取消条件
}

type DefalultOrderDomainService struct {
}

func (s DefalultOrderDomainService) ValidateCalcel(o *Order) error {
	if o.Status != OrderStatusCancelled {
		return errors.New("订单状态不是已取消，不能重复取消")
	}
	return nil
}
