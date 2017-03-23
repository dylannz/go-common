package newrelic

import (
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
	config := newrelic.NewConfig(appName, env.MustGetString("NEWRELIC_API_KEY"))
	app, err = newrelic.NewApplication(config)
	if err != nil {
		panic(err)
	}
}

// App returns the NewRelic application
func App() newrelic.Application {
	return app
}

// WrapHandleFunc is an alias to newrelic.WrapHandleFunc
var WrapHandleFunc = newrelic.WrapHandleFunc
