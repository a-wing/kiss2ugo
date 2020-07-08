package model

type User struct {
	Name string   `json:"name"`
	Pkgs []string `json:"packages"`
	//Link map[string]string `json:"link"`
}

func NewUser() *User {
	return &User{}
}

type Users []*User
