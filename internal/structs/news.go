package structs

import "time"

type News struct {
	ID          int        `json:"news_id" db:"newsId"`
	CategoryID  int        `json:"-" db:"categoryId"`
	Category    *Category  `json:"category" db:"category"`
	Title       string     `json:"title" db:"title"`
	Foreword    string     `json:"foreword" db:"foreword"`
	Content     string     `json:"content,omitempty" db:"content"`
	Author      *string    `json:"author,omitempty" db:"author"`
	CreatedAt   *time.Time `json:"created_at,omitempty" db:"createdAt"`
	PublishedAt *time.Time `json:"published_at,omitempty" db:"publishedAt"`
	TagIDs      []int      `json:"-" db:"tagIds"`
	Tags        []*Tag     `json:"tags,omitempty" db:"-"`
	StatusID    int        `json:"status_id,omitempty" db:"statusId"`
}

type NewsFilter struct {
	CategoryID int `form:"cat,default=0"`
	TagID      int `form:"tag,default=0"`
}
