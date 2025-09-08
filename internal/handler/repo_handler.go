package handler

import (
	"errors"

	"github.com/machillka/shopping-system/internal/item"
	"github.com/machillka/shopping-system/internal/repository"
)

var RepoTable = make(map[string]*repository.RepoModel)

var (
	ErrRepoNotExist = errors.New("该仓库不存在")
)

func CheckRepoExist(itemName string) bool {
	_, exists := RepoTable[itemName]
	return exists
}

func CalculateTotalPrice(itemName string, itemCount int) (float32, error) {
	if repo, exists := RepoTable[itemName]; exists {
		totalPrice := float32(itemCount) * repo.Item.Price * repo.Item.Discount
		return totalPrice, nil
	}
	return -1, ErrRepoNotExist
}

// 对于未出现过的商品，创建其对应的仓库
func NewRepoHandler(itemName string, itemPrice float32, itemDiscount float32) error {
	newItem, err := item.CreateNewItem(itemName, itemPrice, itemDiscount)
	if err != nil {
		return err
	}

	RepoTable[itemName] = &repository.RepoModel{
		Id:        0,
		Item:      newItem,
		ItemCount: 0,
	}

	return nil
}

// 对于已经出现过的商品, 直接添加数量
func AddItemHandler(itemName string, count int) error {
	if repo, exists := RepoTable[itemName]; exists {
		repo.AddItem(count)
		return nil
	}

	return ErrRepoNotExist
}

// 对于已经出现过的商品, 直接减少数量
func RemoveItemHandler(itemName string, count int) error {
	if repo, exists := RepoTable[itemName]; exists {
		err := repo.RemoveItem(count)
		if err != nil {
			return err
		}
		return nil
	}

	return ErrRepoNotExist
}
