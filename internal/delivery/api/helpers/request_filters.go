package helpers

import (
	"synergycommunity/internal/delivery/api/apimodel"
	"synergycommunity/internal/domain"
)

func FilterAllowedFields(filters []apimodel.OptionsFilter, allowed map[string]struct{}) []apimodel.OptionsFilter {
	if len(filters) == 0 {
		return filters
	}

	newFilters := make([]apimodel.OptionsFilter, 0)

	for _, v := range filters {
		_, ok := allowed[v.By]
		if !ok {
			continue
		}

		switch v.Operator {
		case domain.SignEq:
		case domain.SignGt:
		case domain.SignLt:
		case domain.SignLike:
		case "":
			v.Operator = domain.SignEq
		default:
			continue
		}

		newFilters = append(newFilters, v)
	}

	return newFilters
}
