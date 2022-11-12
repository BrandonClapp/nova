package queryable

import (
	"errors"
	"fmt"

	sq "github.com/Masterminds/squirrel"
)

// updates := map[string]interface{}{"start": "2022-10-01T00:00:00Z", "frequency": "monthly"}
func (q *Queryable[T]) Delete(filters *Filters) ([]interface{}, error) {

	// TODO: Ensure filters != nil

	// append initializes slices, if needed
	var deletedIDs []interface{}

	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

	query := psql.
		RunWith(DB).
		Delete(q.Table)

	query = ApplyDeleteFilters(query, filters)
	query = query.Suffix("RETURNING \"id\"")

	rows, err := query.Query()
	defer rows.Close()

	if err != nil {
		fmt.Printf("could not delete entity: %s", err)
		return nil, err
	}

	for rows.Next() {
		var id interface{}
		rows.Scan(&id)
		deletedIDs = append(deletedIDs, id)
	}

	return deletedIDs, nil
}

func (q *Queryable[T]) DeleteOne(filters *Filters) (interface{}, error) {
	deletedIDs, err := q.Delete(filters)
	if err != nil {
		return nil, err
	}

	if len(deletedIDs) != 1 {
		msg := fmt.Sprintf("expected 1 delete result, found %d", len(deletedIDs))
		return nil, errors.New(msg)
	}

	return deletedIDs[0], nil
}
