// Package store provides a set of functions and interfaces that assist in
// working with data stores.
// Created By: Thomas Perry (tjamesperry@hotmail.com)
package store

import (
	"context"
)

// The following types and functions only serve to enforce the store pattern.
// An object should be created to represent a db resource. This object should
// satisfy the necessary interfaces below. The object's logic should be
// executed by passed the object as an argument to appropriate helper function.

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

// Create calls the create method on a Creater.
func Create(ctx context.Context, c Creater) error {
	return c.Create(ctx)
}

// Retrieve calls the Retrieve method on a Retriever.
func Retrieve(ctx context.Context, r Retriever) error {
	return r.Retrieve(ctx)
}

// Delete calls the Delete method on a Deleter.
func Delete(ctx context.Context, d Deleter) error {
	return d.Delete(ctx)
}

// Update calls the Update method on a Updater.
func Update(ctx context.Context, u Updater) error {
	return u.Update(ctx)
}
