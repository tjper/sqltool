// Package store provides a set of interfaces that assist in working with
// data stores.
// Created By: Thomas Perry (tjamesperry@hotmail.com)
package store

import (
	"context"
)

// Creater an object capable of creating.
type Creater interface {
	Create(context.Context) error
}

// Retriever an object capable of retrieving.
type Retriever interface {
	Retrieve(context.Context) error
}

// Deleter an object capable of deleting.
type Deleter interface {
	Delete(context.Context) error
}

// Updater an object capable of updating.
type Updater interface {
	Update(context.Context) error
}
