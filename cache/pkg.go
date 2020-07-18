package cache

import (
	"errors"

	"kiss2u/model"
)

type Pkgs map[string]*model.Pkg

func (s *Storage) PutPkg(pkg *model.Pkg) error {
	// No Name && No log timestamp
	if _, ok := pkg.Log[0]; pkg.Name == "" || ok {
		return errors.New("Name is null Or log no timestamp")
	}

	s.Pkgs[pkg.Name] = pkg
	return nil
}

func (s *Storage) GetAllPkgs() (model.Pkgs, error) {
	var pkgs model.Pkgs
	for _, pkg := range s.Pkgs {
		pkgs = append(pkgs, pkg)
	}
	return pkgs, nil
}

func (s *Storage) GetMapPkgs() (map[string]*model.Pkg, error) {
	return s.Pkgs, nil
}

func (s *Storage) FindPkg(key string) (*model.Pkg, error) {
	if pkg := s.Pkgs[key]; pkg != nil {
		return pkg, nil
	}

	return model.NewPkg(), nil
}

func (s *Storage) RemovePkg(key string) error {
	delete(s.Pkgs, key)
	return nil
}
