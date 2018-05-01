package bboy

import (
	"errors"
	"os"

	"github.com/boltdb/bolt"
)

// BBoy engine struct
type BBoy struct {
	db   *bolt.DB
	path string
	mode os.FileMode
}

const (
	RD_ONLY = 0400
	WR_ONLY = 0200
	RD_WR   = 0600
)

// OpenWrite opens a bbolt DB on write-only mode
func OpenWrite(path string) (*BBoy, error) {
	return NewDB(path, WR_ONLY)
}

// OpenRead opens a bbolt DB on read-only mode
func OpenRead(path string) (*BBoy, error) {
	return NewDB(path, RD_ONLY)
}

// OpenReadWrite opens a bbolt DB on read-only mode
func OpenReadWrite(path string) (*BBoy, error) {
	return NewDB(path, RD_WR)
}

// NewDB opens a bbolt DB using an os.FileMode
func NewDB(path string, mode os.FileMode) (*BBoy, error) {
	db, err := bolt.Open(path, mode, nil)

	if err != nil {
		return nil, errors.New("[ERR ] Could not open DB")
	}

	return &BBoy{
		db:   db,
		path: path,
		mode: mode,
	}, nil
}

func (b *BBoy) ResetLink() (err error) {
	oldLink := b.db
	b.db, err = bolt.Open(b.path, b.mode, nil)

	if err != nil {
		return
	}
	oldLink.Close()
	return nil
}

// Update calls the Update method from bbolt using the Update method from a Query entity
func (b *BBoy) Update(q Query) error {
	return b.db.Update(q.Update)
}

// View calls the View method from bbolt using the View method from a Query entity
func (b *BBoy) View(q Query) error {
	return b.db.View(q.View)
}

// Get combines call to View and FetchResult from Query entity
func (b *BBoy) Get(q Query) (map[string][]byte, error) {
	if err := b.View(q); err != nil {
		return nil, err
	}
	return q.FetchResult(), nil
}

// Close calls bbolt close
func (b *BBoy) Close() {
	b.db.Close()
}
