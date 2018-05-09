package bboy

import (
	"fmt"

	"github.com/boltdb/bolt"
)

const (
	Left  = 1
	Right = 2
)

type MaskBucket struct {
	Bucket []byte
	Keys   map[string][]byte
	Mode   uint8
}

func NewMaskBucket(b string, k map[string][]byte, m uint8) *MaskBucket {
	return &MaskBucket{
		Bucket: []byte(b),
		Keys:   k,
		Mode:   m,
	}
}

// Updates the keys inside a bucket using a map of keys
// Right priority means values inside a bucket won't be overridden by the values provided to the update
func (q *MaskBucket) updateRight(tx *bolt.Tx) error {
	b, err := tx.CreateBucketIfNotExists(q.Bucket)

	if err != nil {
		return fmt.Errorf("[ERR ] Could not find bucket %s in MaskBucket.Update. Reason: %s", q.Bucket, err)
	}

	for k, v := range q.Keys {
		_k := []byte(k)
		_v := b.Get(_k)

		if _v != nil {
			continue
		}
		b.Put([]byte(k), v)
	}

	return nil
}

func (q *MaskBucket) Update(tx *bolt.Tx) error {
	if (q.Mode & Right) != 0 {
		return q.updateRight(tx)
	}

	b := &Bucket{
		Bucket: q.Bucket,
		Keys:   q.Keys,
	}

	return b.Update(tx)
}

// Defines what View from BBoy engine will behave
func (q *MaskBucket) View(tx *bolt.Tx) error {
	return nil
}

// FetchResult is meant to retrieve result after a call to View
func (q *MaskBucket) FetchResult() map[string][]byte {
	return nil
}
