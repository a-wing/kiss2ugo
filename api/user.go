package api

import (
	"net/http"

	"miniflux.app/http/request"
	"miniflux.app/http/response/json"
)

func (h *handler) users(w http.ResponseWriter, r *http.Request) {
	users, err := h.store.GetAllUsers()
	if err != nil {
		json.ServerError(w, r, err)
		return
	}
	json.OK(w, r, users)
}

func (h *handler) findUserPkg(w http.ResponseWriter, r *http.Request) {
	pkg, err := h.store.FindUserPkg(request.RouteStringParam(r, "name"))
	if err != nil {
		json.ServerError(w, r, err)
		return
	}

	json.OK(w, r, pkg)
}
