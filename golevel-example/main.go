package main

import (
	"fmt"

	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/filter"
	"github.com/syndtr/goleveldb/leveldb/opt"
)

var key = "tx2"

func main() {
	o := &opt.Options{
		Filter: filter.NewBloomFilter(10),
	}

	db, err := leveldb.OpenFile("/tmp/goleveltest", o)
	if err != nil {
		fmt.Printf("openfile error: %v \n ", err)
		return
	}

	err = writeBatch(err, db)
	if err != nil {
		fmt.Printf("writeBatch error: %v \n ", err)
		return
	}

	if err = scan(db); err != nil {
		fmt.Printf("scan error: %v \n ", err)
		return
	}

	fmt.Println("success")
	defer db.Close()
}

func scan(db *leveldb.DB) error {
	iter := db.NewIterator(nil, nil)
	for iter.Next() {
		fmt.Println("key:", string(iter.Key()), ", value:", string(iter.Value()))
	}
	iter.Release()

	return iter.Error()
}

func writeBatch(err error, db *leveldb.DB) error {
	batch := new(leveldb.Batch)
	for i := 0; i < 10; i++ {
		batch.Put([]byte(fmt.Sprintf("%s_%d", key, i)), []byte("value"))
	}
	batch.Delete([]byte("baz"))
	err = db.Write(batch, nil)
	return err
}
