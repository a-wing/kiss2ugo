package config

const (
	defaultDatabaseURL = "leveldb"
	defaultListenAddr  = "127.0.0.1:22333"
	defaultLilacLog    = "/home/lilydjwg/.lilac"
	defaultLilacRepo   = "/data/archgitrepo-webhook/archlinuxcn"
)

// Options contains configuration options.
type Options struct {
	databaseURL string
	listenAddr  string
	lilacLog    string
	lilacRepo   string
}

// NewOptions returns Options with default values.
func NewOptions() *Options {
	return &Options{
		databaseURL: defaultDatabaseURL,
		listenAddr:  defaultListenAddr,
		lilacLog:    defaultLilacLog,
		lilacRepo:   defaultLilacRepo,
	}
}

func (o *Options) DatabaseURL() string {
	return o.databaseURL
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
