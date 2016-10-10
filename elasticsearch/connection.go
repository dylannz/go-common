package elasticsearch

import (
	"strings"
	"sync"

	"github.com/HomesNZ/go-common/env"

	"gopkg.in/olivere/elastic.v3"
)

var (
	conn     *elastic.Client
	initOnce = sync.Once{}
)

func initConn() {
	// Create a client
	var err error
	conn, err = elastic.NewClient(
		elastic.SetURL(strings.Split(env.MustGetString("ELASTICSEARCH_URLS"), ";")...),
	)
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
