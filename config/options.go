package config

import (
	"os"
	"path/filepath"
)

const (
	defaultListenAddr = "127.0.0.1:22333"
	defaultLilacLog   = ".lilac"
	defaultLilacRepo  = "Code/archlinuxcn/repo"
	defaultRepoName   = "archlinuxcn"
)

// Options contains configuration options.
type Options struct {
	listenAddr string
	lilacLog   string
	lilacRepo  string
	repoName   string
}

// NewOptions returns Options with default values.
func NewOptions() *Options {
	home, _ := os.UserHomeDir()
	return &Options{
		listenAddr: defaultListenAddr,
		lilacLog:   filepath.Join(home, defaultLilacLog),
		lilacRepo:  filepath.Join(home, defaultLilacRepo),
		repoName:   defaultRepoName,
	}
}

func (o *Options) ListenAddr() string {
	return o.listenAddr
}

func (o *Options) LilacLog() string {
	return o.lilacLog
}

func (o *Options) LilacRepo() string {
	return o.lilacRepo
}

func (o *Options) RepoName() string {
	return o.repoName
}
