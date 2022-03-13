package repository

import (
	"fmt"
	"synergycommunity/internal/domain"
	"synergycommunity/internal/domain/entity"

	"github.com/gocraft/dbr"
)

func setPagination(stmt *dbr.SelectStmt, opts entity.Options) *dbr.SelectStmt {
	if opts.OrderBy != "" {
		isAsc := opts.OrderType == domain.ASC
		stmt = stmt.OrderDir(opts.OrderBy, isAsc)
	}

	limit := opts.Limit
	page := opts.Page
	offset := limit * (page - 1)

	stmt = stmt.Offset(offset).Limit(limit)

	return stmt
}

func setFilter(stmt *dbr.SelectStmt, filters []entity.Filter) *dbr.SelectStmt {
	for _, s := range filters {
		var builder dbr.Builder

		switch s.Operator {
		case domain.SignEq:
			builder = dbr.Eq(s.By, s.Value)
		case domain.SignGt:
			builder = dbr.Gt(s.By, s.Value)
		case domain.SignLt:
			builder = dbr.Lt(s.By, s.Value)
		case domain.SignLike:
			builder = dbr.Like(s.By, fmt.Sprintf("%%%s%%", s.Value), "!")
		}

		stmt = stmt.Where(builder)
	}

	return stmt
}
