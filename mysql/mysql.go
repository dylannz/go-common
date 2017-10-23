package mysql

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/dylannz/go-common/env"

	// SQL driver
	_ "github.com/go-sql-driver/mysql"

	log "github.com/Sirupsen/logrus"
	"github.com/cenkalti/backoff"
)

var (
	// ConnBackoffTimeout is the duration before the backoff will timeout
	ConnBackoffTimeout = time.Duration(30) * time.Second

	// Conn is the current database connection
	conn *sql.DB

	// once prevents InitConnection from being called more than once in Conn
	once = sync.Once{}

	// ErrUnableToParseDBConnection is raised when there are missing or invalid details for the database connection.
	ErrUnableToParseDBConnection = errors.New("Unable to parse database connection details")

	// ErrUnableToConnectToDB is raised when a connection to the database cannot be established.
	ErrUnableToConnectToDB = errors.New("Unable to connect to the database")
)

// InitConnection creates a new new connection to the database and verifies that it succeeds.
func InitConnection() {
	db := PG{}
	db.Open()
	conn = db.Conn
}

// SetConnection manually sets the connection.
func SetConnection(db *sql.DB) {
	// This stops InitConnection from being called again and clobbering the connection..
	once.Do(func() {})

	conn = db
}

// PG is a concrete implementation of a database connection
type PG struct {
	Conn    *sql.DB
	sslMode string
}

// Conn is the SQL database connection accessor. If the connection is nil, it will be initialized.
func Conn() *sql.DB {
	if conn == nil {
		once.Do(InitConnection)
	}
	return conn
}

// Open will initialize the database connection or raise an error.
func (db *PG) Open() {
	c, err := sql.Open("mysql", db.connectionString())
	if err != nil {
		log.Error(err)
		log.Fatal(ErrUnableToParseDBConnection)
	}

	db.Conn = c

	err = db.verifyConnection()
	if err != nil {
		log.Error(err)
		log.Fatal(ErrUnableToConnectToDB)
	}
}

// verifyConnection pings the database to verify a connection is established. If the connection cannot be established,
// it will retry with an exponential back off.
func (db PG) verifyConnection() error {
	log.Infof("Attempting to connect to database: %s", db.logSafeConnectionString())

	pingDB := func() error {
		return db.Conn.Ping()
	}

	expBackoff := backoff.NewExponentialBackOff()
	expBackoff.MaxElapsedTime = ConnBackoffTimeout

	err := backoff.Retry(pingDB, expBackoff)
	if err != nil {
		log.Warning(err)
		return ErrUnableToConnectToDB
	}

	log.Info("Connected to database")

	return nil
}

// connectionString returns the database connection string.
func (db PG) connectionString() string {
	password := env.GetString("DB_PASSWORD", "")
	if password != "" {
		password = ":" + password
	}

	connString := fmt.Sprintf(
		"%s%s@(%s:%s)/%s?parseTime=true",
		env.GetString("DB_USER", "mysql"),
		password,
		env.GetString("DB_HOST", "localhost"),
		env.GetString("DB_PORT", "3306"),
		env.MustGetString("DB_NAME"),
	)
	return connString
}

// logSafeConnectionString is the database connection string with the password replace with `****` so it can be logged
// without revealing the password.
func (db PG) logSafeConnectionString() string {
	c := db.connectionString()

	password := env.GetString("DB_PASSWORD", "")
	if password != "" {
		c = strings.Replace(c, password, "****", 1)
	}

	return c
}
