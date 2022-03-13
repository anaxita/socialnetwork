package service

type DataOptions interface {
	Limit() int64
	Page() int64
	Sort() (filed string, isAsc bool)
	Filters() map[string]interface{}
}
