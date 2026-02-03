package handler

import (
	"assistant-sf-daemon/internal/locale"
	"assistant-sf-daemon/internal/ucase"
	"net/http"
)

type ConfigHandler struct {
	useCase ucase.ConfigUseCase
}

func NewConfigHandler(uCase ucase.ConfigUseCase) *ConfigHandler {
	return &ConfigHandler{
		useCase: uCase,
	}
}

func (h *ConfigHandler) GetInitialStatus(w http.ResponseWriter, r *http.Request) {
	langRequest := locale.GetLangFromContext(r.Context())

	err := h.useCase.GetStatus()
	if err != nil {
		SendErrorResponse(w, buildErrorMessage(langRequest, err), http.StatusUnprocessableEntity, 1)
		return
	}
	SendResponse(w, http.StatusOK, nil)
}
