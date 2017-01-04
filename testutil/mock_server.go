package testutil

import (
	"net/http"
	"net/http/httptest"
	"sync/atomic"
)

const (
	// DefaultContentType is the default content type for the MockFixtureServer.
	DefaultContentType = "application/json"

	// DefaultStatus is the default status for the MockFixtureServer.
	DefaultStatus = http.StatusOK
)

// MockFixtureServer is an httptest.Server which always responses with Fixture and Status and increments RequestCounter
// on each request.
type MockFixtureServer struct {
	*httptest.Server

	RequestCount int64
	Fixture      []byte
	Status       int
	ContentType  string
}

// NewMockFixtureServer initializes a new MockFixtureServer.
func NewMockFixtureServer() *MockFixtureServer {
	m := &MockFixtureServer{
		ContentType: DefaultContentType,
		Status:      DefaultStatus,
	}
	m.Server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		atomic.AddInt64(&m.RequestCount, 1)
		w.WriteHeader(m.Status)
		w.Header().Set("Content-Type", m.ContentType)
		w.Write(m.Fixture)
	}))

	return m
}

// ResetRequestCount resets RequestCount so the MockFixtureServer can be shared between tests.
func (m *MockFixtureServer) ResetRequestCount() {
	atomic.AddInt64(&m.RequestCount, 0-m.RequestCount)
}
