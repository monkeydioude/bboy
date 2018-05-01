package bboy

import (
	"fmt"

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

// Key implments Query interface
type Key struct {
	Bucket []byte
	Key    []byte
	Value  []byte
	Result map[string][]byte
}

// NewKey returns pointer to Key query, allowing to update and retrieve specific key/value inside a bucket
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

// Bucket implments Query interface, made for updating and retrieving every key/value in a bucket
type Bucket struct {
	Bucket []byte
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
