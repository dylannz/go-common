package newrelic

import (
	"fmt"
	"net/http"
	"strings"
	"sync"

	"github.com/HomesNZ/go-common/env"
	newrelic "github.com/newrelic/go-agent"
)

var (
	app      newrelic.Application
	initOnce = sync.Once{}
)

// InitNewRelic initializes the NewRelic configuration and panics if there is an
// error.
func InitNewRelic(appName string) {
	var err error
	e := env.Env()
	if e == "" {
		e = "development"
	}
	config := newrelic.NewConfig(fmt.Sprintf("%s-%s", appName, e), env.MustGetString("NEWRELIC_API_KEY"))
	app, err = newrelic.NewApplication(config)
	if err != nil {
		panic(err)
	}
}

// App returns the NewRelic application
func App() newrelic.Application {
	return app
}

// Middleware is an easy way to implement NewRelic as middleware in an Alice
// chain.
func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		txn := app.StartTransaction(r.URL.Path, w, r)
		for k, v := range r.URL.Query() {
			txn.AddAttribute(k, strings.Join(v, ","))
		}
		defer txn.End()
		next.ServeHTTP(w, r)
	})
}
