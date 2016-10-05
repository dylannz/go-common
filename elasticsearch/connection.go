package elasticsearch

import (
	"sync"

	"github.com/olivere/elastic"
)

var (
	conn     *elastic.Client
	initOnce = sync.Once{}
)

func initConn() {
	// Create a client
	var err error
	conn, err = elastic.NewClient()
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
