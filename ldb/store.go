package ldb

import (
	"github.com/asdine/storm/v3"
	_ "github.com/asdine/storm/v3"
)

const UseDB = 1

func Open(path string, opts ...func(*storm.Options) error) *storm.DB {
	db, err := storm.Open(path, opts...)
	if err != nil {
		panic(err)
	}
	return db
}

var (
	globaldb *storm.DB
)

func SetGlobalDB(db *storm.DB) {
	globaldb = db
}

func CloseDB(db *storm.DB) error {
	return db.Close()
}

func Get() *storm.DB {
	return globaldb
}
