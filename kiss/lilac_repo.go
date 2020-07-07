package kiss

import (
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

func (l *LilacRepo) GetUsers() error {
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
