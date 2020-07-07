package storage

import (
	"github.com/syndtr/goleveldb/leveldb"
)

type Storage struct {
	kv *leveldb.DB
}

func NewStorage(db *leveldb.DB) *Storage {
	return &Storage{db}
}
