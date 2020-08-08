package cache

import (
	"kiss2u/model"
)

type HotPkgs []string

func (s *Storage) SetHotPkgs(pkgs []string) error {
	s.HotPkgs = pkgs
	return nil
}

func (s *Storage) GetHotPkgList() []string {
	return s.HotPkgs
}

func (s *Storage) GetHotPkgs() (model.Pkgs, error) {
	var pkgs model.Pkgs

	for _, key := range s.HotPkgs {
		if pkg := s.Pkgs[key]; pkg != nil {
			pkgs = append(pkgs, pkg)
		}
	}

	return pkgs, nil
}

func (s *Storage) PutHotPkg(pkg string) error {
	s.HotPkgs = append(s.HotPkgs, pkg)
	return nil
}

func (s *Storage) CleanHotPkgs() {
	s.HotPkgs = []string{}
}
