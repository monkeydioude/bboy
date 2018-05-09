package bboy

import (
	"errors"

	"github.com/boltdb/bolt"
)

// Bucket implments Query interface
// Update and retrieve every key/value in a bucket
type Bucket struct {
	Bucket []byte
	Keys   map[string][]byte
	Result map[string][]byte
}

// NewBucket returns new pointer to Bucket query
func NewBucket(bucket string) *Bucket {
	return &Bucket{
		Bucket: []byte(bucket),
		Result: make(map[string][]byte),
	}
}

// Update from Query interface
func (q *Bucket) Update(tx *bolt.Tx) error {
	if q.Keys == nil {
		return errors.New("[ERR ] Keys should not be ")
	}
	b, err := tx.CreateBucketIfNotExists(q.Bucket)

	if err != nil {
		return err
	}

	for k, v := range q.Keys {
		b.Put([]byte(k), v)
	}

	return nil
}

// View from Query interface
func (q *Bucket) View(tx *bolt.Tx) error {
	b := tx.Bucket(q.Bucket)

	err := b.ForEach(func(k, v []byte) error {
		q.Result[string(k)] = v
		return nil
	})

	if err != nil {
		return err
	}

	return nil
}

// FetchResult from Query interface
func (q *Bucket) FetchResult() map[string][]byte {
	return q.Result
}
