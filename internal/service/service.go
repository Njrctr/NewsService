package service

// type NewsService interface {

// }

type Service struct {
	NewsServiceI
	TagServiceI
	CategoryServiceI
}

type Repository interface {
	NewsRepo
	TagRepo
	CategoryRepo
}

func NewService(repo Repository) *Service {
	return &Service{
		NewsServiceI: NewNewsService(repo, repo),
		TagServiceI:  NewTagService(repo),
	}
}
