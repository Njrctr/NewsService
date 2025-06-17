package db

import "time"

type Category struct {
	ID          int `pg:"categoryId,pk"`
	Title       string
	OrderNumber int `pg:"orderNumber"`
}
type News struct {
	ID          int `pg:"newsId,pk"`
	Title       string
	CategoryID  int `pg:"categoryId"`
	Foreword    string
	Content     string
	Author      *string
	CreatedAt   *time.Time `pg:"createdAt"`
	PublishedAt *time.Time `pg:"publishedAt"`
	Category    *Category  `pg:"rel:has-one,fk:categoryId"`
	TagIDs      []int      `pg:"tagIds,array"`
}

type NewsFilter struct {
	CategoryID int `pg:"categoryId"`
	TagID      int
}

type Tag struct {
	ID    int    `pg:"tagId"`
	Title string `pg:"title"`
}
