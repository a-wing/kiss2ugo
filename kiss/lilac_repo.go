package kiss

import (
	"fmt"
	"kiss2u/model"
	"kiss2u/storage"
	"kiss2u/util"
)

type LilacRepo struct {
	path  string
	store *storage.Storage
}

func NewLilacRepo(store *storage.Storage, path string) *LilacRepo {
	return &LilacRepo{
		store: store,
		path:  path,
	}
}

func (l *LilacRepo) Sync() error {
	fmt.Println("Start Sync")
	l.GetPkgUsers()
	l.GetSubName()
	l.GetUsers()
	return nil
}

func (l *LilacRepo) GetUsers() error {
	pkgUsers, err := util.LilacGetMaintainers(l.path)
	if err != nil {
		//return err
	}

	// Reverse user and package relationships
	userPkgs := make(map[string][]string)
	for pkgname, maintainers := range pkgUsers {
		for _, maintainer := range maintainers {
			if value, ok := userPkgs[maintainer]; ok {
				userPkgs[maintainer] = append(value, pkgname)
			} else {
				userPkgs[maintainer] = []string{pkgname}
			}
		}
	}

	for user, pkgs := range userPkgs {
		l.store.PutUser(&model.User{
			Name: user,
			Pkgs: pkgs,
		})
	}
	return nil
}

func (l *LilacRepo) GetPkgUsers() error {
	users, err := util.LilacGetMaintainers(l.path)
	if err != nil {
		//return err
	}

	for name, user := range users {
		pkg, err := l.store.FindPkg(name)
		if err != nil {
			//return err
		}
		pkg.Maintainers = user
		l.store.PutPkg(pkg)
	}
	return nil
}

func (l *LilacRepo) GetSubName() error {
	lilac, err := util.LilacGetSplitList(l.path)
	if err != nil {
		//return err
	}

	for name, subname := range lilac {
		pkg, err := l.store.FindPkg(name)
		if err != nil {
			//return err
		}
		pkg.SubName = subname
		l.store.PutPkg(pkg)
	}
	return nil
}
