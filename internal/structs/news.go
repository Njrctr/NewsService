package structs

import "time"

type News struct {
	ID          int        `json:"news_id" db:"newsId"`
	CategoryID  int        `json:"category_id" db:"categoryId"`
	Title       string     `json:"title" db:"title"`
	Foreword    string     `json:"foreword" db:"foreword"`
	Content     string     `json:"content" db:"content"`
	Author      string     `json:"author,omitempty" db:"author"`
	CreatedAt   *time.Time `json:"created_at,omitempty" db:"createdAt"`
	PublishedAt *time.Time `json:"published_at,omitempty" db:"publishedAt"`
	TagIDs      []int      `json:"tag_ids,omitempty" db:"tagIds"`
	StatusID    int        `json:"status_id,omitempty" db:"statusId"`
}

type NewsFilter struct {
	CategoryID *int
	TagID      *int
}
