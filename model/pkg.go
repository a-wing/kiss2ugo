package model

type Pkg struct {
	Name    string           `json:"name"`
	SubName []string         `json:"subname,omitempty"`
	Version string           `json:"version"`
	Users   []string         `json:"users,omitempty"`
	Log     map[int]BuildLog `json:"log,omitempty"`
}

type BuildLog struct {
	Duration int    `json:"duration"`
	Version  string `json:"version"`
	Status   string `json:"status"`
}

func NewPkg() *Pkg {
	return &Pkg{Log: make(map[int]BuildLog)}
}

type Pkgs []*Pkg
