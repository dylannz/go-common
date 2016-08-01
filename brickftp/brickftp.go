package brickftp

import (
	"net/http"
	"time"

	"github.com/HomesNZ/go-common/util"
)

var (
	// permittedActions is set of permitted actions. All other actions that are received will be ignored.
	permittedActions = util.NewSet("create")
)

// BrickFTP represents a webhook API request from BrickFTP
type BrickFTP struct {
	Action      string
	Interface   string
	Path        string
	Destination string
	At          time.Time
	Username    string
	Type        string

	Function func(BrickFTP)
}

// NewFromReq creates a new BrickFTP from an HTTP request.
func NewFromReq(r *http.Request) BrickFTP {
	params := r.URL.Query()

	return BrickFTP{
		Action:    params.Get("action"),
		Interface: params.Get("interface"),
		Path:      params.Get("path"),
		Username:  params.Get("username"),
		Type:      params.Get("type"),
	}
}

// Execute will execute the webhook Function within its own Go routine. It is the responsibility of the caller to ensure
// the webhook is permitted before calling Execute.
func (b BrickFTP) Execute() error {
	go b.Function(b)

	return nil
}

// Permitted currently just checks if the action is permitted, however if additional checks are required, they
// should be called here.
func (b BrickFTP) Permitted() bool {
	if ok := b.actionPermitted(); !ok {
		return false
	}

	return true
}

func (b BrickFTP) actionPermitted() bool {
	return permittedActions.Contains(b.Action)
}
