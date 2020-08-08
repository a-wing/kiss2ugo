package api

import (
	"net/http"

	"miniflux.app/http/response/json"
)

func (h *handler) hotPkgs(w http.ResponseWriter, r *http.Request) {
	pkgs, err := h.store.GetHotPkgs()
	if err != nil {
		json.ServerError(w, r, err)
		return
	}

	json.OK(w, r, pkgs)
}
