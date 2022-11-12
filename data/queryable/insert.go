package queryable

import (
	"fmt"

	sq "github.com/Masterminds/squirrel"
)

func (q *Queryable[T]) Create(entity *T) (*T, error) {
	lastInsertedID := ""
	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

	query := psql.
		RunWith(DB).
		Insert(q.Table).
		Columns(q.Columns).
		Values(q.InsertFn(entity)...).
		Suffix("RETURNING \"id\"") // Assumes every Querable has an id field

	// QueryRow does not return an error
	// query.QueryRow().Scan(&lastInsertedID)
	rows, err := query.Query()
	defer rows.Close()

	if err != nil {
		fmt.Printf("could not create entity: %s", err)
		return nil, err
	}

	// Only 1 record should be created
	rows.Next()
	rows.Scan(&lastInsertedID)

	inserted, err := q.GetOne(&Filters{
		Filter{
			FieldName:  "id",
			Expression: "eq",
			Value:      lastInsertedID,
		},
	})

	if err != nil {
		return nil, err
	}

	return inserted, nil
}
