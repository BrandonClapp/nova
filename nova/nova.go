package core

import (
	"log"
	"net/http"

	"github.com/brandonclapp/nova/auth"
	"github.com/brandonclapp/nova/config"
	"github.com/brandonclapp/nova/cors"
	"github.com/brandonclapp/nova/data"
	"github.com/brandonclapp/nova/identity"
	id "github.com/brandonclapp/nova/identity"
	"github.com/brandonclapp/nova/logging"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type Nova struct {
	Identity *id.Identity
	DB       *gorm.DB
	Config   *config.Config
	Router   *mux.Router
}

func New() *Nova {
	// load configuration data from env variables
	cfg := config.Get()

	// open sql connection pool in data package for every other package to use
	data.OpenConnectionPool(cfg.Database.ConnString)

	// initialize nova app
	nova := &Nova{
		Identity: &id.Identity{},
		DB:       data.DB,
		Config:   cfg,
		Router:   mux.NewRouter(),
	}

	// Injects the User onto the request context
	opts := auth.NewCookiePerstenceOptions(cfg.Database.ConnString, cfg.Session.Key)
	nova.Router.Use(auth.InjectUserMiddleware(opts))

	// Configure CORS
	nova.Router.Use(cors.AccessControlMiddleware)

	// Add Sentry
	logging.AddSentry(cfg.Sentry.Dsn, cfg.Sentry.Env)

	// AutoMigrate all core packages (schema + seed data)
	identity.AutoMigrate()

	return nova
}

func (nova *Nova) Run() {
	port := nova.Config.Host.Port
	if port == "" {
		log.Println("No port in env, defaulting to 8080")
		port = "8080"
	}

	log.Printf("Running on http://localhost:%s/", port)

	err := http.ListenAndServe(":"+port, nova.Router)

	// Clean up sentry
	// defer sentry.Flush(2 * time.Second)

	if err != nil {
		panic(err)
	}
}
