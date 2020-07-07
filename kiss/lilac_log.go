package kiss

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"

	"kiss2u/model"
	"kiss2u/storage"
	"kiss2u/util"

	"github.com/hpcloud/tail"
)

const (
	lilacBuildLogJSON = "/build-log.json"
	lilacBuildLogPath = "/log"
)

type lilacJSONLog struct {
	LoggerName string  `json:"logger_name"`
	Pkgbase    string  `json:"pkgbase"`
	Version    string  `json:"pkg_version"`
	During     float64 `json:"elapsed"`
	Status     string  `json:"event"`
	Time       float64 `json:"ts"`
}

type LilacLog struct {
	store *storage.Storage
	path  string
}

func NewLilacLog(store *storage.Storage, path string) *LilacLog {
	return &LilacLog{
		store: store,
		path:  path,
	}
}

func (l *LilacLog) Migrate() error {
	file, err := os.Open(l.path + lilacBuildLogJSON)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		l.DoLog(scanner.Bytes())
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}

func (l *LilacLog) WatchJSON() error {
	t, err := tail.TailFile(l.path+lilacBuildLogJSON, tail.Config{Follow: true})
	if err != nil {
		return err
	}
	for line := range t.Lines {
		l.DoLog([]byte(line.Text))
	}
	return nil
}

func (l *LilacLog) DoLog(data []byte) error {
	var log lilacJSONLog
	err := json.Unmarshal(data, &log)
	if err != nil {
		return err
	}

	pkg, err := l.store.FindPkg(log.Pkgbase)
	if err != nil {
		fmt.Println(err)
	}

	pkg.Name = log.Pkgbase
	pkg.Version = log.Version
	pkg.Log[int(log.Time)] = model.BuildLog{
		Version: log.Version,
		During:  int(log.During),
		Status:  log.Status,
	}

	return l.store.PutPkg(pkg)
}

func (l *LilacLog) GetLog(name string, timestamp int64) (io.Reader, error) {
	path, err := util.LilacGetLog(l.path+lilacBuildLogPath, name, "+08:00", timestamp)
	if err != nil {
		r, _ := io.Pipe()
		return r, err
	}
	return os.Open(path)
}
