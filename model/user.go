package model

type User struct {
	Name string            `json:"name"`
	Link map[string]string `json:"link"`
}
