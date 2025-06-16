package db

import "time"

type Category struct {
	ID          int    `json:"category_id" db:"categoryId"`
	Title       string `json:"title" db:"title"`
	OrderNumber int    `json:"order_number,omitempty" db:"orderNumber"`
}
type News struct {
	ID          int        `json:"news_id" db:"newsId"`
	CategoryID  int        `json:"-" db:"categoryId"`
	Title       string     `json:"title" db:"title"`
	Foreword    string     `json:"foreword" db:"foreword"`
	Content     string     `json:"content,omitempty" db:"content"`
	Author      *string    `json:"author,omitempty" db:"author"`
	CreatedAt   *time.Time `json:"created_at,omitempty" db:"createdAt"`
	PublishedAt *time.Time `json:"published_at,omitempty" db:"publishedAt"`
	Category    *Category  `json:"category" db:"category"`
	TagIDs      []int      `json:"-" db:"tagIds"`
}

type NewsFilter struct {
	CategoryID int `form:"cat,default=0"`
	TagID      int `form:"tag,default=0"`
}

type Tag struct {
	ID    int    `json:"tag_id" db:"tagId"`
	Title string `json:"title" db:"title"`
}
