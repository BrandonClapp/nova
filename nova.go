package main

import (
	"log"
	"net/http"

	"github.com/brandonclapp/nova/auth"
	"github.com/brandonclapp/nova/config"
	"github.com/brandonclapp/nova/cors"
	"github.com/brandonclapp/nova/data"
	authHandlers "github.com/brandonclapp/nova/handlers/auth"
	"github.com/brandonclapp/nova/handlers/health"
	"github.com/brandonclapp/nova/identity"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type Nova struct {
	Identity *identity.Identity
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
		Identity: &identity.Identity{},
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
	// logging.AddSentry(cfg.Sentry.Dsn, cfg.Sentry.Env)

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

	registerRoutes(nova)

	log.Printf("Running on http://localhost:%s/", port)

	err := http.ListenAndServe(":"+port, nova.Router)

	// Clean up sentry
	// defer sentry.Flush(2 * time.Second)

	if err != nil {
		panic(err)
	}
}

func registerRoutes(nova *Nova) {

	// Serve static files for admin dashboard

	nova.Router.HandleFunc("/health-check", health.HealthCheckHandler)
	nova.Router.HandleFunc("/auth/current-user", authHandlers.CurrentUserHandler)
	nova.Router.HandleFunc("/auth/login", authHandlers.LoginHandler)
	nova.Router.HandleFunc("/auth/logout", authHandlers.LogoutHandler)
	nova.Router.HandleFunc("/auth/users", authHandlers.GetUsersHandler)
}
