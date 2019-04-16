// Package store provides a set of functions, types and interfaces that
// assist in working with data stores.
// Created By: Thomas Perry (tjamesperry@hotmail.com)
package store

import (
	"database/sql"
	"log"
)

// Queryer is a single method interface encompassing the action of Querying a
// datastore.
type Queryer interface {
	Query(string, ...interface{}) (*sql.Rows, error)
}

// QueryRower is a single method interface encompasing the action of
// QueryRowing a datastore.
type QueryRower interface {
	QueryRow(string, ...interface{}) *sql.Row
}

// Queryest encompasses the Queryer and QueryRower interfaces.
type Queryest interface {
	Queryer
	QueryRower
}

// Execer is a single method interface encompassing the action of Execing a
// SQL statement in a datastore.
type Execer interface {
	Exec(SQL string, args ...interface{}) (sql.Result, error)
}

// QueryExecer encompasses the Queryest and Exerer interfaces.
type QueryExecer interface {
	Queryest
	Execer
}

// Selector is a single method interface that encapsulates the process of
// SELECTing an object's data from a Queryest.
type Selector interface {
	Select(Queryest) error
}

// Creator is a single method interface that encapsulates the process of
// INSERTing an object's data into an Execer.
type Creator interface {
	Create(Queryest) error
}

// Deletor is a single method interface that encapsulates the process of
// DELETEing an object's data from an Execer.
type Deletor interface {
	Delete(Execer) error
}

// Updater is a single method interface that encapsulates the process of
// UPDATEing an objects data in an Execer.
type Updater interface {
	Update(Execer) error
}

// Storer encompasses various store interfaces.
type Storer interface {
	Selector
	Creator
	Deletor
	Updater
}

// Store is a type that manages the access and modification of a
// Storer.
type Store struct {
	Conn QueryExecer
}

// DefaultPostgresLocalhostConnstr is the default connection string for a
// localhost postgres db.
const DefaultPostgresLocalhostConnStr = "postgres://postgres:password@localhost:5432/postgres?sslmode=disable"

// New initializes a Store object and returns it to the caller. The default
// Store.Conn is a postgres localhost DB. To configure the Store object pass a
// set of option functions to New.
func New(options ...func(*Store)) *Store {
	store := new(Store)
	for _, option := range options {
		option(store)
	}

	if store.Conn == nil {
		db, err := sql.Open("postgres", DefaultPostgresLocalhostConnStr)
		if err != nil {
			log.Fatal(err)
		}
		if err := db.Ping(); err != nil {
			log.Fatal(err)
		}
		store.Conn = db
	}

	return store
}

// WithConn sets the Store.Conn field to a Storer object.
func WithConn(c QueryExecer) func(*Store) {
	return func(s *Store) {
		s.Conn = c
	}
}

// Select retrieves data from the Datastore as specified by the Selector.
func (ds *Store) Select(s Selector) error {
	return s.Select(ds.Conn)
}

// Create inserts data into the Datastore as specified by the Creator.
func (ds *Store) Create(c Creator) error {
	return c.Create(ds.Conn)
}

// Delete deletes data from the Datastore as specified by the Deletor.
func (ds *Store) Delete(d Deletor) error {
	return d.Delete(ds.Conn)
}

// Update updates data in the Store as specified by the Updater.
func (ds *Store) Update(u Updater) error {
	return u.Update(ds.Conn)
}
