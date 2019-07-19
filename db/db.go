// Package db provides functions and data structures
// for manipulating local go boltdb database in terms of
// creating and editing tasks.
package db

import (
	"encoding/binary"
	"fmt"
	"time"

	"github.com/boltdb/bolt"
)

var db *bolt.DB
var taskBucket = []byte("tasks")

// InitDB initialize default buckets in given path.
func InitDB(dbPath string) error {
	var err error

	db, err = bolt.Open(
		dbPath, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		return err
	}

	return db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(taskBucket)

		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		return nil
	})
}

func AddTask(t *Task) error {
	return db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(taskBucket)

		id, err := b.NextSequence()
		if err != nil {
			return err
		}
		t.Id = int(id)

		buf, err := t.ToJson()
		if err != nil {
			return err
		}

		return b.Put(itob(t.Id), buf)
	})
}

func GetTask(id int) (Task, error) {
	var t Task
	var jsonTask []byte

	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(taskBucket)
		jsonTask = b.Get(itob(id))
		return nil
	})
	if err != nil {
		return t, err
	}

	err = t.ReadFromJson(jsonTask)
	if err != nil {
		return t, err
	}

	return t, nil

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
