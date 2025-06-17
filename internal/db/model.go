package db

import "time"

type Category struct {
	ID          int    `db:"categoryId"`
	Title       string `db:"title"`
	OrderNumber int    `db:"orderNumber"`
}
type News struct {
	ID          int        `db:"newsId"`
	CategoryID  int        `db:"categoryId"`
	Title       string     `db:"title"`
	Foreword    string     `db:"foreword"`
	Content     string     `db:"content"`
	Author      *string    `db:"author"`
	CreatedAt   *time.Time `db:"createdAt"`
	PublishedAt *time.Time `db:"publishedAt"`
	Category    *Category  `db:"category"`
	TagIDs      []int      `db:"tagIds"`
}

type NewsFilter struct {
	CategoryID int
	TagID      int
}

type Tag struct {
	ID    int    `db:"tagId"`
	Title string `db:"title"`
}
