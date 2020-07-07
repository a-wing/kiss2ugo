package tests

import (
	"testing"

	"kiss2u/kiss"
	"kiss2u/storage"

	"github.com/syndtr/goleveldb/leveldb"
)

func TestLilac(t *testing.T) {
	db, err := leveldb.OpenFile(testDB, nil)
	if err != nil {
		t.Error(err)
	}
	defer db.Close()
	store := storage.NewStorage(db)

	lilackiss := kiss.NewLilacLog(store, "./")

	log1 := `{"logger_name": "build", "pkgbase": "emacs-org-mode-git", "nv_version": "1.45.b6e7d3dcb6a927a19b78ca87575fd05509992542", "pkg_version": "1:9.3.7.r666.gb6e7d3dcb-1", "elapsed": 111.30243873596191, "event": "successful", "ts": 1593825901.2374923, "level": "info"}`
	log2 := `{"logger_name": "build", "pkgbase": "emacs-org-mode-git", "nv_version": "1.45.b6e7d3dcb6a927a19b78ca87575fd05509992542", "pkg_version": "1:9.3.7.r666.gb6e7d3dcb-1", "elapsed": 111.30243873596191, "event": "successful", "ts": 1593825902.2374923, "level": "info"}`
	log3 := `{"logger_name": "build", "pkgbase": "firefox-nightly", "nv_version": "80.0a1 04-Jul-2020 00:01", "pkg_version": "80.0a1.20200704.00-1", "elapsed": 414.96375370025635, "event": "successful", "ts": 1593826316.202076, "level": "info"}`

	err = lilackiss.DoLog([]byte(log1))
	err = lilackiss.DoLog([]byte(log2))
	err = lilackiss.DoLog([]byte(log3))
	if err != nil {
		t.Error(err)
	}

	pkg, err := store.FindPkg("firefox-nightly")
	if err != nil {
		t.Error(err)
	}

	pkg, err = store.FindPkg("null_test")
	if err == nil {
		t.Error(err)
	}

	pkg, err = store.FindPkg("emacs-org-mode-git")
	if err != nil {
		t.Error(err)
	}

	if _, ok := pkg.Log[1593825901]; !ok {
		t.Error(pkg)
	}

	if _, ok := pkg.Log[1593825902]; !ok {
		t.Error(pkg)
	}

	// key is null
	if _, ok := pkg.Log[1593825903]; ok {
		t.Error(pkg)
	}
}
