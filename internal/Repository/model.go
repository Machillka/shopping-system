package repository

import "github.com/machillka/shopping-system/internal/item"

type RepoModel struct {
	Id uint
	RepoName string
	Items []ItemModel
}