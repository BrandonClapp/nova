package queryable

import (
	"fmt"

	sq "github.com/Masterminds/squirrel"
)

// TODO: Update return type to  ([]*T, error)
// TODO: Better yet, experiment with returning []*Queryable[T]
// TODO: #5 - Support pagination
func (q *Queryable[T]) Get(filters *Filters) []*T {

	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

	query := psql.Select(q.Columns).From(q.Table)
	query = ApplyFilters(query, filters)

	sql, args, err := query.ToSql()

	if err != nil {
		panic(err)
	}

	rows, err := DB.Query(sql, args...)

	if err != nil {
		panic(err)
	}

	// TODO: Don't panic here. Just return nil
	defer rows.Close()

	records := []*T{}

	for rows.Next() {
		t, err := q.ScanFn(rows)
		if err != nil {
			panic(err)
		}

		records = append(records, t)
	}

	return records
}

func (q *Queryable[T]) GetOne(filters *Filters) (*T, error) {
	arr := q.Get(filters)

	if len(arr) != 1 {
		return nil, fmt.Errorf("could not find single record from `%s`. found: %v", q.Table, len(arr))
	}

	return arr[0], nil
}
