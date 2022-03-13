package dto

import (
	"synergycommunity/internal/delivery/api/apimodel"
	"synergycommunity/internal/domain/entity"
)

func OptionsToRest(g entity.Options) apimodel.OptionsInput {
	e := apimodel.OptionsInput{
		Page:      int64(g.Page),
		Limit:     int64(g.Limit),
		OrderType: g.OrderType,
		OrderBy:   g.OrderBy,
		Filters:   make([]apimodel.OptionsFilter, len(g.Filters)),
	}

	for i, v := range g.Filters {
		e.Filters[i] = apimodel.OptionsFilter{
			By:       v.By,
			Operator: v.Operator,
			Value:    v.Value,
		}
	}

	return e
}

func OptionsFromRest(g *apimodel.OptionsInput) entity.Options {
	e := entity.Options{
		Page:      uint64(g.Page),
		Limit:     uint64(g.Limit),
		OrderType: g.OrderType,
		OrderBy:   g.OrderBy,
		Filters:   make([]entity.Filter, len(g.Filters)),
	}

	for i, v := range g.Filters {
		e.Filters[i] = entity.Filter{
			By:       v.By,
			Operator: v.Operator,
			Value:    v.Value,
		}
	}

	return e
}
