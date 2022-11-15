package nova

import (
	"github.com/brandonclapp/nova/data"
)

func main() {
	nova := New()

	data.DB.AutoMigrate(
	// &models.Movie{},
	)

	// Execute app seeds

	// App routes

	nova.Run()
}
