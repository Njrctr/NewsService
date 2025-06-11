package service

import (
	"context"
	repo "news-service/internal/repository"
	"news-service/internal/structs"
)

func GetCategories(ctx context.Context) ([]*structs.Category, error) {
	return repo.GetCategories(ctx)
}
