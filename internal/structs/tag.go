package structs

type Tag struct {
	ID       int    `json:"tag_id" db:"tagId"`
	Title    string `json:"title" db:"title"`
	StatusID int    `json:"status_id,omitempty" db:"statusId"`
}
