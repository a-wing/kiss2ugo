package cache

type Storage struct {
	Pkgs
	Users
}

func NewStorage() *Storage {
	return &Storage{
		Pkgs:  make(Pkgs),
		Users: make(Users),
	}
}
