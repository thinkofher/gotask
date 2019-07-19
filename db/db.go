// Package db provides functions and data structures
// for manipulating local go boltdb database in terms of
// creating and editing tasks
package db

import (
	"fmt"
	"time"

	"github.com/boltdb/bolt"
)

var taskBucket = []byte("tasks")

// InitDB initialize default buckets in given path
func InitDB(dbPath string) error {
	db, err := bolt.Open(
		dbPath, 0600, &bolt.Options{Timeout: 1 * time.Second})

	if err != nil {
		return err
	}
	defer db.Close()

	return db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(taskBucket)

		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		return nil
	})
}
// itob returns an 8-byte big endian representation of v.
func itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}

// btoi returns int representation of b.
func btoi(b []byte) int {
	return int(binary.BigEndian.Uint64(b))
}
