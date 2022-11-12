package queryable

import (
	"fmt"

	sq "github.com/Masterminds/squirrel"
)

type Filter struct {
	FieldName  string
	Expression string
	Value      interface{}
}

func IDEquals(id interface{}) *Filters {
	return &Filters{
		Filter{
			FieldName:  "id",
			Expression: "eq",
			Value:      id,
		},
	}
}

type Filters []Filter

func ApplyFilters(query sq.SelectBuilder, filters *Filters) sq.SelectBuilder {

	if filters == nil {
		filters = &Filters{}
	}

	for _, f := range *filters {
		if f.Expression == "eq" {
			query = query.Where(sq.Eq{f.FieldName: f.Value})
		}
		if f.Expression == "neq" {
			query = query.Where(sq.NotEq{f.FieldName: f.Value})
		}
		if f.Expression == "like" {
			value := fmt.Sprintf("%%%v%%", f.Value)
			query = query.Where(sq.ILike{f.FieldName: value})
		}
		if f.Expression == "gt" {
			query = query.Where(sq.Gt{f.FieldName: f.Value})
		}
		if f.Expression == "lt" {
			query = query.Where(sq.Lt{f.FieldName: f.Value})
		}
	}

	return query
}

// // TODO: Generic?
func ApplyDeleteFilters(query sq.DeleteBuilder, filters *Filters) sq.DeleteBuilder {
	if filters == nil {
		filters = &Filters{}
	}

	for _, f := range *filters {
		if f.Expression == "eq" {
			query = query.Where(sq.Eq{f.FieldName: f.Value})
		}
		if f.Expression == "neq" {
			query = query.Where(sq.NotEq{f.FieldName: f.Value})
		}
		if f.Expression == "like" {
			value := fmt.Sprintf("%%%v%%", f.Value)
			query = query.Where(sq.ILike{f.FieldName: value})
		}
		if f.Expression == "gt" {
			query = query.Where(sq.Gt{f.FieldName: f.Value})
		}
		if f.Expression == "lt" {
			query = query.Where(sq.Lt{f.FieldName: f.Value})
		}
	}

	return query
}

func ApplyUpdateFilters(query sq.UpdateBuilder, filters *Filters) sq.UpdateBuilder {
	if filters == nil {
		filters = &Filters{}
	}

	for _, f := range *filters {
		if f.Expression == "eq" {
			query = query.Where(sq.Eq{f.FieldName: f.Value})
		}
		if f.Expression == "neq" {
			query = query.Where(sq.NotEq{f.FieldName: f.Value})
		}
		if f.Expression == "like" {
			value := fmt.Sprintf("%%%v%%", f.Value)
			query = query.Where(sq.ILike{f.FieldName: value})
		}
		if f.Expression == "gt" {
			query = query.Where(sq.Gt{f.FieldName: f.Value})
		}
		if f.Expression == "lt" {
			query = query.Where(sq.Lt{f.FieldName: f.Value})
		}
	}

	return query
}
