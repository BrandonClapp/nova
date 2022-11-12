package auth

import (
	"context"
	"net/http"

	"github.com/brandonclapp/nova/identity"
	"github.com/brandonclapp/nova/sessions"
	gSessions "github.com/gorilla/sessions"
)

type CookiePersistenceOptions struct {
	// The backing Postgres database connection string to store session data.
	// If a connection string is provided, we will use the pg store instead
	// of in-memory store.
	ConnString string

	// The secret that will be used to encode/decode sessions.
	SessionKey []byte

	// The name of the session cookie that will be created
	SessionName string
}

// Create a new instance of CookiePersistenceOptions.
// If connString is an empty string (""), In-memory storage will be configured
// instead of Postgres
func NewCookiePerstenceOptions(connString string, sessionSecret string) *CookiePersistenceOptions {
	return &CookiePersistenceOptions{
		ConnString:  connString,
		SessionKey:  []byte(sessionSecret),
		SessionName: "sess",
	}
}

// InjectUserMiddleware handles reading the userId from the session cooking and then
// retreiving and injecting the current user into onto the request context
func InjectUserMiddleware(opts *CookiePersistenceOptions) func(http.Handler) http.Handler {

	if opts.ConnString == "" {
		panic("postgres connection string must be provided to initialize session middleware")
	}

	// Create the gorilla in-memory cookie store
	store := gSessions.NewCookieStore(opts.SessionKey)

	// Design Question: Should this middleware open a connection pool for all of "identity"?
	if opts.ConnString != "" {
		// Open a connection pool for the application
		// store, err := pgstore.NewPGStore("postgres://user:password@127.0.0.1:5432/database?sslmode=verify-full", []byte("secret-key"))
		// if err != nil {
		// 	log.Fatalf(err.Error())
		// }
		// defer store.Close()

		// // Run a background goroutine to clean up expired sessions from the database.
		// defer store.StopCleanup(store.Cleanup(time.Minute * 5))
	}

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			userKeyContext := sessions.UserCtxKey
			storeCtxKey := sessions.HTTPKey("store")
			httpRequestKey := sessions.HTTPKey("http")

			httpContext := sessions.HTTP{
				W: &w,
				R: r,
			}

			ctx := context.WithValue(r.Context(), httpRequestKey, httpContext)
			ctx = context.WithValue(ctx, storeCtxKey, store)

			r = r.WithContext(ctx)

			session, err := store.Get(r, opts.SessionName)

			if err != nil {
				// Session exists but could not be decoded
				// TODO: Delete the session and invoke next.
				next.ServeHTTP(w, r)
				return
			}

			userId := session.Values["userId"]
			if userId == nil {
				// No userId in the cookie, user is unauthenticated.
				// Invoke the request without the User attached to context

				next.ServeHTTP(w, r)
				return
			}

			user := identity.Users.GetUserByID(userId.(uint))

			// If there is a user on our cookie but can't find the user,
			// delete the cookie, don't attach the user to the context.
			if user == nil {
				session.Options.MaxAge = -1
				sessions.SaveSession(ctx, session)
				next.ServeHTTP(w, r)
				return
			}

			ctx = context.WithValue(ctx, userKeyContext, user)

			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}
