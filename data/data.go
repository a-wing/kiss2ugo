// +build dev

package data

import (
	"net/http"
)

var (
	Docs = http.Dir("docs/")
)
