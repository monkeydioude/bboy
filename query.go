package bboy

import (
	"github.com/boltdb/bolt"
)

// Query interface is made to be used with BBoy Query Engine
type Query interface {
	// Defines how Update from BBoy engine will behave
	Update(tx *bolt.Tx) error
	// Defines what View from BBoy engine will behave
	View(tx *bolt.Tx) error
	// FetchResult is meant to retrieve result after a call to View
	FetchResult() map[string][]byte
}
