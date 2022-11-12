package queryable

import (
	"database/sql"
)

// Usage:
// In appropriate module that needs exposed to the database,
//
// var Adapter = data.Queryable[module.Model]{
// 	Table:   "core.table",
// 	Columns: "id, start, end, duration, frequency, metadata",
// 	ScanFn: func(rows *sql.Rows) (*module.Model, error) {
// 		var c = &module.Model{}
// 		err := rows.Scan(&c.ID, &c.Start, &c.End, &c.Duration, &c.Frequency, &c.Metadata)
// 		return c, err
// 	},
// }
//
// You can then access models via
// - Adapter.Single(&Filters{})
// - Adapter.Get(&Filters{})
// - Adapter.Create(&Model{})

type Queryable[T any] struct {
	Table    string                           // The schema.name of the table in the db
	Columns  string                           // Comma delimited list of columns
	ScanFn   func(rows *sql.Rows) (*T, error) // Maps the returned sql rows to the model
	InsertFn func(t *T) []interface{}         // Returns the values to be inserted for the entity
	Seed     string                           // String of SQL to be ran when seeding this table
	inner    T
}
