package lilac

import (
	"os"
	"path/filepath"
	"testing"
)

const testPath = "test"

func TestGetMaintainers(t *testing.T) {
	conf := `maintainers:
  - github: Universebenzene
    email: universebenzene@sina.com
  - github: MarvelousBlack

build_prefix: extra-x86_64

pre_build: aur_pre_build

post_build: aur_post_build

update_on:
  - aur: wps-office-cn
`

	name := "wps-office-cn"

	os.MkdirAll(filepath.Join(testPath, name), os.ModePerm)
	f, err := os.Create(filepath.Join(testPath, name, ConfName))
	if err != nil {
		t.Error(err)
	}

	f.Write([]byte(conf))
	f.Close()

	defer os.RemoveAll(testPath)

	pkgs, err := GetMaintainers(testPath)
	if err != nil {
		t.Error(err)
	}

	users, ok := pkgs[name]
	if !ok {
		t.Errorf("does not exist")
	}

	u := 0
	for _, user := range users {
		if user == "Universebenzene" || user == "MarvelousBlack" {
			u++
		}
	}

	if u != 2 {
		t.Error(users)
		t.Error(pkgs)
	}

}

func TestGetSplitList(t *testing.T) {
	conf := `frpc
frps
`

	name := "frp"

	os.MkdirAll(filepath.Join(testPath, name), os.ModePerm)
	f, err := os.Create(filepath.Join(testPath, name, ListName))
	if err != nil {
		t.Error(err)
	}

	f.Write([]byte(conf))
	f.Close()

	defer os.RemoveAll(testPath)

	pkgs, err := GetSplitList(testPath)
	if err != nil {
		t.Error(err)
	}

	users, ok := pkgs[name]
	if !ok {
		t.Errorf("does not exist")
	}

	u := 0
	for _, user := range users {
		if user == "frpc" || user == "frps" {
			u++
		}
	}

	if u != 2 {
		t.Error(users)
		t.Error(pkgs)
	}

}
