package sessions

import (
	"context"
	"fmt"
	"net/http"

	"github.com/brandonclapp/nova/config"
	"github.com/gorilla/sessions"
)

// HTTPKey is the key used to extract the Http struct.
type HTTPKey string

// HTTP is the struct used to inject the response writer and request http structs.
type HTTP struct {
	W *http.ResponseWriter
	R *http.Request
}

// These should probably go in the auth middleware
var UserCtxKey = &userCtxKey{"user"}

type userCtxKey struct {
	name string
}

func getSessionCookieOptions(session *sessions.Session) *sessions.Options {
	session.Options = &sessions.Options{SameSite: http.SameSiteNoneMode, Secure: true, Path: "/", Domain: "nova.io"}

	if config.Get().Session.Domain != "" {
		domain := config.Get().Session.Domain
		fmt.Printf("Setting cookie domain to %s\n", domain)
		session.Options.Domain = domain
	}

	return session.Options
}

func CreateSession(ctx context.Context, userId uint) error {
	session := GetSession(ctx, "sess")

	// Set the SameSite mode via one of the typed constants described
	// at https://golang.org/pkg/net/http/#SameSite
	session.Options = getSessionCookieOptions(session)
	session.Options.MaxAge = 60 * 60 * 24 * 14

	// Setting userId cookie value
	session.Values["userId"] = userId

	// Save session. This will write a cookie to the http response.
	return SaveSession(ctx, session)
}

func DeleteSession(ctx context.Context, userId uint) error {
	session := GetSession(ctx, "sess")
	session.Options = getSessionCookieOptions(session)
	session.Options.MaxAge = -1

	return SaveSession(ctx, session)
}

// GetSession returns a cached session of the given name
// Gorilla returns a session if the named sesssion does not exist
func GetSession(ctx context.Context, name string) *sessions.Session {
	// This relies on having a gorilla CookieStore and the http request/responseWriter in the context
	store, httpContext := getContextStore(ctx), getContextHttpRequest(ctx)

	// Ignore err because a session is always returned even if one doesn't exist
	session, _ := store.Get(httpContext.R, name)
	return session
}

// // SaveSession saves the session by writing it to the response
func SaveSession(ctx context.Context, session *sessions.Session) error {
	httpContext := getContextHttpRequest(ctx)

	// session.Save will also set-cookie on the response
	// but it has to be called before the response body is written
	err := session.Save(httpContext.R, *httpContext.W)
	return err
}

func getContextStore(ctx context.Context) (store *sessions.CookieStore) {
	// TODO: Fix magic strings
	store = ctx.Value(HTTPKey("store")).(*sessions.CookieStore)
	return store
}

func getContextHttpRequest(ctx context.Context) (http HTTP) {
	http = ctx.Value(HTTPKey("http")).(HTTP)
	return http
}
