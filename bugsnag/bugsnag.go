package bugsnag

import (
	"github.com/Sirupsen/logrus"
	"github.com/HomesNZ/go-common/version"
	"github.com/HomesNZ/go-common/env"
	"github.com/bugsnag/bugsnag-go"
)

// InitBugsnag initializes bugsnag to capture panics if BUGSNAG_API_KEY is defined. Note that because bugsnag spawns a
// new process, logs will show some initial duplicate entries.
func InitBugsnag() {
	apiKey := env.Get("BUGSNAG_API_KEY")

	if apiKey != "" {
		bugsnag.Configure(bugsnag.Configuration{
			APIKey:       apiKey,
			ReleaseStage: env.Env(),
			AppVersion:   version.Version,
		})
		logrus.Info("Bugsnag configured to capture panics")
	}
}
