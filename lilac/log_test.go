package lilac

import (
	"os"
	"testing"
)

func TestGetLogPath(t *testing.T) {

	p1 := "test/2020-07-07T01:17:01/fish-git.log"
	p2 := "test/2020-07-07T09:03:01/alacritty-git.log"

	os.MkdirAll(p1, os.ModePerm)
	os.MkdirAll(p2, os.ModePerm)

	path, _ := GetLogPath("test", "+08:00", "fish-git", 1687954859)
	if path != p1 {
		t.Errorf(path)
	}

	path, _ = GetLogPath("test", "+08:00", "fish-git", 1587954859)
	if path == p1 {
		t.Errorf(path)
	}

	os.RemoveAll("test/")
}
