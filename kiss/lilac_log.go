package kiss

import (
	"bufio"
	"encoding/json"
	"io"
	"os"
	"path/filepath"

	storage "kiss2u/cache"
	"kiss2u/lilac"
	"kiss2u/model"

	"github.com/hpcloud/tail"
)

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
	file, err := os.Open(filepath.Join(l.path, lilac.BuildLogJSON))
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		l.HandleLog(scanner.Bytes())
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}

func (l *LilacLog) WatchJSON() error {
	t, err := tail.TailFile(filepath.Join(l.path, lilac.BuildLogJSON), tail.Config{Follow: true})
	if err != nil {
		return err
	}
	for line := range t.Lines {
		l.HandleLog([]byte(line.Text))
	}
	return nil
}

func (l *LilacLog) HandleLog(data []byte) error {
	var log lilac.LogJSON
	err := json.Unmarshal(data, &log)
	if err != nil {
		return err
	}

	pkg, _ := l.store.FindPkg(log.Pkgbase)

	pkg.Name = log.Pkgbase
	pkg.Version = log.Version
	pkg.Log[int(log.Time)] = model.BuildLog{
		Version:  log.Version,
		Duration: int(log.Duration),
		Status:   log.Status,
	}

	// Hot pkg
	l.store.PutHotPkg(log.Pkgbase)
	if log.Status == lilac.BuildStatusStart {
		l.store.CleanHotPkgs()
	}

	return l.store.PutPkg(pkg)
}

func (l *LilacLog) GetLog(name string, timestamp int64) (io.Reader, error) {
	path, err := lilac.GetLogPath(filepath.Join(l.path, lilac.BuildLogPath), "+08:00", name, timestamp)
	if err != nil {
		r, _ := io.Pipe()
		return r, err
	}
	return os.Open(path)
}
