package logging

import (
	"fmt"
	"log"

	"github.com/getsentry/sentry-go"
)

func AddSentry(dsn string, env string) {

	if dsn == "" || env == "" {
		log.Fatal("dsn and env variables are requried to configure sentry")
	}

	err := sentry.Init(sentry.ClientOptions{
		Dsn:              dsn,
		Environment:      env,
		AttachStacktrace: true,
	})

	if err != nil {
		log.Fatalf("sentry.Init: %s", err)
	}
}

func CaptureException(err error) {
	eventID := sentry.CaptureException(err)
	fmt.Printf("%s sent to sentry", *eventID)
}
