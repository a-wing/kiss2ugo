package model

type Pkg struct {
	Name        string           `json:"name"`
	SubName     []string         `json:"subname,omitempty"`
	Version     string           `json:"version"`
	Maintainers []string         `json:"maintainers,omitempty"`
	Log         map[int]BuildLog `json:"log,omitempty"`
}

type BuildLog struct {
	During  int    `json:"during"`
	Version string `json:"version"`
	Status  string `json:"status"`
}

func NewPkg() *Pkg {
	return &Pkg{Log: make(map[int]BuildLog)}
}

type Pkgs []*Pkg
