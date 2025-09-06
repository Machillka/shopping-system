package repository

import "github.com/machillka/shopping-system/internal/item"

// 单个商品的仓库模型
type RepoModel struct {
	Id       uint
	RepoName string
	Items    []item.ItemModel
}
