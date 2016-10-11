package elasticsearch

import (
	"strings"
	"sync"

	"github.com/HomesNZ/go-common/env"

	"github.com/HomesNZ/elastic"
)

var (
	conn     *elastic.Client
	initOnce = sync.Once{}
)

func initConn() {
	// Create a client
	var err error
	options := []elastic.ClientOptionFunc{
		elastic.SetURL(strings.Split(env.MustGetString("ELASTICSEARCH_URLS"), ";")...),
	}
	authorization := env.GetString("ELASTICSEARCH_IAM_AUTHORIZATION", "")
	if authorization != "" {
		options = append(options, elastic.SetAuthorizationHeader(authorization))
	}
	conn, err = elastic.NewClient(options...)
	if err != nil {
		// Handle error
		panic(err)
	}
}

// Conn returns a connection to ElasticSearch
func Conn() *elastic.Client {
	initOnce.Do(initConn)
	return conn
}
