package queryable

import (
	"fmt"

	"github.com/brandonclapp/nova/utils"
)

// Dependency ordering matters
func Seed[T any](queryables ...Queryable[T]) {
	fmt.Println("\nRunning seed scripts...")
	for _, q := range queryables {
		if q.Seed != "" {
			utils.ExecuteSql(DB, q.Seed)
			fmt.Printf("[✓] %s \n", q.Table)
		}
	}
}
