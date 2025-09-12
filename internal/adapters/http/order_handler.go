package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/machillka/shopping-system/internal/application"
	"github.com/machillka/shopping-system/internal/domain"
)

// 构建 POST /orders 的请求结构
type createOrderRequest struct {
	UserID string `json:"user_id" binding:"required"`
	Items []struct {
		SKU string `json:"sku" binding:"required"`
		UnitPrice float32 `json:"unit_price" binding:"required"`
		Quantity int `json:"quantity" bingding:"required,gt=0"`
	} `json:"items" binding:"required,dive"`
}


// 构建 创建订单成功的响应结构
type createOrderResponse struct {
    OrderID string `json:"order_id"`
}

type OrderHandler struct {
	svc application.OrderUseCase
}

func NewOrderHandler(svc application.OrderUseCase) *OrderHandler {
	return &OrderHandler{ svc: svc }
}

func (h *OrderHandler) CreateOrder(c *gin.Context) {
	var req createOrderRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H {
			"error": err.Error(),
		})
		return
	}
	items := make([]domain.OrderItem, 0, len(req.Items))
	for _, it := range req.Items {
		items = append(items, domain.OrderItem{
			SKU: it.SKU,
			UnitPrice: it.UnitPrice,
			Quantity: it.Quantity,
		})
	}
	id, err := h.svc.Create(c.Request.Context(), application.CreateOrderInput{
		UserID: req.UserID,
		Items: items,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, createOrderResponse{OrderID: id})
}

func (h *OrderHandler) RegisterRoutes(r *gin.Engine) {
	grp := r.Group("/orders")
	{
		grp.POST("")
	}
}