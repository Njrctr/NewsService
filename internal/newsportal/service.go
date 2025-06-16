package newsportal

import (
	"news-service/internal/db"
)

type Service struct {
	repo *db.Repository
}

func New(repo *db.Repository) *Service {
	return &Service{repo: repo}
}

func pagination(pageNum, pageSize uint) (uint, uint) {
	if pageSize == 0 {
		pageSize = 5 // Дефолтный размер страницы
	}

	if pageNum > 0 {
		pageNum--
	}

	return pageNum, pageSize * pageNum
}
