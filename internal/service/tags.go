package service

import (
	"context"
	repo "news-service/internal/repository"
	"news-service/internal/structs"
)

func GetTags(ctx context.Context) ([]*structs.Tag, error) {
	return repo.GetTags(ctx)
}
