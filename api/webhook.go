package api

import (
	"net/http"

	"miniflux.app/http/response/json"
)

func (h *handler) hookSync(w http.ResponseWriter, r *http.Request) {
	if err := h.kiss.LilacRepo.GetUsers(); err != nil {
		json.ServerError(w, r, err)
		return
	}

	if err := h.kiss.LilacRepo.GetSubName(); err != nil {
		json.ServerError(w, r, err)
		return
	}
}
