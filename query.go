package bboy

import (
	"github.com/coreos/bbolt"
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

// SafeCopy protects []byte from unsafe pointer manipulation from boltdb
func SafeCopy(source []byte) []byte {
	dest := make([]byte, len(source))
	copy(dest, source)
	return dest
}
