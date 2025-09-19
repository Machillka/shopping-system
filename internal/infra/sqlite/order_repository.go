package sqlite

import (
	"context"
	"database/sql"
	"time"

	"github.com/machillka/shopping-system/internal/domain"
)

// 实现 domain.OrderRepository 接口
type orderRepo struct {
	db *sql.DB
}

// 返回 sqlsite 版本的 OrderRepository
func NewOrderRepository() domain.OrderRepository {
	return &orderRepo{db: GetDB()}
}

func (r *orderRepo) Save(ctx context.Context, o *domain.Order) error {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()
	_, err = tx.ExecContext(
		ctx,
		`
			INSERT INTO orders(id, user_id, total_amount, status, created_at, updated_at)
			VALUES(?, ?, ?, ?, ?, ?)
			ON CONFLICT(id) DO UPDATE SET
				total_amount=excluded.total_amount,
				status=excluded.status,
				updated_at=excluded.updated_at
		`,
		o.ID, o.UserID, o.TotalAmount, o.Status, o.CreateAt, o.UpdateAt)
	if err != nil {
		return err
	}

	_, err = tx.ExecContext(ctx, `
	DELETE FROM order_items WHERE order_id = ?
	`, o.ID)
	if err != nil {
		return err
	}
	stms, err := tx.PrepareContext(ctx, `
        INSERT INTO order_items(order_id, sku, unit_price, quantity)
        VALUES(?, ?, ?, ?)
    `)
	if err != nil {
		return nil
	}
	defer stms.Close()

	for _, item := range o.Items {
		_, err = stms.ExecContext(
			ctx,
			o.ID,
			item.SKU,
			item.UnitPrice,
			item.Quantity,
		)

		if err != nil {
			return err
		}
	}
	return tx.Commit()
}

func (r *orderRepo) FindById(ctx context.Context, id string) (*domain.Order, error) {
	row := r.db.QueryRowContext(ctx, `
	SELECT user_id, total_amount, status, create_at, update_at
	FROM orders WHERE id = ?
	`, id)
	var userId, status string
	var totalAmount float32
	var createAt, updateAt time.Time

	if err := row.Scan(&userId, &totalAmount, &status, &createAt, &updateAt); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	o := &domain.Order{
		ID:          id,
		UserID:      userId,
		TotalAmount: totalAmount,
		Status:      domain.OrderStatus(status),
		CreateAt:    createAt,
		UpdateAt:    updateAt,
	}

	rows, err := r.db.QueryContext(ctx, `
		SELECT sku, unit_price, quantity
		FROM order_items where order_id = ?
	`, id)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var (
			sku       string
			unitPrice float32
			quantity  int
		)

		if err := rows.Scan(&sku, &unitPrice, &quantity); err != nil {
			return nil, err
		}
		o.Items = append(o.Items, domain.OrderItem{
			SKU:       sku,
			UnitPrice: unitPrice,
			Quantity:  quantity,
		})
	}

	return o, nil
}
