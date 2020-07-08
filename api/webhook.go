package api

import (
	"net/http"
)

func (h *handler) hookSync(w http.ResponseWriter, r *http.Request) {
	h.kiss.LilacRepo.Sync()
}
