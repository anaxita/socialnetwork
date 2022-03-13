package helpers

import (
	"strings"
	"synergycommunity/internal/delivery/api/apimodel"
	"synergycommunity/internal/domain"
)

/*ValidateOptions checks options for filtering, pagination.
Not valid ones are replaced by default ones or ignored.*/
func ValidateOptions(o *apimodel.OptionsInput, fields ...string) *apimodel.OptionsInput {
	if o == nil {
		tmp := apimodel.DefaultOptions
		o = &tmp
	}

	o.OrderType = strings.ToLower(o.OrderType)

	if o.OrderType != domain.DESC {
		o.OrderType = domain.ASC
	}

	if o.Page < domain.Page {
		o.Page = domain.Page
	}

	if o.Limit < domain.LimitMin {
		o.Limit = domain.Limit
	}

	f := make(map[string]struct{}, len(fields))

	for _, field := range fields {
		f[field] = struct{}{}
	}

	_, ok := f[o.OrderBy]
	if !ok {
		o.OrderBy = ""
	}

	allowedFilters := FilterAllowedFields(o.Filters, f)

	o.Filters = allowedFilters

	return o
}

// TODO
func IsAccessAllowed(got []domain.Perm, allowed ...domain.Perm) bool {
	if len(allowed) == 0 {
		return true
	}

loop:
	for _, v := range allowed {
		for _, d := range got {
			if v == d {
				continue loop
			}
		}

		return false
	}

	return true
}

func IsDisallowAccess(got []domain.Perm, disallowed ...domain.Perm) bool {
	if len(disallowed) == 0 {
		return false
	}

	for _, v := range got {
		for _, d := range disallowed {
			if v == d {
				return true
			}
		}
	}

	return false
}

func HasAccess(got, allowed, disallowed []domain.Perm) bool {
	return !IsDisallowAccess(got, disallowed...) && IsAccessAllowed(got, allowed...)
}
