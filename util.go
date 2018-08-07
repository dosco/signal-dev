package main

import "github.com/boltdb/bolt"

func writeDB(id []byte, key, value []byte) error {
	return db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists(id)
		if err != nil {
			return err
		}

		return b.Put(key, value)
	})
}

func readDB(id []byte, key []byte) ([]byte, error) {
	var v []byte
	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(id)
		v = b.Get(key)
		return nil
	})
	return v, err
}
