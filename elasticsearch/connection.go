package elasticsearch

import (
	"net/http"
	"strings"
	"sync"

	"github.com/HomesNZ/go-common/env"

	"github.com/HomesNZ/elastic"
	"github.com/smartystreets/go-aws-auth"
)

var (
	conn     *elastic.Client
	initOnce = sync.Once{}
)

func awsAuth() bool {
	key := env.GetString("AWS_ACCESS_KEY_ID", "")
	secret := env.GetString("AWS_SECRET_ACCESS_KEY", "")
	token := env.GetString("AWS_SECURITY_TOKEN", "")
	return key != "" && secret != "" || token != ""
}

func initConn() {
	// Create a client
	var err error
	options := []elastic.ClientOptionFunc{
		elastic.SetURL(strings.Split(env.MustGetString("ELASTICSEARCH_URLS"), ";")...),
	}
	if awsAuth() {
		options = append(options, elastic.SetPrepareRequest(func(req *http.Request) {
			awsauth.Sign(req)
		}))
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
