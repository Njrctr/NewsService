package rpc

import (
	"news-service/internal/newsportal"
	"time"
)

type Category struct {
	ID    int    `json:"category_id"`
	Title string `json:"title"`
}
type News struct {
	ID          int       `json:"news_id"`
	Title       string    `json:"title"`
	Foreword    string    `json:"foreword"`
	Content     string    `json:"content"`
	Author      *string   `json:"author"`
	PublishedAt time.Time `json:"published_at"`
	Category    Category  `json:"category"`
	Tags        []Tag     `json:"tags,omitempty"`
}

type NewsFilter struct {
	CategoryID int `json:"cat"`
	TagID      int `json:"tag"`
}

type Tag struct {
	ID    int    `json:"tag_id" db:"tagId"`
	Title string `json:"title" db:"title"`
}

func newCategory(dto newsportal.Category) Category {
	return Category{
		ID:    dto.ID,
		Title: dto.Title,
	}
}
func newCategorySlice(dtos []newsportal.Category) []Category {
	res := make([]Category, 0, len(dtos))
	for _, c := range dtos {
		res = append(res, Category{
			ID:    c.ID,
			Title: c.Title,
		})
	}
	return res
}

func newTags(dtos []newsportal.Tag) []Tag {
	res := make([]Tag, 0, len(dtos))
	for _, t := range dtos {
		res = append(res, Tag{
			ID:    t.ID,
			Title: t.Title,
		})
	}

	return res
}

func newNews(dto newsportal.News) News {
	category := newCategory(dto.Category)
	tags := newTags(dto.Tags)
	return News{
		ID:          dto.ID,
		Title:       dto.Title,
		Foreword:    dto.Foreword,
		Content:     dto.Content,
		Author:      dto.Author,
		PublishedAt: dto.PublishedAt,
		Category:    category,
		Tags:        tags,
	}
}

func newNewsSlice(dtos []newsportal.News) []News {
	req := make([]News, 0, len(dtos))
	for _, n := range dtos {
		category := newCategory(n.Category)
		tags := newTags(n.Tags)
		req = append(req, News{
			ID:          n.ID,
			Title:       n.Title,
			Foreword:    n.Foreword,
			Content:     n.Content,
			Author:      n.Author,
			PublishedAt: n.PublishedAt,
			Category:    category,
			Tags:        tags,
		})
	}

	return req
}
