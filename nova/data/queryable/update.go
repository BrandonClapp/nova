package queryable

import (
	"errors"
	"fmt"

	sq "github.com/Masterminds/squirrel"
)

// updates := map[string]interface{}{"start": "2022-10-01T00:00:00Z", "frequency": "monthly"}
func (q *Queryable[T]) Update(updates map[string]interface{}, filters *Filters) ([]*T, error) {

	// TODO: Ensure filters != nil

	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

	if len(updates) == 0 {
		return nil, errors.New("at least 1 update is required")
	}

	query := psql.
		RunWith(DB).
		Update(q.Table)

	for key, val := range updates {
		query = query.Set(key, val)
	}

	query = ApplyUpdateFilters(query, filters)
	query = query.Suffix("RETURNING \"id\"")

	rows, err := query.Query()
	defer rows.Close()

	if err != nil {
		fmt.Printf("could not update entity: %s", err)
		return nil, err
	}

	records := q.Get(filters)

	return records, nil
}

func (q *Queryable[T]) UpdateOne(updates map[string]interface{}, filters *Filters) (*T, error) {
	updated, err := q.Update(updates, filters)
	if err != nil {
		return nil, err
	}

	if len(updated) != 1 {
		msg := fmt.Sprintf("expected 1 updated result, found %d", len(updated))
		return nil, errors.New(msg)
	}

	return updated[0], nil
}
