package client

import (
	"encoding/json"
	"fmt"
	"net/http"

	"kiss2u/model"
)

const (
	UrlAllPackages = "packages"
	UrlFindPackage = "packages/%s"
	UrlAllUsers    = "users"
	UrlFindUser    = "users/%s"
)

type Client struct {
	API string
}

func NewClient(api string) *Client {
	return &Client{
		API: api,
	}
}

func (c *Client) GetPkgs() (pkgs *model.Pkgs, err error) {
	res, err := http.Get(c.API + UrlAllPackages)
	if err != nil {
		return
	}

	err = json.NewDecoder(res.Body).Decode(&pkgs)
	return
}

func (c *Client) FindPkg(name string) (pkg *model.Pkg, err error) {
	res, err := http.Get(c.API + fmt.Sprint(UrlFindPackage, name))
	if err != nil {
		return
	}

	err = json.NewDecoder(res.Body).Decode(&pkg)
	return
}

func (c *Client) GetUsers() (users *model.Users, err error) {
	res, err := http.Get(c.API + UrlAllUsers)
	if err != nil {
		return
	}

	err = json.NewDecoder(res.Body).Decode(&users)
	return
}

func (c *Client) FindUser(name string) (user *model.User, err error) {
	res, err := http.Get(c.API + fmt.Sprint(UrlFindUser, name))
	if err != nil {
		return
	}

	err = json.NewDecoder(res.Body).Decode(&user)
	return
}
