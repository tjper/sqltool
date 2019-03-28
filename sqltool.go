// Package sqltool provides a set of functions, types and interfaces that
// assist in working with SQL.
// Created By: Thomas Perry (tjamesperry@hotmail.com)
package sqltool

import "database/sql"

// Queryer is a single method interface encompassing the action of Querying a
// datastore.
type Queryer interface {
	Query(SQL string, args ...interface{}) (sql.Rows, error)
}

// QueryRower is a single method interface encompasing the action of
// QueryRowing a datastore.
type QueryRower interface {
	QueryRow(SQL string, args ...interface{}) sql.Row
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

// Datastore encompasses the Queryest and Exerer interfaces.
type DataStore interface {
	Queryest
	Execer
}
