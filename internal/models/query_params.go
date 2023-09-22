package models

type QueryParams struct {
	Inactive   *bool  `json:"is_active"`
	Pagination bool   `json:"pagination"`
	Page       int    `json:"page"`
	Limit      int    `json:"limit"`
	TableName  string `json:"table_name"`
	TableId    string `json:"table_id"`
}

type Pagination struct {
	CurrentPage  int   `json:"page_number"`
	TotalPage    int   `json:"total_pages"`
	Limit        int   `json:"limit"`
	CurrentCount int   `json:"current_page_total"`
	TotalCount   int64 `json:"total"`
}
