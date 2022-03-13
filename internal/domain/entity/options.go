package entity

type Options struct {
	Page      uint64
	Limit     uint64
	OrderType string
	OrderBy   string
	Filters   []Filter
}

type Filter struct {
	By       string
	Operator string
	Value    string
}
