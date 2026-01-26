package handler

import (
	"net/http"
)

type HeartbeatHandler struct {
}

func NewHeartbeatHandler() *HeartbeatHandler {
	return &HeartbeatHandler{}
}

func (h *HeartbeatHandler) Heartbeat(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNoContent)
}
