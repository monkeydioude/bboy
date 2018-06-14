package bboy

import (
	"errors"
	"fmt"

	"github.com/coreos/bbolt"
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
	if tx == nil {
		return errors.New("empty tx")
	}
	b := tx.Bucket(q.Bucket)
	if b == nil {
		return fmt.Errorf("Bucket '%s' does not exist", q.Bucket)
	}

	res := b.Get(q.Key)

	if res == nil {
		return fmt.Errorf("Could not get result from bucket '%s' and key '%s'", string(q.Bucket), string(q.Key))
	}

	q.Result[string(q.Key)] = SafeCopy(res)
	return nil
}

// FetchResult from Query interface
func (q *Key) FetchResult() map[string][]byte {
	return q.Result
}
