package testsrv

import (
	"net/http"
	"time"
)

const (
	UUIDHeaderName = "HTTP-Sub-UUID"
)

// ReceivedRequest represents a request that the subscriber received.
type ReceivedRequest struct {
	// The request the subscriber received
	Request *http.Request
	// The receipt time
	Time time.Time
	// Unique identifier of the request. will be returned as UUIDHeaderName in each
	// response that the subscriber sends
	UUID string
}
