package lilac

import (
	"bufio"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v2"
)

const (
	ConfName = "lilac.yaml"
	ListName = "package.list"
)

type PkgConf struct {
	Maintainers []struct {
		Github string `yaml:"github"`
		Email  string `yaml:"email,omitempty"`
	} `yaml:"maintainers"`
}

// Fork From: https://github.com/imlonghao/archlinuxcn-log/blob/master/main.go#L55
func GetMaintainers(path string) (map[string][]string, error) {
	r := make(map[string][]string)
	packages, err := ioutil.ReadDir(path)
	if err != nil {
		return r, err
	}

	for _, pkg := range packages {
		if strings.HasPrefix(pkg.Name(), ".") {
			continue
		}

		item := filepath.Join(path, pkg.Name(), ConfName)
		conf, err := ioutil.ReadFile(item)
		if err != nil {
			continue
		}

		var pkgConf PkgConf
		err = yaml.Unmarshal(conf, &pkgConf)
		if err != nil {
			continue
		}
		for _, maintainer := range pkgConf.Maintainers {
			if value, ok := r[pkg.Name()]; ok {
				value = append(value, maintainer.Github)
				r[pkg.Name()] = value
			} else {
				r[pkg.Name()] = []string{maintainer.Github}
			}
		}
	}
	return r, nil
}

func GetSplitList(path string) (map[string][]string, error) {
	r := make(map[string][]string)
	packages, err := ioutil.ReadDir(path)
	if err != nil {
		return r, err
	}

	for _, pkg := range packages {
		if strings.HasPrefix(pkg.Name(), ".") {
			continue
		}

		item := filepath.Join(path, pkg.Name(), ListName)
		if file, err := os.Open(item); err == nil {
			scanner := bufio.NewScanner(file)
			for scanner.Scan() {
				if value, ok := r[pkg.Name()]; ok {
					value = append(value, scanner.Text())
					r[pkg.Name()] = value
				} else {
					r[pkg.Name()] = []string{scanner.Text()}
				}
			}

		}
	}
	return r, nil
}
