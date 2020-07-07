package api

import (
	"net/http"

	"miniflux.app/http/request"
	"miniflux.app/http/response/json"
)

func (h *handler) pkgs(w http.ResponseWriter, r *http.Request) {
	pkgs, err := h.store.GetAllPkgs()
	if err != nil {
		json.ServerError(w, r, err)
		return
	}

	json.OK(w, r, pkgs)
}

func (h *handler) findPkg(w http.ResponseWriter, r *http.Request) {
	pkg, err := h.store.FindPkg(request.RouteStringParam(r, "name"))
	if err != nil {
		json.ServerError(w, r, err)
		return
	}

	json.OK(w, r, pkg)
}
