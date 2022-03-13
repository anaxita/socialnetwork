package apimodel

import (
	"synergycommunity/internal/domain"
)

const (
	ColID      = "id"
	ColUserID  = "user_id"
	ColGroupID = "group_id"
	ColName    = "name"
	ColText    = "text"

	ColNickname  = "nickname"
	ColFirstName = "first_name"
	ColLastName  = "last_name"
)

type OptionsInput struct {
	Page      int64           `json:"page"`
	Limit     int64           `json:"limit"`
	OrderType string          `json:"order_type"`
	OrderBy   string          `json:"order_by"`
	Filters   []OptionsFilter `json:"filters"`
}

type OptionsFilter struct {
	By       string `json:"by" validation:"required"`
	Operator string `json:"operator" validation:"required"`
	Value    string `json:"value" validation:"required"`
}

var DefaultOptions = OptionsInput{
	Page:      domain.Page,
	Limit:     domain.Limit,
	OrderType: domain.ASC,
	OrderBy:   ColID,
}
