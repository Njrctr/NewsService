package structs

type Category struct {
	ID          int    `json:"category_id" db:"categoryId"`
	Title       string `json:"title" db:"title"`
	OrderNumber int    `json:"order_number,omitempty" db:"orderNumber"`
	StatusID    int    `json:"status_id" db:"statusId"`
}
