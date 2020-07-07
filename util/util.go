package util

import (
	"bufio"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"

	"gopkg.in/yaml.v2"
)

type Lilac struct {
	Maintainers []struct {
		Github string `yaml:"github"`
		Email  string `yaml:"email,omitempty"`
	} `yaml:"maintainers"`
}

// Fork From: https://github.com/imlonghao/archlinuxcn-log/blob/master/main.go#L55
func LilacGetMaintainers(path string) (map[string][]string, error) {
	r := make(map[string][]string)
	packages, err := ioutil.ReadDir(path)
	if err != nil {
		//return r, err
	}

	for _, pkg := range packages {
		if strings.HasPrefix(pkg.Name(), ".") {
			continue
		}

		path := fmt.Sprintf("%s/%s/lilac.yaml", path, pkg.Name())
		conf, err := ioutil.ReadFile(path)
		if err != nil {
			log.Printf("fail to read %s, %v", path, err)
			continue
		}
		pkgInfo := Lilac{}
		err = yaml.Unmarshal(conf, &pkgInfo)
		if err != nil {
			log.Printf("fail to unmarshal %s, %v", path, err)
			continue
		}
		for _, maintainer := range pkgInfo.Maintainers {
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

func LilacGetSplitList(path string) (map[string][]string, error) {
	r := make(map[string][]string)
	packages, err := ioutil.ReadDir(path)
	if err != nil {
		//return r, err
	}

	for _, pkg := range packages {
		if strings.HasPrefix(pkg.Name(), ".") {
			continue
		}

		path := fmt.Sprintf("%s/%s/package.list", path, pkg.Name())

		if file, err := os.Open(path); err == nil {
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

func LilacGetLog(path, name, timezone string, timestamp int64) (string, error) {
	//timezone := "+08:00"
	filename := func(name string) string {
		return fmt.Sprintf("%s.log", name)
	}

	filepath := func(p1, p2 string) string {
		return fmt.Sprintf("%s/%s/", p1, p2)
	}

	dirs, err := ioutil.ReadDir(path)
	if err != nil {
		return "", err
	}

	// Whether the detection time meets
	l := len(dirs)
	for i, _ := range dirs {
		dir := dirs[l-i-1]
		if t, err := time.Parse(time.RFC3339, dir.Name()+timezone); timestamp > t.Unix() && err == nil {

			// Check if there is a record within this time
			if names, err := ioutil.ReadDir(filepath(path, dir.Name())); err == nil {
				for _, n := range names {
					if n.Name() == filename(name) {
						return filepath(path, dirs[l-i-1].Name()) + filename(name), nil
					}
				}
			}

		}
	}

	return "", errors.New("Not Found Log")
}
