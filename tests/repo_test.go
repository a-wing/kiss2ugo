package tests

import (
	"testing"

	"kiss2u/kiss"
	"kiss2u/storage"

	"github.com/syndtr/goleveldb/leveldb"
)

const (
	path = "/home/metal/Code/archlinuxcn/repo/archlinuxcn"
)

func TestLilacRepo(t *testing.T) {
	db, err := leveldb.OpenFile(testDB, nil)
	if err != nil {
		t.Error(err)
	}
	defer db.Close()
	store := storage.NewStorage(db)

	lilacrepo := kiss.NewLilacRepo(store, path)
	lilacrepo.GetUsers()
	//t.Error("E")

}
