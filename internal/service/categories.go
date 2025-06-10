package service

type CategoryRepo interface {
}
type CategoryServiceI interface {
}

type CategoryService struct {
	repo CategoryRepo
}

func NewCategoryService(repo CategoryRepo) *CategoryService {
	return &CategoryService{repo: repo}
}
