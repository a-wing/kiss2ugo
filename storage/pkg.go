package storage

import (
	"encoding/json"

	"kiss2u/model"

	"github.com/syndtr/goleveldb/leveldb/util"
)

const (
	pkgPrefix = "pkg."
)

func getKeyName(name string) []byte {
	return []byte(pkgPrefix + name)
}

func (s *Storage) PutPkg(pkg *model.Pkg) error {
	data, err := json.Marshal(pkg)
	if err != nil {
		return err
	}

	err = s.kv.Put(getKeyName(pkg.Name), data, nil)
	return err
}

func (s *Storage) GetAllPkgs() (model.Pkgs, error) {
	iter := s.kv.NewIterator(util.BytesPrefix([]byte(pkgPrefix)), nil)
	var pkgs model.Pkgs
	for iter.Next() {
		var pkg model.Pkg
		json.Unmarshal(iter.Value(), &pkg)
		pkgs = append(pkgs, &pkg)
	}
	iter.Release()
	return pkgs, iter.Error()
}

func (s *Storage) GetMapPkgs() (map[string]*model.Pkg, error) {
	iter := s.kv.NewIterator(util.BytesPrefix([]byte(pkgPrefix)), nil)
	pkgs := make(map[string]*model.Pkg)
	for iter.Next() {
		var pkg model.Pkg
		json.Unmarshal(iter.Value(), &pkg)
		pkgs[pkg.Name] = &pkg
	}
	iter.Release()
	return pkgs, iter.Error()
}

func (s *Storage) FindPkg(key string) (*model.Pkg, error) {
	data, err := s.kv.Get(getKeyName(key), nil)
	pkg := model.NewPkg()
	if err != nil {
		return pkg, err
	}

	err = json.Unmarshal(data, pkg)
	if err != nil {
		return pkg, err
	}

	return pkg, nil
}

func (s *Storage) RemovePkg(key string) error {
	return s.kv.Delete(getKeyName(key), nil)
}
