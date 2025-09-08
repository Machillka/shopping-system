package repository

import (
	"errors"
	"fmt"

	"github.com/machillka/shopping-system/internal/item"
)

// 单个商品的仓库模型
type RepoModel struct {
	Id        uint           `json:"id"`
	Item      item.ItemModel `json:"item"`
	ItemCount int            `json:"item_count"`
}

// 定义单个仓库的操作
type IRepo interface {
	GetItemInfo() (item.ItemModel, error)
	AddItem() error
	RemoveItem() error
}

func (r RepoModel) GetItemInfo() (item.ItemModel, error) {
	if r.ItemCount == 0 {
		return item.ItemModel{}, errors.New("不存在商品")
	}
	return r.Item, nil
}

func (r *RepoModel) AddItem(count int) error {
	r.ItemCount += count
	return nil
}

func (r *RepoModel) RemoveItem(count int) error {
	if r.ItemCount < count {
		return fmt.Errorf("仓库中 %s 数量不足", r.Item.ItemName)
	}
	r.ItemCount -= count
	return nil
}
