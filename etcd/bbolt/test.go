package main

import (
	"fmt"
	"log"

	"time"

	bolt "go.etcd.io/bbolt"
)

var world = []byte("world")

var world2 = []byte("world22")

func main() {
	db, err := bolt.Open("/tmp/bolt.db", 0644, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	key := []byte("k_1000")
	value := []byte("Hello World!")

	/*// retrieve the data
	err = db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(world)
		if bucket == nil {
			return fmt.Errorf("Bucket %q not found!", world)
		}

		val := bucket.Get(key1)
		fmt.Println("haha2: " + string(val))

		return nil
	})*/

	// store some data
	err = db.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists(world)

		if err != nil {
			return err
		}

		err = bucket.Put(key, value)
		if err != nil {
			return err
		}

		for i := 0; i < 30; i += 10 {
			k := fmt.Sprintf("k_%v", i)
			v := fmt.Sprintf("v_%v", i)
			err = bucket.Put([]byte(k), []byte(v))
			if err != nil {
				return err
			}
		}
		bucket2, err := tx.CreateBucketIfNotExists(world2)

		if err != nil {
			return err
		}

		err = bucket2.Put(key, value)
		if err != nil {
			return err
		}

		for i := 0; i < 30; i += 10 {
			k := fmt.Sprintf("22_%v", i)
			v := fmt.Sprintf("55_%v", i)
			err = bucket.Put([]byte(k), []byte(v))
			if err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		log.Fatal(err)
	}

	time.Sleep(time.Minute * 10)
	// retrieve the data
	err = db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(world)
		if bucket == nil {
			return fmt.Errorf("Bucket %q not found!", world)
		}

		val := bucket.Get(key)
		fmt.Println("haha1: " + string(val))

		return nil
	})

	if err != nil {
		log.Fatal(err)
	}
}
