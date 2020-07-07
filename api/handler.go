package api

import (
	"net/http"

	"kiss2u/kiss"
	"kiss2u/storage"

	"miniflux.app/http/response/json"
)

type handler struct {
	store *storage.Storage
	kiss  *kiss.Kiss
}

func (h *handler) test(w http.ResponseWriter, r *http.Request) {
	json.OK(w, r, &struct {
		Status string `json:"status"`
	}{
		Status: "ok",
	})
}
