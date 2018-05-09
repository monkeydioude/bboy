package bboy

import (
	"fmt"

	"github.com/boltdb/bolt"
)

// Key implments Query interface
// Update and retrieve specific key/value inside a bucket
type Key struct {
	Bucket []byte
	Key    []byte
	Value  []byte
	Result map[string][]byte
}

// NewKey returns pointer to Key query
func NewKey(bucket, key, value string) *Key {
	k := &Key{
		Bucket: []byte(bucket),
		Key:    []byte(key),
		Value:  nil,
		Result: make(map[string][]byte),
	}

	if value != "" {
		k.Value = []byte(value)
	}

	return k
}

// Update from Query interface
func (q *Key) Update(tx *bolt.Tx) error {
	b, err := tx.CreateBucketIfNotExists(q.Bucket)

	if err != nil {
		return err
	}

	err = b.Put(q.Key, q.Value)

	if err != nil {
		return err
	}

	return nil
}

// View from Query interface
func (q *Key) View(tx *bolt.Tx) error {
	b := tx.Bucket(q.Bucket)

	res := b.Get(q.Key)

	if res == nil {
		return fmt.Errorf("[WARN] Could not get result from bucket '%s' and key '%s'", string(q.Bucket), string(q.Key))
	}

	q.Result[string(q.Key)] = res
	return nil
}

// FetchResult from Query interface
func (q *Key) FetchResult() map[string][]byte {
	return q.Result
}