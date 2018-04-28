package bboy

import (
	"fmt"

	"github.com/boltdb/bolt"
)

type Query interface {
	Update(tx *bolt.Tx) error
	View(tx *bolt.Tx) error
}

type Key struct {
	Bucket []byte
	Key    []byte
	Value  []byte
}

func NewKey(bucket, key, value string) *Key {
	k := &Key{
		Bucket: []byte(bucket),
		Key:    []byte(key),
		Value:  nil,
	}

	if value != "" {
		k.Value = []byte(value)
	}

	return k
}

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

func (q *Key) View(tx *bolt.Tx) error {
	b := tx.Bucket(q.Bucket)

	res := b.Get(q.Key)

	if res == nil {
		return fmt.Errorf("[WARN] Could not get result from bucket '%s' and key '%s'", string(q.Bucket), string(q.Key))
	}

	q.Value = res
	return nil
}
