package lilac

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"time"
)

const (
	BuildLogJSON = "build-log.json"
	BuildLogPath = "log"
)

const (
	BuildStatusStart = "build start"
	BuildStatusStop  = "build end"
)

type LogJSON struct {
	LoggerName string  `json:"logger_name"`
	Pkgbase    string  `json:"pkgbase"`
	Version    string  `json:"pkg_version"`
	Duration   float64 `json:"elapsed"`
	Status     string  `json:"event"`
	Time       float64 `json:"ts"`
}

// Base
//   path: log base path
//   timezone: filename timezone offset
// Query
//   name: build item name
//   timestamp: Timestamp
func GetLogPath(path, timezone, name string, timestamp int64) (string, error) {
	//timezone := "+08:00"
	filename := func(name string) string {
		return fmt.Sprintf("%s.log", name)
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
			if names, err := ioutil.ReadDir(filepath.Join(path, dir.Name())); err == nil {
				for _, n := range names {
					if n.Name() == filename(name) {
						return filepath.Join(path, dirs[l-i-1].Name(), filename(name)), nil
					}
				}
			}

		}
	}

	return "", fmt.Errorf("Not Found Log")
}
