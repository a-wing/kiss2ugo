package api

import (
	"io"
	"net/http"

	"miniflux.app/http/request"
	"miniflux.app/http/response/json"
)

func (h *handler) findLog(w http.ResponseWriter, r *http.Request) {
	name := request.RouteStringParam(r, "name")
	reader, err := h.kiss.LilacLog.GetLog(name, request.RouteInt64Param(r, "timestamp"))
	if err != nil {
		json.ServerError(w, r, err)
		return
	}

	_, err = io.Copy(w, reader)
	if err != nil {
		json.ServerError(w, r, err)
		return
	}

	w.Header().Add("Content-type", "application/octet-stream")
	w.Header().Add("content-disposition", "attachment; filename=\""+name+"\"")
}
