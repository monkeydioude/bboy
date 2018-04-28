package bboy

import (
	"errors"
	"os"

	"github.com/boltdb/bolt"
)

type BBoy struct {
	db *bolt.DB
}

const (
	RD_ONLY = 0400
	WR_ONLY = 0200
)

// WriteDB opens a bbolt DB on write-only mode
func WriteDB(path string) (*BBoy, error) {
	return NewDB(path, WR_ONLY)
}

// ReadDB opens a bbolt DB on read-only mode
func ReadDB(path string) (*BBoy, error) {
	return NewDB(path, RD_ONLY)
}

// NewDB opens a bbolt DB using a FileMode
func NewDB(path string, mode os.FileMode) (*BBoy, error) {
	db, err := bolt.Open(path, mode, nil)

	if err != nil {
		return nil, errors.New("[ERR ] Could not open DB")
	}

	return &BBoy{
		db: db,
	}, nil
}

func (b *BBoy) Update(q Query) error {
	return b.db.Update(q.Update)
}

func (b *BBoy) View(q Query) error {
	return b.db.View(q.View)
}
