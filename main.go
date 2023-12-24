package main

import (
	"fmt"
	"os"

	"github.com/syndtr/goleveldb/leveldb"
)

func main() {
	db, err := leveldb.OpenFile("./db", nil)
	if err != nil {
		os.Exit(-1)
	}
	defer db.Close()

	os.Exit(simpleAccess(db))
	// os.Exit(iteratorAccess(db))
}

func simpleAccess(db *leveldb.DB) int {
	_ = db.Put([]byte("key1"), []byte("value1"), nil)
	_ = db.Put([]byte("key2"), []byte("value2"), nil)

	value, _ := db.Get([]byte("key1"), nil)
	fmt.Printf("value is %s\n", value)

	_, err := db.Get([]byte("key3"), nil)
	if err != nil {
		fmt.Printf("key3 is not exist\n")
	}
	return 0
}

func iteratorAccess(db *leveldb.DB) int {
	iter := db.NewIterator(nil, nil)
	for iter.Next() {
		fmt.Printf("%s value is %s\n", iter.Key(), iter.Value())
	}
	iter.Release()
	return 0
}
