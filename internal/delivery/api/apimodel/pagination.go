package apimodel

type Pagination struct {
	Page       int64 `json:"page"`
	CountPages int64 `json:"count_pages"`
	CountItems int64 `json:"count_items"`
}
