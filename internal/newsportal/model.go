package newsportal

import (
	"news-service/internal/db"
	"time"
)

type Category struct {
	ID          int
	Title       string
	OrderNumber int
}
type News struct {
	ID          int
	Title       string
	Foreword    string
	Content     string
	Author      *string
	PublishedAt time.Time
	Category    Category
	Tags        []Tag
}

type NewsFilter struct {
	CategoryID int `form:"cat,default=0"`
	TagID      int `form:"tag,default=0"`
}

type Tag struct {
	ID    int
	Title string
}

//func newNewsFilter(dto *db.NewsFilter) *NewsFilter {
//	return &NewsFilter{
//		CategoryID: dto.CategoryID,
//		TagID:      dto.TagID,
//	}
//}

func newTag(dto db.Tag) Tag {
	return Tag{
		ID:    dto.ID,
		Title: dto.Title,
	}
}

func newNews(dto *db.News) *News {
	category := newCategory(*dto.Category)
	return &News{
		ID:          dto.ID,
		Title:       dto.Title,
		Foreword:    dto.Foreword,
		Content:     *dto.Content,
		Author:      dto.Author,
		PublishedAt: dto.PublishedAt,
		Category:    category,
	}
}

func newCategory(dto db.Category) Category {
	return Category{
		ID:          dto.ID,
		Title:       dto.Title,
		OrderNumber: dto.OrderNumber,
	}
}
