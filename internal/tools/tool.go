package tools

func Pagination(pageNum, pageSize uint) (uint, uint) {
	if pageSize == 0 {
		pageSize = 5 // Дефолтный размер страницы
	}

	if pageNum > 0 {
		pageNum--
	}

	return pageNum, pageSize * pageNum
}
