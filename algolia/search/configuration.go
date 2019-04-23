package search

import (
	"time"

	"github.com/algolia/algoliasearch-client-go/algolia/transport"
)

// Configuration contains all the different parameters one can change to
// instantiate a new client for the Search API.
type Configuration struct {
	AppID        string
	APIKey       string
	Hosts        []string
	MaxBatchSize int
	Requester    transport.Requester
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	Headers      map[string]string
}
