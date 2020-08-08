package cache

type Storage struct {
	Pkgs
	Users
	HotPkgs
}

func NewStorage() *Storage {
	return &Storage{
		Pkgs:  make(Pkgs),
		Users: make(Users),
	}
}
