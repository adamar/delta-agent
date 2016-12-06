package delta

import (
	"bytes"
	"fmt"
	"github.com/boltdb/bolt"
)

var bucket = []byte("delta")

func SetupDB() (*bolt.DB, error) {

	DB, err := bolt.Open("/var/lib/bolt.db", 0644, nil)
	if err != nil {
		return nil, fmt.Errorf("Failed to create DB File: %s", err)
	}

	DB.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucket(bucket)
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		return nil
	})

	return DB, nil

}

func Put(db *bolt.DB, key []byte, value []byte) error {

	err := db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(bucket)
		err := b.Put(key, value)
		return err
	})

	if err != nil {
		return err
	}

	return nil

}

func Get(db *bolt.DB, key []byte) []byte {

	var val []byte

	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(bucket)
		val = b.Get(key)
		return nil
	})

	return val

}

func GetLimit(db *bolt.DB, min []byte, max []byte) {

	db.View(func(tx *bolt.Tx) error {
		c := tx.Bucket(bucket).Cursor()

		for k, v := c.Seek(min); k != nil && bytes.Compare(k, max) <= 0; k, v = c.Next() {
			fmt.Printf("%s: %s\n", k, v)
		}
		return nil
	})

}
